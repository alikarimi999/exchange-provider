package kucoin

import (
	"fmt"
	"order_service/internal/delivery/exchanges/kucoin/dto"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
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
	ot   *orderTracker
	wt   *withdrawalTracker
	wa   *withdrawalAggregator
	l    logger.Logger
	adds []*kucoinAdress
}

func NewKucoinExchange(cfg *Configs, rc *redis.Client, l logger.Logger) *kucoinExchange {
	const op = errors.Op("Kucoin-Exchange-Service.NewKucoinExchange")
	k := &kucoinExchange{
		cfg: cfg,
		l:   l,
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

func (k *kucoinExchange) ID() string {
	return "kucoin"
}

func (k *kucoinExchange) Exchange(from, to entity.Coin, vol string) (string, error) {
	const op = errors.Op("Kucoin.Exchange")

	oDTO, err := dto.CreateOrderRequest(from, to, vol)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	k.l.Debug(string(op), fmt.Sprintf("kucoin opening order request: %+v", oDTO))
	res, err := k.api.CreateOrder((*kucoin.CreateOrderModel)(oDTO))
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}

	if res.Code != "200000" {
		return "", errors.Wrap(errors.New(string(res.Message)), op, errors.ErrInternal)
	}

	resp := &kucoin.CreateOrderResultModel{}

	if err = res.ReadData(resp); err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}
	return resp.OrderId, nil

}

func (k *kucoinExchange) Withdrawal(coin entity.Coin, addr, vol string) (string, error) {
	const op = errors.Op("Kucoin.Withdrawal")

	opts, err := dto.Options(coin)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	// first transfer from trade account to main account
	res, err := k.api.InnerTransferV2(uuid.New().String(), coin.Symbol, "trade", "main", vol)
	if err != nil || res.Code != "200000" {
		return "", errors.Wrap(errors.New(fmt.Sprintf("InnerTransfer   %s:%s:%s", res.Message, res.Code, err)), op, errors.ErrInternal)
	}

	res, err = k.api.ApplyWithdrawal(coin.Symbol, addr, vol, opts)
	if err != nil || res.Code != "200000" {
		return "", errors.Wrap(errors.New(fmt.Sprintf("ApplyWithdrawal   %s:%s:%s", res.Message, res.Code, err)), op, errors.ErrInternal)
	}
	w := &kucoin.ApplyWithdrawalResultModel{}
	if err = res.ReadData(w); err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}
	return w.WithdrawalId, nil
}

func (k *kucoinExchange) TrackOrder(o *entity.ExchangeOrder, done chan<- struct{},
	err chan<- error) {

	feed := &trackerFedd{
		eo:   o,
		done: done,
		err:  err,
	}

	k.ot.track(feed)
	return
}

func (k *kucoinExchange) TrackWithdrawal(w *entity.Withdrawal, done chan<- struct{},
	err chan<- error, proccessedCh <-chan bool) error {

	feed := &wtFeed{
		w:            w,
		done:         done,
		err:          err,
		proccessedCh: proccessedCh,
	}

	k.wt.track(feed)
	return nil
}

func (k *kucoinExchange) ping() error {
	const op = errors.Op("Kucoin.Ping")

	resp, err := k.api.Accounts("", "")
	if err != nil || resp.Code != "200000" {
		return errors.Wrap(errors.New(fmt.Sprintf("%s:%s:%s", resp.Message, resp.Code, err)), op, errors.ErrInternal)
	}

	return nil
}
