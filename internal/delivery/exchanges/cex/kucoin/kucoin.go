package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

type kucoinExchange struct {
	cfg *Configs
	mux *sync.Mutex

	readApi  *kucoin.ApiService
	writeApi *kucoin.ApiService

	cache *cache
	da    *depositAggregator
	wa    *withdrawalAggregator
	pls   *pairList

	l              logger.Logger
	supportedCoins *supportedCoins
	pairs          entity.PairsRepo
	repo           entity.OrderRepo
	pc             entity.PairConfigs
	fee            entity.FeeService

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewKucoinExchange(cfgi interface{}, pairs entity.PairsRepo, l logger.Logger, readConfig bool,
	repo entity.OrderRepo, pc entity.PairConfigs, fee entity.FeeService) (entity.Cex, error) {
	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap("NewKucoinExchange", err)
	}

	k := &kucoinExchange{
		cfg:   cfg,
		mux:   &sync.Mutex{},
		pairs: pairs,
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

		supportedCoins: newSupportedCoins(),
		repo:           repo,
		pc:             pc,
		fee:            fee,
		l:              l,

		stopCh: make(chan struct{}),
	}

	agent := k.agent("NewKucoinExchange")

	if err := k.ping(); err != nil {
		return nil, err
	}
	k.l.Debug(agent, "ping was successful")
	k.cache = newCache(k, k.l)

	k.da = newDepositAggregator(k, k.cache)
	k.wa = newWithdrawalAggregator(k, k.cache)
	k.pls = newPairList(k, k.readApi, l)

	if readConfig {
		ps := k.pairs.GetAll(k.Id())
		cs := []*entity.Token{}
		for _, p := range ps {
			if err := k.pls.support(p, true); err != nil {
				k.l.Error(agent, err.Error())
				continue
			}
			bc := p.T1
			qc := p.T2
			cs = append(cs, bc)
			cs = append(cs, qc)
		}
		k.supportedCoins.add(cs)

	}

	k.l.Debug(agent, fmt.Sprintf("exchange %s started successfully", k.Name()))
	return k, nil
}

func (k *kucoinExchange) Run() {
	k.l.Debug(fmt.Sprintf("%s.Run", k.Name()), "started")
}

func (k *kucoinExchange) Remove() {
	op := fmt.Sprintf("%s.Stop", k.Name())
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.l.Debug(string(op), "stopped")
}

func (k *kucoinExchange) Type() entity.ExType {
	return entity.CEX
}

func (k *kucoinExchange) Id() uint {
	return k.cfg.Id
}

func (k *kucoinExchange) Name() string {
	return "kucoin"
}

func (k *kucoinExchange) ping() error {
	resp, err := k.readApi.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(k.agent("ping"), errors.NewMesssage(err.Error()))
	}

	return nil
}
