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

type exchange struct {
	cfg *Configs
	mux *sync.Mutex

	readApi  *kucoin.ApiService
	writeApi *kucoin.ApiService

	cache *cache
	da    *depositAggregator
	wa    *withdrawalAggregator
	pls   *pairList

	l     logger.Logger
	pairs entity.PairsRepo
	repo  entity.OrderRepo

	fee entity.FeeTable
	st  entity.SpreadTable

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewExchange(cfgi interface{}, pairs entity.PairsRepo, l logger.Logger, fromDB bool,
	repo entity.OrderRepo, fee entity.FeeTable, st entity.SpreadTable) (entity.Cex, error) {

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap("NewKucoinExchange", err)
	}

	k := &exchange{
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

		repo: repo,
		fee:  fee,
		st:   st,
		l:    l,

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

	if fromDB {
		ps := k.pairs.GetAll(k.Id())
		if len(ps) > 0 {
			if err := k.pls.downloadList(); err != nil {
				return nil, err
			}
			wg := &sync.WaitGroup{}
			waitChan := make(chan struct{}, max_conccurrent_jobs)

			for _, p := range ps {
				waitChan <- struct{}{}
				wg.Add(1)
				go func(p *entity.Pair) {
					defer func() {
						<-waitChan
						wg.Done()
					}()
					if err := k.support(p); err != nil {
						k.pairs.Remove(k.cfg.Id, p.T1.String(), p.T2.String(), false)
						k.l.Debug(agent, err.Error())
						return
					}

					if err := k.setInfos(p); err != nil {
						k.pairs.Remove(k.cfg.Id, p.T1.String(), p.T2.String(), false)
						k.l.Debug(agent, err.Error())
						return
					}
					if err := k.pairs.Update(k.Id(), p); err != nil {
						k.pairs.Remove(k.cfg.Id, p.T1.String(), p.T2.String(), false)
						k.l.Debug(agent, err.Error())
						return
					}
				}(p)
			}
			wg.Wait()
			if err := k.retreiveOrders(); err != nil {
				return nil, err
			}
		}
	}
	go k.da.run(k.stopCh)
	go k.wa.run(k.stopCh)
	k.l.Debug(agent, fmt.Sprintf("exchange '%s' started successfully", k.NID()))
	return k, nil
}

func (k *exchange) Remove() {
	op := fmt.Sprintf("%s.Stop", k.NID())
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.l.Debug(string(op), "stopped")
}

func (k *exchange) Type() entity.ExType {
	return entity.CEX
}

func (k *exchange) Id() uint {
	return k.cfg.Id
}

func (k *exchange) Name() string {
	return "kucoin"
}

func (k *exchange) EnableDisable(enable bool) {
	k.mux.Lock()
	defer k.mux.Unlock()
	k.cfg.Enable = enable
}
func (k *exchange) IsEnable() bool {
	k.mux.Lock()
	defer k.mux.Unlock()
	return k.cfg.Enable
}

func (k *exchange) NID() string {
	return fmt.Sprintf("%s-%d", k.Name(), k.Id())
}

func (k *exchange) ping() error {
	resp, err := k.readApi.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(k.agent("ping"), errors.NewMesssage(err.Error()))
	}

	return nil
}
