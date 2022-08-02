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

func (k *kucoinExchange) Exchange(o *entity.UserOrder) (string, error) {
	const op = errors.Op("Kucoin.Exchange")

	req, err := k.createOrderRequest(o)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	// transfer from main account to trade account
	// if it's a buy order, we transfer the qoute coin from main account to trade account
	// if it's a sell order, we transfer the base coin from main account to trade account
	switch req.Side {
	case "buy":
		k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from main account to trade account", req.Funds, o.QC.CoinId))
		res, err := k.api.InnerTransferV2(uuid.New().String(), o.QC.CoinId, "main", "trade", req.Funds)
		if err = handleSDKErr(err, res); err != nil {
			return "", errors.Wrap(err, op, errors.ErrBadRequest)
		}
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Funds, o.QC.CoinId))
	case "sell":
		k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from main account to trade account", req.Size, o.BC.CoinId))
		res, err := k.api.InnerTransferV2(uuid.New().String(), o.BC.CoinId, "main", "trade", req.Size)
		if err = handleSDKErr(err, res); err != nil {
			return "", errors.Wrap(err, op, errors.ErrBadRequest)
		}
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Size, o.BC.CoinId))
	}

	// create order, after transfer is done
	k.l.Debug(string(op), fmt.Sprintf("kucoin opening order request: %+v", req))
	res, err := k.api.CreateOrder(req)
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

	wc, err := k.withdrawalCoins.get(coin.CoinId, coin.ChainId)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	vol = trim(vol, wc.precision)

	// first transfer from trade account to main account
	k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from trade account to main account", vol, coin.CoinId))
	res, err := k.api.InnerTransferV2(uuid.New().String(), coin.CoinId, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from trade account to main account", vol, coin.CoinId))

	// then withdraw from main account
	k.l.Debug(string(op), fmt.Sprintf("withdrawing %s `%s` from main account", vol, coin.CoinId))
	res, err = k.api.ApplyWithdrawal(coin.CoinId, addr, vol, opts)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s withdrawn from main account", vol, coin.CoinId))

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
