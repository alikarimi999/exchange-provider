package dex

import (
	"exchange-provider/internal/app"
	pv2 "exchange-provider/internal/delivery/exchanges/dex/pancakeswap/v2"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	uv3 "exchange-provider/internal/delivery/exchanges/dex/uniswap/v3"
	"exchange-provider/internal/delivery/exchanges/dex/utils"

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

	cfg *Config

	wallet *eth.HDWallet
	ws     app.WalletStore
	types.Dex

	confirms uint64

	tokens *supportedTokens
	pairs  *supportedPairs

	tt *utils.TxTracker
	am *utils.ApproveManager
	rc *redis.Client
	v  *viper.Viper
	l  logger.Logger

	stopCh    chan struct{}
	stoppedAt time.Time
}

func NewDEX(cfg *Config, ws app.WalletStore, rc *redis.Client, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.Cex, error) {

	agent := "NewDEX"

	if err := cfg.Validate(readConfig); err != nil {
		return nil, err
	}

	ex := &dex{
		mux: &sync.Mutex{},

		cfg:      cfg,
		ws:       ws,
		confirms: 1,

		tokens: newSupportedTokens(),
		pairs:  newSupportedPairs(),

		rc: rc,
		v:  v,
		l:  l,

		stopCh: make(chan struct{}),
	}

	if readConfig {
		ex.l.Debug(agent, fmt.Sprintf("retriving `%s` data", ex.Id()))
		acc, ok := ex.v.Get(fmt.Sprintf("%s.account_count", ex.Id())).(float64)
		if ok {
			ex.cfg.AccountCount = uint64(acc)
		}

		ex.cfg.NativeToken = ex.v.GetString(fmt.Sprintf("%s.native_token", ex.Id()))
		ex.cfg.TokenStandard = ex.v.GetString(fmt.Sprintf("%s.token_standard", ex.Id()))
		ex.cfg.Factory = common.HexToAddress(ex.v.GetString(fmt.Sprintf("%s.factory", ex.Id())))
		ex.cfg.Router = common.HexToAddress(ex.v.GetString(fmt.Sprintf("%s.router", ex.Id())))
		ex.cfg.TokensFile = ex.v.GetString(fmt.Sprintf("%s.tokens_file", ex.Id()))
		ex.cfg.BlockTime = time.Duration(ex.v.GetFloat64(fmt.Sprintf("%s.block_time", ex.Id())))

		ex.tt = utils.NewTxTracker(ex.Id(), ex.cfg.BlockTime, ex.confirms, l)

		i := ex.v.Get(fmt.Sprintf("%s.providers", ex.Id()))
		if i == nil {
			return nil, errors.New("no provider available in config file")
		}

		psi := i.(map[string]interface{})
		for _, v := range psi {
			ex.cfg.Providers = append(ex.cfg.Providers, &types.EthProvider{URL: v.(string)})
		}
		ex.am = utils.NewApproveManager(int64(ex.cfg.ChainId), ex.tt, ex.wallet, ex.l, ex.cfg.Providers)

		if err := ex.checkProviders(); err != nil {
			return nil, err
		}

		if err := ex.setupWallet(); err != nil {
			return nil, err
		}

		if err := ex.setDEX(); err != nil {
			return nil, err
		}

		ps := ex.v.GetStringSlice(fmt.Sprintf("%s.pairs", ex.Id()))
		wg := &sync.WaitGroup{}
		for _, v := range ps {
			p := strings.Split(v, types.Delimiter)
			if len(p) == 2 {
				wg.Add(1)
				go func(t1, t2 string) {
					defer wg.Done()
					if err := ex.addPair(t1, t2); err != nil {
						ex.l.Error(agent, err.Error())
						return
					}
					ex.l.Debug(agent, fmt.Sprintf("pair %s added", t1+"/"+t2))
				}(p[0], p[1])
			}
		}
		wg.Wait()
	} else {

		ex.v.Set(fmt.Sprintf("%s.factory", ex.Id()), ex.cfg.Factory)
		ex.v.Set(fmt.Sprintf("%s.router", ex.Id()), ex.cfg.Router)
		ex.v.Set(fmt.Sprintf("%s.native_token", ex.Id()), ex.cfg.NativeToken)
		ex.v.Set(fmt.Sprintf("%s.token_standard", ex.Id()), ex.cfg.TokenStandard)
		ex.v.Set(fmt.Sprintf("%s.account_count", ex.Id()), ex.cfg.AccountCount)
		ex.v.Set(fmt.Sprintf("%s.tokens_file", ex.Id()), ex.cfg.TokensFile)
		ex.v.Set(fmt.Sprintf("%s.block_time", ex.Id()), ex.cfg.BlockTime)

		ex.tt = utils.NewTxTracker(ex.Id(), ex.cfg.BlockTime, ex.confirms, l)

		for i, p := range ex.cfg.Providers {
			ex.v.Set(fmt.Sprintf("%s.providers.%d", ex.Id(), i), p.URL)
		}

		if err := ex.checkProviders(); err != nil {
			return nil, err
		}

		if err := ex.setupWallet(); err != nil {
			return nil, err
		}

		if err := ex.setDEX(); err != nil {
			return nil, err
		}

		ex.am = utils.NewApproveManager(int64(ex.cfg.ChainId), ex.tt, ex.wallet, ex.l, ex.cfg.Providers)

		if err := ex.v.WriteConfig(); err != nil {
			return nil, err
		}

	}

	return ex, nil
}

func (d *dex) setDEX() error {
	switch d.cfg.Name {
	case "uniswapv3":
		dex, err := uv3.NewUniSwapV3(d.Id(), d.cfg.NativeToken, d.cfg.Providers,
			d.cfg.Factory, d.cfg.Router, d.wallet, d.tt, d.l)
		if err != nil {
			return err
		}
		d.Dex = dex
		return nil
	case "panckakeswapv2":
		dex, err := pv2.NewPanckakeswapV2(d.Id(), d.cfg.NativeToken, d.wallet, d.tt, d.cfg.Router,
			d.cfg.Providers, d.l)
		if err != nil {
			return err
		}
		d.Dex = dex
		return nil
	default:
		return fmt.Errorf("'%s' unknown exchange name", d.cfg.Name)
	}

}
