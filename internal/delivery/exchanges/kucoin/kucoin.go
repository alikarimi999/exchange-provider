package kucoin

import (
	"order_service/internal/entity"
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
	ApiKey        string `json:"api_key"`
	ApiSecret     string `json:"api_secret"`
	ApiPassphrase string `json:"api_passphrase"`
	ApiVersion    string
	ApiUrl        string
	// WsTopics      []string
}

// kucoinExchange is a concrete implementation of entity.Exchange interface.
type kucoinExchange struct {
	cfg *Configs
	api *kucoin.ApiService
	// ws   *webSocket
	ot  *orderTracker
	wt  *withdrawalTracker
	wa  *withdrawalAggregator
	pls *pairList
	l   logger.Logger

	exchangePairs   *exPairs
	withdrawalCoins *withdrawalCoins
}

func NewKucoinExchange() *kucoinExchange {
	return &kucoinExchange{}
}

func (k *kucoinExchange) Setup(cfgi interface{}, rc *redis.Client, l logger.Logger) (entity.Exchange, error) {
	const op = errors.Op("Kucoin-Exchange.Setup")

	cfg, err := validateConfigs(cfgi)
	if err != nil {
		return nil, errors.Wrap(string(op), err)
	}

	k = &kucoinExchange{
		cfg:             cfg,
		l:               l,
		exchangePairs:   newExPairs(),
		withdrawalCoins: newWithdrawalCoins(),
	}
	k.api = kucoin.NewApiService(
		kucoin.ApiBaseURIOption(cfg.ApiUrl),
		kucoin.ApiKeyOption(cfg.ApiKey),
		kucoin.ApiSecretOption(cfg.ApiSecret),
		kucoin.ApiPassPhraseOption(cfg.ApiPassphrase),
		kucoin.ApiKeyVersionOption(cfg.ApiVersion),
	)
	if err := k.ping(); err != nil {
		return nil, err
	}

	k.l.Debug(string(op), "ping was successful")

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
	go k.ot.run(w)
	w.Add(1)
	go k.wt.run(w)
	w.Add(1)
	go k.wa.run(w)

	k.l.Debug("Kucoin-Exchange.Run", "started")
	w.Wait()
}
