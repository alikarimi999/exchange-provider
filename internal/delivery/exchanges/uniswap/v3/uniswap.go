package uniswapv3

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Config struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ChianId       uint64 `json:"chian_id,omitempty"`
	Network       string `json:"network,omitempty"`
	NativeToken   string `json:"native_token,omitempty"`
	TokenStandard string `json:"token_standard,omitempty"`

	Providers []*Provider `json:"providers,omitempty"`

	Factory       common.Address
	Router        common.Address
	Mnemonic      string             `json:"mnemonic,omitempty"`
	AccountCount  uint64             `json:"account_count,omitempty"`
	Accounts      []accounts.Account `json:"accounts,omitempty"`
	ConfirmBlocks uint64             `json:"confirm_blocks,omitempty"`
	TokensFile    string             `json:"tokens_file,omitempty"`
}

func (c *Config) wrapNative() string {
	return fmt.Sprintf("W%s", c.NativeToken)
}

type dex struct {
	mux *sync.Mutex

	cfg       *Config
	accountId string

	wallet *eth.HDWallet

	confirms  uint64
	blockTime time.Duration

	tokens *supportedTokens
	pairs  *supportedPairs

	tt *txTracker
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

	if cfg.ConfirmBlocks == 0 {
		cfg.ConfirmBlocks = 1
	}

	if cfg.TokensFile == "" {
		cfg.TokensFile = "./tokens.json"
	}

	if cfg.Mnemonic == "" {
		cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}

	cfg.Factory = factory
	cfg.Router = routerV2

	ex := &dex{
		mux: &sync.Mutex{},

		accountId: accountId(cfg.Mnemonic),
		cfg:       cfg,

		confirms:  cfg.ConfirmBlocks,
		blockTime: time.Duration(15 * time.Second),

		tokens: newSupportedTokens(),
		pairs:  newSupportedPairs(),

		rc: rc,
		v:  v,
		l:  l,

		stopCh: make(chan struct{}),
	}

	ex.tt = newTxTracker(ex)

	ex.am = newApproveManager(ex)

	if readConfig {
		ex.l.Debug(agent, fmt.Sprintf("retriving pairs from config file %s", ex.v.ConfigFileUsed()))
		acc, ok := ex.v.Get(fmt.Sprintf("%s.account_count", ex.NID())).(float64)
		if ok {
			ex.cfg.AccountCount = uint64(acc)
		}

		ex.cfg.NativeToken = ex.v.Get(fmt.Sprintf("%s.native_token", ex.NID())).(string)
		ex.cfg.TokenStandard = ex.v.Get(fmt.Sprintf("%s.token_standard", ex.NID())).(string)

		i := ex.v.Get(fmt.Sprintf("%s.providers", ex.NID()))
		if i == nil {
			return nil, errors.New("no provider available in config file")
		}
		psi := i.(map[string]interface{})
		for _, v := range psi {
			ex.cfg.Providers = append(ex.cfg.Providers, &Provider{URL: v.(string)})
		}

		if err := ex.generalSets(); err != nil {
			return nil, err
		}

		i = ex.v.Get(fmt.Sprintf("%s.pairs", ex.NID()))
		if i != nil {
			psi := i.(map[string]interface{})
			wg := &sync.WaitGroup{}
			for _, v := range psi {
				wg.Add(1)
				go func(v interface{}) {
					defer wg.Done()
					p := v.(map[string]interface{})
					pair := &pair{}

					if p["bt"] != nil && p["qt"] != nil {
						bChainId := int64(p["bt"].(map[string]interface{})["chainid"].(float64))
						qChainId := int64(p["qt"].(map[string]interface{})["chainid"].(float64))

						if bChainId == qChainId && int64(ex.cfg.ChianId) == bChainId {
							pair.BT = Token{
								Name:     p["bt"].(map[string]interface{})["name"].(string),
								Symbol:   p["bt"].(map[string]interface{})["symbol"].(string),
								Address:  common.HexToAddress(p["bt"].(map[string]interface{})["address"].(string)),
								Decimals: int(p["bt"].(map[string]interface{})["decimals"].(float64)),
								ChainId:  bChainId,
							}

							pair.QT = Token{
								Name:     p["qt"].(map[string]interface{})["name"].(string),
								Symbol:   p["qt"].(map[string]interface{})["symbol"].(string),
								Address:  common.HexToAddress(p["qt"].(map[string]interface{})["address"].(string)),
								Decimals: int(p["qt"].(map[string]interface{})["decimals"].(float64)),
								ChainId:  qChainId,
							}

							pair, err := ex.pairWithPrice(pair.BT, pair.QT)
							if err != nil {
								ex.l.Error(agent, err.Error())
								return
							}

							ex.pairs.add(*pair)
							ex.tokens.add(pair.BT, pair.QT)
							ex.l.Debug(agent, fmt.Sprintf("pair %s added", pair.String()))
						}
					}
				}(v)
			}
			wg.Wait()
		}

	} else {
		if err := ex.generalSets(); err != nil {
			return nil, err
		}
		ex.v.Set(fmt.Sprintf("%s.native_token", ex.NID()), ex.cfg.NativeToken)
		ex.v.Set(fmt.Sprintf("%s.token_standard", ex.NID()), ex.cfg.TokenStandard)
		ex.v.Set(fmt.Sprintf("%s.account_count", ex.NID()), ex.cfg.AccountCount)
		for i, p := range ex.cfg.Providers {
			ex.v.Set(fmt.Sprintf("%s.providers.%d", ex.NID(), i), p.URL)
		}
		if err := ex.v.WriteConfig(); err != nil {
			return nil, err
		}
	}

	return ex, nil
}

func (u *dex) Run(wg *sync.WaitGroup) {
	defer wg.Done()

}

func (u *dex) Type() entity.ExType {
	return entity.DEX
}
