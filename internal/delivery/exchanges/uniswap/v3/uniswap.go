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

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Config struct {
	Id              string             `json:"id,omitempty"`
	DefaultProvider string             `json:"default_provider,omitempty"`
	BackupProviders []string           `json:"backup_providers,omitempty"`
	Mnemonic        string             `json:"mnemonic,omitempty"`
	AccountCount    uint64             `json:"account_count,omitempty"`
	Accounts        []accounts.Account `json:"accounts,omitempty"`
	ConfirmBlocks   uint64             `json:"confirm_blocks,omitempty"`
	TokensFile      string             `json:"tokens_file,omitempty"`
	TokensUrl       string             `json:"tokens_url,omitempty"`
	Msg             string             `json:"msg,omitempty"`
}

type UniSwapV3 struct {
	mux *sync.Mutex

	cfg       *Config
	accountId string

	wallet             *eth.HDWallet
	provider           *Provider
	backupProvidersURL []string

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

	if cfg.DefaultProvider == "" && len(cfg.DefaultProvider) == 0 && !readConfig {
		return nil, errors.Wrap(errors.NewMesssage("default provider and backup providers cannot by empty"))
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
	if cfg.Mnemonic == "" {
		cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}

	v3 := &UniSwapV3{
		mux: &sync.Mutex{},

		provider: &Provider{
			URL: cfg.DefaultProvider,
		},

		backupProvidersURL: cfg.BackupProviders,
		accountId:          hash(hash(cfg.Mnemonic)),
		cfg:                cfg,

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

	if readConfig {
		v3.l.Debug(agent, fmt.Sprintf("retriving pairs from config file %s", v3.v.ConfigFileUsed()))
		acc, ok := v3.v.Get(fmt.Sprintf("%s.account_count", v3.NID())).(float64)
		if ok {
			v3.cfg.AccountCount = uint64(acc)
		}

		i := v3.v.Get(fmt.Sprintf("%s.providers", v3.NID()))
		if i != nil {
			psi := i.(map[string]interface{})
			for k, v := range psi {
				if k == "default" {
					v3.provider.URL = v.(string)
					continue
				}
				v3.backupProvidersURL = append(v3.backupProvidersURL, v.(string))

			}
		}

		if err := v3.generalSets(); err != nil {
			return nil, err
		}

		i = v3.v.Get(fmt.Sprintf("%s.pairs", v3.NID()))
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

							pair, err := v3.highestLiquidPool(pair.BT, pair.QT)
							if err != nil {
								v3.l.Error(agent, err.Error())
								return
							}

							v3.pairs.add(*pair)
							v3.tokens.add(pair.BT, pair.QT)
							v3.l.Debug(agent, fmt.Sprintf("pair %s added", pair.String()))
						}
					}
				}(v)
			}
			wg.Wait()
		}

	} else {
		if err := v3.generalSets(); err != nil {
			return nil, err
		}
		v3.v.Set(fmt.Sprintf("%s.providers.default", v3.NID()), v3.provider.URL)
		v3.v.Set(fmt.Sprintf("%s.account_count", v3.NID()), v3.cfg.AccountCount)
		for i, url := range v3.backupProvidersURL {
			v3.v.Set(fmt.Sprintf("%s.providers.%d", v3.NID(), i), url)
		}
		if err := v3.v.WriteConfig(); err != nil {
			return nil, err
		}
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

func (u *UniSwapV3) pingProvider() error {
	agent := u.agent("pingProvider")
	_, err := u.provider.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(errors.Op(agent), err)
	}
	return nil
}

func (u *UniSwapV3) Type() entity.ExType {
	return entity.DEX
}
