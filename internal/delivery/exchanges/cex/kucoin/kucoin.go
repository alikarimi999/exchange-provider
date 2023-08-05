package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

type exchange struct {
	cfg *Configs

	readApi  *kucoin.ApiService
	writeApi *kucoin.ApiService

	cache *cache
	da    *depositAggregator
	wa    *withdrawalAggregator
	si    *serverInfos

	l     logger.Logger
	pairs entity.PairsRepo
	repo  entity.OrderRepo

	fee entity.FeeTable
	st  entity.SpreadTable

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewExchange(cfgi interface{}, pairs entity.PairsRepo, l logger.Logger, fromDB bool, lastUpdate time.Time,
	repo entity.OrderRepo, fee entity.FeeTable, st entity.SpreadTable) (entity.Cex, error) {

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap("NewKucoinExchange", err)
	}

	k := &exchange{
		cfg:   cfg,
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

		repo: repo,
		fee:  fee,
		st:   st,
		l:    l,

		stopCh: make(chan struct{}),
	}

	if err := k.ping(); err != nil {
		return nil, err
	}

	si, err := newServerInfos(k, k.readApi)
	if err != nil {
		return nil, err
	}
	k.si = si

	k.cache = newCache(k, k.l)

	k.da = newDepositAggregator(k, k.cache)
	k.wa = newWithdrawalAggregator(k, k.cache)

	if fromDB {
		ps := k.pairs.GetAll(k.Id())
		if len(ps) > 0 {
			for _, p := range ps {
				if err := k.setInfos(p); err != nil {
					k.pairs.Remove(k.cfg.Id, p.T1.String(), p.T2.String(), false)
					continue
				}
				if err := k.pairs.Update(k.Id(), p); err != nil {
					k.pairs.Remove(k.cfg.Id, p.T1.String(), p.T2.String(), false)
					continue
				}
			}

			if err := k.retreiveOrders(lastUpdate); err != nil {
				return nil, err
			}
		}
	}

	go k.si.run(k, k.stopCh)
	go k.da.run(k.stopCh)
	go k.wa.run(k.stopCh)
	return k, nil
}

func (ex *exchange) Remove() {
	op := fmt.Sprintf("%s.Stop", ex.NID())
	close(ex.stopCh)
	ex.stopedAt = time.Now()
	ex.l.Debug(string(op), "stopped")
}

func (ex *exchange) Type() entity.ExType {
	return entity.CEX
}

func (ex *exchange) Id() uint {
	return ex.cfg.Id
}

func (ex *exchange) Name() string {
	return "kucoin"
}

func (ex *exchange) EnableDisable(enable bool) {
	ex.cfg.Enable = enable
}
func (ex *exchange) IsEnable() bool {
	return ex.cfg.Enable
}

func (ex *exchange) NID() string {
	return fmt.Sprintf("%s-%d", ex.Name(), ex.Id())
}
func (ex *exchange) UpdateStatus(eo entity.Order) error { return nil }

func (ex *exchange) ping() error {
	resp, err := ex.readApi.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(ex.agent("ping"), errors.NewMesssage(err.Error()))
	}

	return nil
}
