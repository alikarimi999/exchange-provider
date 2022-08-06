package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"
	"time"

	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/go-redis/redis/v9"
)

type kucoinAdress struct {
	Address  string
	Chain    string
	Currency string
}

type Configs struct {
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	ApiPassphrase string `json:"api_passphrase"`
	ApiVersion    string
	ApiUrl        string
	// WsTopics      []string
}

// kucoinExchange is a concrete implementation of entity.Exchange interface.
type kucoinExchange struct {
	cfg       *Configs
	mux       *sync.Mutex
	accountId string

	api *kucoin.ApiService
	// ws   *webSocket
	ot  *orderTracker
	wt  *withdrawalTracker
	wa  *withdrawalAggregator
	pls *pairList
	l   logger.Logger

	exchangePairs   *exPairs
	withdrawalCoins *withdrawalCoins

	stopCh   chan struct{}
	stopedAt time.Time
}

func NewKucoinExchange(cfgi interface{}) (*kucoinExchange, error) {
	const op = errors.Op("Kucoin-Exchange.NewKucoinExchange")

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	return &kucoinExchange{
		cfg:       cfg,
		mux:       &sync.Mutex{},
		accountId: hash(cfg.ApiKey, cfg.ApiSecret, cfg.ApiPassphrase),
	}, nil
}

func (k *kucoinExchange) Setup(rc *redis.Client, l logger.Logger) (entity.Exchange, error) {
	const op = errors.Op("Kucoin-Exchange.Setup")

	k.stopCh = make(chan struct{})
	k.l = l
	k.exchangePairs = newExPairs()
	k.withdrawalCoins = newWithdrawalCoins()

	k.api = kucoin.NewApiService(
		kucoin.ApiBaseURIOption(k.cfg.ApiUrl),
		kucoin.ApiKeyOption(k.cfg.ApiKey),
		kucoin.ApiSecretOption(k.cfg.ApiSecret),
		kucoin.ApiPassPhraseOption(k.cfg.ApiPassphrase),
		kucoin.ApiKeyVersionOption(k.cfg.ApiVersion),
	)
	if err := k.ping(); err != nil {
		return nil, err
	}

	k.l.Debug(string(op), fmt.Sprintf("ping was successful"))

	k.ot = newOrderTracker(k.api, l)
	k.wt = newWithdrawalTracker(rc, l)
	k.wa = newWithdrawalAggregator(k.api, l, rc)
	k.pls = newPairList(k.api, l)
	// k.setupWebSocket(rc)

	return k, nil
}

func (k *kucoinExchange) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}
	w.Add(1)
	go k.ot.run(w, k.stopCh)
	w.Add(1)
	go k.wt.run(w, k.stopCh)
	w.Add(1)
	go k.wa.run(w, k.stopCh)

	k.l.Debug("Kucoin-Exchange.Run", "started")
	w.Wait()
}

func (k *kucoinExchange) Stop() {
	close(k.stopCh)
	k.stopedAt = time.Now()
	k.l.Debug("Kucoin-Exchange", "stopped")
}

func (k *kucoinExchange) StartAgain() (*entity.StartAgainResult, error) {
	const op = errors.Op("Kucoin-Exchange.StartAgain")
	k.stopCh = make(chan struct{})

	k.l.Debug(string(op), "starting again")
	k.l.Debug(string(op), fmt.Sprintf("stopped at %s", k.stopedAt.Format("2006-01-02 15:04:05")))

	// check if ping is successful
	if err := k.ping(); err != nil {
		return nil, errors.Wrap(string(op), err)
	}
	k.l.Debug(string(op), fmt.Sprintf("ping was successful"))

	k.l.Debug(string(op), "downloading pairs list from kucoin")

	// download pairs list from kucoin again
	if err := k.pls.download(); err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	res := &entity.StartAgainResult{}
	// check if current pairs are still supported by kucoin
	ps := k.exchangePairs.snapshot()
	k.exchangePairs.purge()
	cs := k.withdrawalCoins.snapshot()
	k.withdrawalCoins.purge()
	newPs := []*pair{}
	newCs := map[string]*withdrawalCoin{}
	for _, p := range ps {
		pe := p.toEntity()
		ok, err := k.pls.support(pe)
		if err != nil {
			return nil, errors.Wrap(string(op), err)
		}

		if !ok {
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: pe,
				Err:  fmt.Errorf("pair is not supported by kucoin anymore so it will be removed"),
			})
			continue
		}

		if err := k.setInfos(pe); err != nil {
			res.Removed = append(res.Removed, &entity.PairsErr{
				Pair: pe,
				Err:  fmt.Errorf("retrieving infos for pair failed due to error ( %s ) so it well be removed", err.Error()),
			})
			continue
		}
		newPs = append(newPs, fromEntity(pe))
		newCs[pe.BC.CoinId+pe.BC.ChainId] = &withdrawalCoin{
			precision: cs[pe.BC.CoinId+pe.BC.ChainId].precision,
			needChain: pe.BC.SetChain,
		}

		newCs[pe.QC.CoinId+pe.QC.ChainId] = &withdrawalCoin{
			precision: cs[pe.QC.CoinId+pe.QC.ChainId].precision,
			needChain: pe.QC.SetChain,
		}

	}

	k.exchangePairs.add(newPs)
	k.withdrawalCoins.add(newCs)

	k.l.Debug(string(op), fmt.Sprintf("%d pairs were added", len(newPs)))
	k.l.Debug(string(op), fmt.Sprintf("%d pairs were removed", len(ps)-len(newPs)))
	k.l.Debug(string(op), fmt.Sprintf("exchange %s started again successfully", k.AccountId()))

	return res, nil
}
