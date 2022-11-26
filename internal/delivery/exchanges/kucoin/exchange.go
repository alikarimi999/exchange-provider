package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

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

func (k *kucoinExchange) Exchange(o *entity.Order, index int) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Exchange", k.NID()))

	in := o.Routes[index].Input
	out := o.Routes[index].Output

	p, err := k.exchangePairs.get(in, out)
	if err != nil {
		return "", err
	}

	var side, size, funds string
	if p.BC.CoinId == in.CoinId && p.QC.ChainId == in.ChainId {
		size = o.Swaps[index].InAmount
		side = "sell"
	} else {
		funds = o.Swaps[index].InAmount
		side = "buy"
	}

	req, err := k.createOrderRequest(p, side, size, funds)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	// transfer from main account to trade account
	// if it's a buy order, we transfer the qoute coin from main account to trade account
	// if it's a sell order, we transfer the base coin from main account to trade account

	res, err := k.api.InnerTransferV2(uuid.New().String(), in.CoinId, "main", "trade", req.Funds)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}
	switch req.Side {
	case "buy":
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Funds, in.CoinId))
	case "sell":
		k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account", req.Size, in.CoinId))
	}

	// create order, after transfer is done
	res, err = k.api.CreateOrder(req)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	resp := &kucoin.CreateOrderResultModel{}

	if err = res.ReadData(resp); err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}
	return resp.OrderId, nil

}

func (k *kucoinExchange) TrackExchangeOrder(o *entity.Order, index int, done chan<- struct{}, p <-chan bool) {
	op := errors.Op(fmt.Sprintf("%s.TrackExchangeOrder", k.NID()))

	s := o.Swaps[index]
	resp, err := k.api.Order(s.ExId)
	if err = handleSDKErr(err, resp); err != nil {
		k.l.Error(string(op), err.Error())
		s.Status = entity.ExOrderFailed
		s.FailedDesc = err.Error()
		done <- struct{}{}
		<-p
		return
	}

	order := &kucoin.OrderModel{}
	if err = resp.ReadData(order); err != nil {
		k.l.Error(string(op), err.Error())
		s.Status = entity.ExOrderFailed
		s.FailedDesc = err.Error()
		done <- struct{}{}
		<-p
		return
	}

	s.InAmount = order.DealFunds
	s.OutAmount = order.DealSize

	if order.Side == "sell" {
		s.OutAmount = order.DealFunds
	} else {
		s.OutAmount = order.DealSize
	}

	s.Fee = order.Fee
	s.FeeCurrency = order.FeeCurrency
	s.Status = entity.ExOrderSucceed
	done <- struct{}{}
	<-p

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
