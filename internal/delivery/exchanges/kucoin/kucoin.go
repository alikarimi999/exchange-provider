package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"strings"
	"sync"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/spf13/viper"
)

type API struct {
	ApiKey        string `json:"apiKey"`
	ApiSecret     string `json:"apiSecret"`
	ApiPassphrase string `json:"apiPassphrase"`
}

type Configs struct {
	ReadApi  *API `json:"readApi,omitempty"`
	WriteApi *API `json:"writeApi,omitempty"`

	ApiVersion string
	ApiUrl     string
	Message    string
}

type kucoinExchange struct {
	cfg *Configs
	mux *sync.Mutex

	readApi  *kucoin.ApiService
	writeApi *kucoin.ApiService

	cache *cache
	wt    *withdrawalTracker
	da    *depositAggregator
	dt    *depositTracker
	wa    *withdrawalAggregator
	pls   *pairList

	v *viper.Viper
	l logger.Logger

	exchangePairs  *exPairs
	pairs          entity.PairRepo
	supportedCoins *supportedCoins

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewKucoinExchange(cfgi interface{}, pairs entity.PairRepo, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.Cex, error) {
	const op = errors.Op("Kucoin-Exchange.NewKucoinExchange")

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	k := &kucoinExchange{
		cfg: cfg,
		mux: &sync.Mutex{},

		readApi: kucoin.NewApiService(
			kucoin.ApiBaseURIOption(cfg.ApiUrl),
			kucoin.ApiKeyOption(cfg.ReadApi.ApiKey),
			kucoin.ApiSecretOption(cfg.ReadApi.ApiSecret),
			kucoin.ApiPassPhraseOption(cfg.ReadApi.ApiPassphrase),
			kucoin.ApiKeyVersionOption(cfg.ApiVersion),
		),
		writeApi: kucoin.NewApiService(
			kucoin.ApiBaseURIOption(cfg.ApiUrl),
			kucoin.ApiKeyOption(cfg.WriteApi.ApiKey),
			kucoin.ApiSecretOption(cfg.WriteApi.ApiSecret),
			kucoin.ApiPassPhraseOption(cfg.WriteApi.ApiPassphrase),
			kucoin.ApiKeyVersionOption(cfg.ApiVersion),
		),

		exchangePairs:  newExPairs(),
		supportedCoins: newSupportedCoins(),
		pairs:          pairs,
		v:              v,
		l:              l,

		stopCh: make(chan struct{}),
	}

	if err := k.ping(); err != nil {
		return nil, err
	}
	k.l.Debug(string(op), "ping was successful")
	k.cache = newCache(k, k.l)

	k.wt = newWithdrawalTracker(k, k.cache)
	k.da = newDepositAggregator(k, k.cache)
	k.dt = newDepositTracker(k, k.cache)
	k.wa = newWithdrawalAggregator(k, k.cache)
	k.pls = newPairList(k, k.readApi, l)

	if readConfig {
		k.l.Debug(string(op), fmt.Sprintf("retriving pairs from config file %s", k.v.ConfigFileUsed()))

		i := k.v.Get(fmt.Sprintf("%s.pairs", k.Id()))
		if i != nil {
			if err := k.pls.download(); err != nil {
				return nil, err
			}

			psi := i.(map[string]interface{})

			ps := make(map[string]*pair)
			for _, v := range psi {
				p := v.(map[string]interface{})
				pc := &pair{}
				if p["bc"] != nil && p["qc"] != nil {
					pc.BC = &kuToken{
						TokenId: p["bc"].(map[string]interface{})["tokenid"].(string),
						ChainId: chainId(p["bc"].(map[string]interface{})["chainid"].(string)),

						BlockTime:           time.Duration(p["bc"].(map[string]interface{})["block_time"].(float64)),
						ConfirmBlocks:       int64(p["bc"].(map[string]interface{})["confirm_blocks"].(float64)),
						WithdrawalPrecision: int(p["bc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}

					pc.QC = &kuToken{
						TokenId: p["qc"].(map[string]interface{})["tokenid"].(string),
						ChainId: chainId(p["qc"].(map[string]interface{})["chainid"].(string)),

						BlockTime:           time.Duration(p["qc"].(map[string]interface{})["block_time"].(float64)),
						ConfirmBlocks:       int64(p["qc"].(map[string]interface{})["confirm_blocks"].(float64)),
						WithdrawalPrecision: int(p["qc"].(map[string]interface{})["withdrawal_precision"].(float64)),
						needChain:           true,
					}

					ps[pc.Id()] = pc
				}
			}

			newPs := []*pair{}
			newCs := map[string]*kuToken{}
			for _, p := range ps {
				ok, _ := k.pls.support(p)
				if !ok {
					k.l.Debug(string(op), fmt.Sprintf("pair %s is not supported by kucoin anymore", p.String()))
					delete(k.v.Get(fmt.Sprintf("%s.pairs", k.Id())).(map[string]interface{}), strings.ToLower(p.Id()))
					if err := k.v.WriteConfig(); err != nil {
						k.l.Error(string(op), err.Error())
					}
					continue
				}

				if err := k.setInfos(p); err != nil {
					k.l.Error(string(op), fmt.Sprintf("failed to set infos for pair %s du to error (%s)",
						p.String(), err.Error()))
					continue
				}

				newPs = append(newPs, p)
				newCs[p.BC.TokenId+string(p.BC.ChainId)] = p.BC.snapshot()
				newCs[p.QC.TokenId+string(p.QC.ChainId)] = p.QC.snapshot()

			}

			k.exchangePairs.add(newPs...)
			k.supportedCoins.add(newCs)
			eps := []*entity.Pair{}
			for _, p := range newPs {
				ep := p.toEntity()
				ep.FeeRate = k.orderFeeRate(p)
				eps = append(eps, ep)
			}
			psPrice, err := k.Price(eps...)
			if err != nil {
				return nil, err
			}
			k.pairs.Add(k, psPrice...)

			k.l.Info(string(op), fmt.Sprintf("%d pairs loaded", len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("%d pairs couldn't be loaded", len(ps)-len(newPs)))
			k.l.Info(string(op), fmt.Sprintf("exchange %s started successfully", k.Id()))

		}
	}
	return k, nil
}

func (k *kucoinExchange) Run() {
	k.l.Debug(fmt.Sprintf("%s.Run", k.Id()), "started")
}

func (k *kucoinExchange) Remove() {
	op := fmt.Sprintf("%s.Stop", k.Id())
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.pairs.RemoveExchange(k.Id())
	k.l.Debug(string(op), "stopped")
}

func (k *kucoinExchange) Type() entity.ExType {
	return entity.CEX
}
