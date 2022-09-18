package kucoin

import (
	"fmt"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) NID() string {
	k.mux.Lock()
	defer k.mux.Unlock()
	return fmt.Sprintf("%s-%s", k.Name(), k.accountId)
}

func (k *kucoinExchange) Name() string {
	return "kucoin"
}

func (k *kucoinExchange) AccountId() string {
	k.mux.Lock()
	defer k.mux.Unlock()
	return k.accountId
}

func (k *kucoinExchange) Exchange(o *entity.UserOrder, size, funds string) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Exchange", k.NID()))

	bc := o.BC
	qc := o.QC
	side := o.Side
	req, err := k.createOrderRequest(bc, qc, side, size, funds)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	// transfer from main account to trade account
	// if it's a buy order, we transfer the qoute coin from main account to trade account
	// if it's a sell order, we transfer the base coin from main account to trade account
	switch req.Side {
	case "buy":
		k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from main account to trade account", req.Funds, qc.CoinId))
		res, err := k.api.InnerTransferV2(uuid.New().String(), qc.CoinId, "main", "trade", req.Funds)
		if err = handleSDKErr(err, res); err != nil {
			return "", errors.Wrap(err, op, errors.ErrBadRequest)
		}
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Funds, qc.CoinId))
	case "sell":
		k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from main account to trade account", req.Size, bc.CoinId))
		res, err := k.api.InnerTransferV2(uuid.New().String(), bc.CoinId, "main", "trade", req.Size)
		if err = handleSDKErr(err, res); err != nil {
			return "", errors.Wrap(err, op, errors.ErrBadRequest)
		}
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Size, bc.CoinId))
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

func (k *kucoinExchange) Withdrawal(o *entity.UserOrder, coin *entity.Coin, a *entity.Address, vol string) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Withdrawal", k.NID()))

	opts, err := k.withdrawalOpts(coin, a.Tag)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	wc, err := k.supportedCoins.get(coin.CoinId, coin.ChainId)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	vol = trim(vol, wc.WithdrawalPrecision)

	// first transfer from trade account to main account
	k.l.Debug(string(op), fmt.Sprintf("transferring %s `%s` from trade account to main account", vol, coin.CoinId))
	res, err := k.api.InnerTransferV2(uuid.New().String(), coin.CoinId, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from trade account to main account", vol, coin.CoinId))

	// then withdraw from main account
	k.l.Debug(string(op), fmt.Sprintf("withdrawing %s `%s` from main account", vol, coin.CoinId))
	res, err = k.api.ApplyWithdrawal(coin.CoinId, a.Addr, vol, opts)
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

func (k *kucoinExchange) TrackExchangeOrder(o *entity.UserOrder, done chan<- struct{}, p <-chan bool) {
	feed := &trackerFedd{
		eo:   o.ExchangeOrder,
		done: done,
		pCh:  p,
	}

	k.ot.track(feed)

}

func (k *kucoinExchange) TrackWithdrawal(w *entity.Withdrawal, done chan<- struct{},
	proccessedCh <-chan bool) {

	feed := &wtFeed{
		w:            w,
		done:         done,
		proccessedCh: proccessedCh,
	}

	k.wt.track(feed)
}

func (k *kucoinExchange) ping() error {
	op := errors.Op(fmt.Sprintf("%s.ping", k.NID()))

	resp, err := k.api.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(op, errors.NewMesssage(err.Error()))
	}

	return nil
}

func (k *kucoinExchange) TrackDeposit(d *entity.Deposit, done chan<- struct{},
	proccessed <-chan bool) {

	c, err := k.supportedCoins.get(d.CoinId, d.ChainId)
	if err != nil {
		d.Status = entity.DepositFailed
		d.FailedDesc = err.Error()
		done <- struct{}{}
		<-proccessed
		return
	}

	f := &dtFeed{
		d:         d,
		blockTime: c.BlockTime,
		confirms:  c.ConfirmBlocks,
		done:      done,
		pCh:       proccessed,
	}

	k.dt.fCh <- f
}

func (k *kucoinExchange) GetAddress(c *entity.Coin) (*entity.Address, error) {
	kc, err := k.supportedCoins.get(c.CoinId, c.ChainId)
	if err != nil {
		return nil, err
	}

	return &entity.Address{
		Addr: kc.address,
		Tag:  kc.tag,
	}, nil
}
