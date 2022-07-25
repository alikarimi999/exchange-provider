package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) ID() string {
	return "kucoin"
}

func (k *kucoinExchange) Exchange(from, to *entity.Coin, vol string) (string, error) {
	const op = errors.Op("Kucoin.Exchange")

	oDTO, err := k.createOrderRequest(from, to, vol)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	k.l.Debug(string(op), fmt.Sprintf("kucoin opening order request: %+v", oDTO))
	res, err := k.api.CreateOrder((*kucoin.CreateOrderModel)(oDTO))
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	resp := &kucoin.CreateOrderResultModel{}

	if err = res.ReadData(resp); err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}
	return resp.OrderId, nil

}

func (k *kucoinExchange) Withdrawal(coin *entity.Coin, addr, vol string) (string, error) {
	const op = errors.Op("Kucoin.Withdrawal")

	opts, err := k.withdrawalOpts(coin)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	// first transfer from trade account to main account
	res, err := k.api.InnerTransferV2(uuid.New().String(), coin.Id, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	res, err = k.api.ApplyWithdrawal(coin.Id, addr, vol, opts)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
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
	const op = errors.Op("Kucoin.ping")

	resp, err := k.api.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}