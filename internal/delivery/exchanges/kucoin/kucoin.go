package kucoin

import (
	"order_service/pkg/logger"
	"sync"

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
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
	ApiVersion    string
	ApiUrl        string
	// WsTopics      []string
}

// kucoinExchange is a concrete implementation of entity.Exchange interface.
type kucoinExchange struct {
	cfg *Configs
	api *kucoin.ApiService
	// ws   *webSocket
	ot *orderTracker
	wt *withdrawalTracker
	wa *withdrawalAggregator
	l  logger.Logger

	exchangePairs   *exchangePairs
	withdrawalCoins *withdrawalCoins
}

func NewKucoinExchange(cfg *Configs, rc *redis.Client, l logger.Logger) *kucoinExchange {
	const op = errors.Op("Kucoin-Exchange-Service.NewKucoinExchange")
	k := &kucoinExchange{
		cfg:             cfg,
		l:               l,
		exchangePairs:   newExchangePairs(),
		withdrawalCoins: newWithdrawalCoins(),
	}
	k.api = kucoin.NewApiService(
		kucoin.ApiBaseURIOption(cfg.ApiUrl),
		kucoin.ApiKeyOption(cfg.ApiKey),
		kucoin.ApiSecretOption(cfg.ApiSecret),
		kucoin.ApiPassPhraseOption(cfg.ApiPassphrase),
		kucoin.ApiKeyVersionOption(cfg.ApiVersion),
	)

	l.Debug(string(op), "kucoin: ping...")
	if err := k.ping(); err != nil {
		l.Fatal(string(op), errors.Wrap(err, op).Error())
	}
	l.Debug(string(op), "kucoin: ping ok")

	k.ot = newOrderTracker(k.api, l)
	k.wt = newWithdrawalTracker(rc, l)
	k.wa = newWithdrawalAggregator(k.api, l, rc)
	// k.setupWebSocket(rc)

	return k
}

func (k *kucoinExchange) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	w := &sync.WaitGroup{}
	w.Add(1)
	go k.ot.run(w)
	w.Add(1)
	go k.wt.run(w)
	w.Add(1)
	go k.wa.run(w)
	w.Wait()
}
