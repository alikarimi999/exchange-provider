package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) Id() uint {
	return k.cfg.Id
}

func (k *kucoinExchange) Name() string {
	return "kucoin"
}

func (k *kucoinExchange) Swap(o *entity.CexOrder, index int) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Swap", k.Name()))

	in := o.Routes[index].In
	out := o.Routes[index].Out

	p, ok := k.pairs.Get(k.Id(), in.String(), out.String())
	if !ok {
		return "", errors.Wrap(errors.ErrNotFound)
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	var side, size, funds, amount string
	if p.T1.Equal(in) {
		size = o.Swaps[index].InAmount
		amount = size
		side = "sell"
	} else {
		funds = o.Swaps[index].InAmount
		amount = funds
		side = "buy"
	}

	req, err := k.createOrderRequest(bc, qc, side, size, funds)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	res, err := k.writeApi.InnerTransferV2(uuid.New().String(), in.Symbol, "main", "trade", amount)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from main account to trade account",
		amount, in.Symbol))

	// create order, after transfer is done
	res, err = k.writeApi.CreateOrder(req)
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
	op := errors.Op(fmt.Sprintf("%s.TrackSap", k.Name()))

	s := o.Swaps[index]
	resp, err := k.readApi.Order(s.TxId)
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

func (k *kucoinExchange) ping() error {
	op := errors.Op(fmt.Sprintf("%s.ping", k.Name()))

	resp, err := k.readApi.Accounts("", "")
	if err = handleSDKErr(err, resp); err != nil {
		return errors.Wrap(op, errors.NewMesssage(err.Error()))
	}

	return nil
}

func (k *kucoinExchange) trackDeposit(o *entity.CexOrder, done chan<- struct{},
	proccessed <-chan bool) {
	d := o.Deposit
	c, err := k.supportedCoins.get(d.Symbol, d.Standard)
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
	o.UpdatedAt = time.Now().Unix()
}

func (k *kucoinExchange) SetDepositddress(o *entity.CexOrder) error {
	kc, err := k.supportedCoins.get(o.Deposit.Symbol, o.Deposit.Standard)
	if err != nil {
		return err
	}

	o.Deposit.Address.Addr = kc.Address
	o.Deposit.Address.Tag = kc.Tag
	return nil
}
