package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) Id() string {
	return "kucoin"
}

func (k *kucoinExchange) Name() string {
	return "kucoin"
}

func (k *kucoinExchange) Swap(o *entity.CexOrder, index int) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Swap", k.Id()))

	in := o.Routes[index].In
	out := o.Routes[index].Out

	p, err := k.exchangePairs.get(in, out)
	if err != nil {
		return "", err
	}

	var side, size, funds, amount string
	if p.BC.TokenId == in.TokenId && string(p.QC.ChainId) == in.ChainId {
		size = o.Swaps[index].InAmount
		amount = size
		side = "sell"
	} else {
		funds = o.Swaps[index].InAmount
		amount = funds
		side = "buy"
	}

	req, err := k.createOrderRequest(p, side, size, funds)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	res, err := k.api.InnerTransferV2(uuid.New().String(), in.TokenId, "main", "trade", amount)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account",
		amount, in.TokenId))

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

func (k *kucoinExchange) TrackSwap(o *entity.CexOrder, index int, done chan<- struct{}, p <-chan bool) {
	op := errors.Op(fmt.Sprintf("%s.TrackSap", k.Id()))

	s := o.Swaps[index]
	resp, err := k.api.Order(s.TxId)
	if err = handleSDKErr(err, resp); err != nil {
		k.l.Error(string(op), err.Error())
		s.Status = entity.SwapFailed
		s.FailedDesc = err.Error()
		done <- struct{}{}
		<-p
		return
	}

	order := &kucoin.OrderModel{}
	if err = resp.ReadData(order); err != nil {
		k.l.Error(string(op), err.Error())
		s.Status = entity.SwapFailed
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
	s.Status = entity.SwapSucceed
	done <- struct{}{}
	<-p

}

func (k *kucoinExchange) TrackWithdrawal(o *entity.CexOrder, done chan<- struct{},
	proccessedCh <-chan bool) {

	feed := &wtFeed{
		w:            o.Withdrawal,
		done:         done,
		proccessedCh: proccessedCh,
	}

	k.wt.track(feed)
}

func (k *kucoinExchange) ping() error {
	op := errors.Op(fmt.Sprintf("%s.ping", k.Id()))

	resp, err := k.api.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(op, errors.NewMesssage(err.Error()))
	}

	return nil
}

func (k *kucoinExchange) TrackDeposit(o *entity.CexOrder, done chan<- struct{},
	proccessed <-chan bool) {
	d := o.Deposit
	c, err := k.supportedCoins.get(d.TokenId, d.ChainId)
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

	k.dt.track(f)
}

func (k *kucoinExchange) GetAddress(c *entity.Token) (*entity.Address, error) {
	kc, err := k.supportedCoins.get(c.TokenId, c.ChainId)
	if err != nil {
		return nil, err
	}

	return &entity.Address{
		Addr: kc.address,
		Tag:  kc.tag,
	}, nil
}
