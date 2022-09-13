package uniswapv3

import (
	"context"
	"fmt"
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
	"order_service/pkg/wallet/eth"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Config struct {
	Id            string `json:"id,omitempty"`
	ProviderURL   string
	Mnemonic      string
	Addresses     []common.Address
	ConfirmBlocks uint64
	TokensFile    string
	TokensUrl     string
	Msg           string
}

type UniSwapV3 struct {
	mux *sync.Mutex

	cfg       *Config
	accountId string

	wallet   *eth.HDWallet
	Provider *entity.Provider

	confirms  uint64
	blockTime time.Duration
	chainId   *big.Int

	factory *contracts.Uniswapv3Factory

	tokens *supportedTokens
	pairs  *supportedPairs

	tt *txTracker
	dt *depostiTracker
	am *approveManager
	rc *redis.Client
	v  *viper.Viper
	l  logger.Logger

	stopCh    chan struct{}
	stoppedAt time.Time
}

func NewExchange(cfg *Config, rc *redis.Client, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.Exchange, error) {

	agent := "uniswapv3.NewExchange"

	if cfg.ProviderURL == "" {
		return nil, errors.Wrap("Provider URL must cannot be empty")
	}

	if cfg.ConfirmBlocks == 0 {
		cfg.ConfirmBlocks = 1
	}

	if cfg.TokensFile == "" {
		cfg.TokensFile = "./tokens.json"
	}
	if cfg.TokensUrl == "" {
		cfg.TokensUrl = "https://tokens.uniswap.org"
	}

	client, err := ethclient.Dial(cfg.ProviderURL)
	if err != nil {
		return nil, err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	w, err := eth.NewWallet(cfg.Mnemonic, client)
	if err != nil {
		return nil, err
	}

	cfg.Mnemonic = w.Mnemonic()

	v3 := &UniSwapV3{
		mux: &sync.Mutex{},

		wallet: w,
		Provider: &entity.Provider{
			Client: client,
			URL:    cfg.ProviderURL,
		},
		cfg: cfg,

		chainId:   chainId,
		accountId: hash(hash(w.Mnemonic())),

		confirms:  cfg.ConfirmBlocks,
		blockTime: time.Duration(15 * time.Second),

		tokens: newSupportedTokens(),
		pairs:  newSupportedPairs(),

		rc: rc,
		v:  v,
		l:  l,

		stopCh: make(chan struct{}),
	}

	v3.tt = newTxTracker(v3)
	v3.dt = newDepositTracker(v3)
	v3.am = newApproveManager(v3)

	f, err := contracts.NewUniswapv3Factory(factory, v3.Provider)
	if err != nil {
		return nil, err
	}
	v3.factory = f

	if readConfig {
		v3.l.Debug(agent, fmt.Sprintf("retriving pairs from config file %s", v3.v.ConfigFileUsed()))

		i := v3.v.Get(fmt.Sprintf("%s.pairs", v3.NID()))
		if i != nil {
			psi := i.(map[string]interface{})

			for _, v := range psi {
				p := v.(map[string]interface{})
				pair := &pair{}

				if p["bt"] != nil && p["qt"] != nil {
					bChainId := int64(p["bt"].(map[string]interface{})["chainid"].(float64))
					qChainId := int64(p["qt"].(map[string]interface{})["chainid"].(float64))

					if bChainId == qChainId && v3.chainId.Int64() == bChainId {
						pair.BT = token{
							Name:     p["bt"].(map[string]interface{})["name"].(string),
							Symbol:   p["bt"].(map[string]interface{})["symbol"].(string),
							Address:  common.HexToAddress(p["bt"].(map[string]interface{})["address"].(string)),
							Decimals: int(p["bt"].(map[string]interface{})["decimals"].(float64)),
							ChainId:  bChainId,
						}

						pair.QT = token{
							Name:     p["qt"].(map[string]interface{})["name"].(string),
							Symbol:   p["qt"].(map[string]interface{})["symbol"].(string),
							Address:  common.HexToAddress(p["qt"].(map[string]interface{})["address"].(string)),
							Decimals: int(p["qt"].(map[string]interface{})["decimals"].(float64)),
							ChainId:  qChainId,
						}

						pair, err = v3.highestLiquidPool(pair.BT, pair.QT)
						if err != nil {
							v3.l.Error(agent, err.Error())
							continue
						}

						v3.pairs.add(*pair)
						v3.tokens.add(pair.BT, pair.QT)
						v3.l.Debug(agent, fmt.Sprintf("pair %s added", pair.String()))
					}
				}
			}

		}

	}

	v3.cfg.Addresses, err = v3.wallet.AllAddresses()
	if err != nil {
		return nil, err
	}

	return v3, nil
}

func (u *UniSwapV3) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}

	w.Add(1)
	go u.tt.run(w, u.stopCh)

	w.Add(1)
	go u.dt.run(w, u.stopCh)

}
