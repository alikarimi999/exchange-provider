package dex

import (
	pv2 "exchange-provider/internal/delivery/exchanges/dex/pancakeswap/v2"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	uv3 "exchange-provider/internal/delivery/exchanges/dex/uniswap/v3"

	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type dex struct {
	mux *sync.Mutex

	cfg       *Config
	accountId string

	wallet *eth.HDWallet

	types.Dex

	confirms uint64

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

func NewDEX(cfg *Config, rc *redis.Client, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.Exchange, error) {

	agent := "NewDEX"

	if err := cfg.Validate(readConfig); err != nil {
		return nil, err
	}

	ex := &dex{
		mux: &sync.Mutex{},

		accountId: accountId(cfg.Mnemonic),
		cfg:       cfg,

		confirms: 1,

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
		ex.l.Debug(agent, fmt.Sprintf("retriving `%s` data", ex.NID()))
		acc, ok := ex.v.Get(fmt.Sprintf("%s.account_count", ex.NID())).(float64)
		if ok {
			ex.cfg.AccountCount = uint64(acc)
		}

		ex.cfg.NativeToken = ex.v.Get(fmt.Sprintf("%s.native_token", ex.NID())).(string)
		ex.cfg.TokenStandard = ex.v.Get(fmt.Sprintf("%s.token_standard", ex.NID())).(string)
		ex.cfg.Factory = common.HexToAddress(ex.v.Get(fmt.Sprintf("%s.factory", ex.NID())).(string))
		ex.cfg.Router = common.HexToAddress(ex.v.Get(fmt.Sprintf("%s.router", ex.NID())).(string))
		ex.cfg.TokensFile = ex.v.Get(fmt.Sprintf("%s.tokens_file", ex.NID())).(string)
		ex.cfg.BlockTime = time.Duration(ex.v.Get(fmt.Sprintf("%s.block_time", ex.NID())).(float64))

		i := ex.v.Get(fmt.Sprintf("%s.providers", ex.NID()))
		if i == nil {
			return nil, errors.New("no provider available in config file")
		}
		psi := i.(map[string]interface{})
		for _, v := range psi {
			ex.cfg.Providers = append(ex.cfg.Providers, &types.Provider{URL: v.(string)})
		}

		if err := ex.generalSets(); err != nil {
			return nil, err
		}

		if err := ex.setDEX(); err != nil {
			return nil, err
		}

		ps := ex.v.GetStringSlice(fmt.Sprintf("%s.pairs", ex.NID()))
		wg := &sync.WaitGroup{}
		for _, v := range ps {
			p := strings.Split(v, types.Delimiter)
			if len(p) == 2 {
				wg.Add(1)
				go func(bt, qt string) {
					defer wg.Done()
					if err := ex.addPair(bt, qt); err != nil {
						ex.l.Error(agent, err.Error())
						return
					}
					ex.l.Debug(agent, fmt.Sprintf("pair %s added", v))
				}(p[0], p[1])
			}
		}
		wg.Wait()
	} else {

		if err := ex.generalSets(); err != nil {
			return nil, err
		}

		if err := ex.setDEX(); err != nil {
			return nil, err
		}

		ex.v.Set(fmt.Sprintf("%s.factory", ex.NID()), ex.cfg.Factory)
		ex.v.Set(fmt.Sprintf("%s.router", ex.NID()), ex.cfg.Router)
		ex.v.Set(fmt.Sprintf("%s.native_token", ex.NID()), ex.cfg.NativeToken)
		ex.v.Set(fmt.Sprintf("%s.token_standard", ex.NID()), ex.cfg.TokenStandard)
		ex.v.Set(fmt.Sprintf("%s.account_count", ex.NID()), ex.cfg.AccountCount)
		ex.v.Set(fmt.Sprintf("%s.tokens_file", ex.NID()), ex.cfg.TokensFile)
		ex.v.Set(fmt.Sprintf("%s.block_time", ex.NID()), ex.cfg.BlockTime)

		for i, p := range ex.cfg.Providers {
			ex.v.Set(fmt.Sprintf("%s.providers.%d", ex.NID(), i), p.URL)
		}
		if err := ex.v.WriteConfig(); err != nil {
			return nil, err
		}

	}

	return ex, nil
}

func (d *dex) setDEX() error {
	switch d.cfg.Name {
	case "uniswapv3":
		u, err := uv3.NewUniSwapV3(d.NID(), d.cfg.Providers, d.cfg.Factory, d.cfg.Router, d.wallet, d.l)
		if err != nil {
			return err
		}
		d.Dex = u
		return nil
	case "panckakeswapv2":
		u, err := pv2.NewPanckakeswapV2(d.NID(), d.wallet, d.cfg.Router, d.cfg.Providers, d.l)
		if err != nil {
			return err
		}
		d.Dex = u
		return nil
	default:
		return fmt.Errorf("'%s' unknown exchange name", d.cfg.Name)
	}

}
