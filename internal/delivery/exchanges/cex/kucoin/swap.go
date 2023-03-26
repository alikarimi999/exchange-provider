package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) swap(o *entity.CexOrder, index int) (string, error) {
	agent := k.agent("swap")

	in := o.Routes[index].In
	out := o.Routes[index].Out

	p, ok := k.pairs.Get(k.Id(), in.String(), out.String())
	if !ok {
		return "", fmt.Errorf("swap: pair not found")
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	var side, size, funds string
	if p.T1.Equal(in) {
		size = o.Swaps[index].InAmount
		side = "sell"
	} else {
		funds = o.Swaps[index].InAmount
		side = "buy"
	}

	req, err := k.createOrderRequest(bc, qc, side, size, funds)
	if err != nil {
		k.l.Error(agent, err.Error())
		return "", err
	}

	// res, err := k.writeApi.InnerTransferV2(uuid.New().String(), in.Symbol, "main", "trade", amount)
	// if err = handleSDKErr(err, res); err != nil {
	// 	return "", errors.Wrap(err, op, errors.ErrBadRequest)
	// }

	// k.l.Debug(agent, fmt.Sprintf("%s %s transferred from main account to trade account",
	// 	amount, in.Symbol))

	res, err := k.writeApi.CreateOrder(req)
	if err = handleSDKErr(err, res); err != nil {
		k.l.Error(agent, err.Error())
		return "", err
	}

	resp := &kucoin.CreateOrderResultModel{}
	if err = res.ReadData(resp); err != nil {
		k.l.Error(agent, err.Error())
		return "", err
	}
	return resp.OrderId, nil

}

func (k *kucoinExchange) trackSwap(o *entity.CexOrder, index int) {
	agent := k.agent("TrackSap")

	s := o.Swaps[index]
	resp, err := k.readApi.Order(s.TxId)
	if err = handleSDKErr(err, resp); err != nil {
		k.l.Error(agent, err.Error())
		s.Status = entity.SwapFailed
		s.FailedDesc = err.Error()
		return
	}

	order := &kucoin.OrderModel{}
	if err = resp.ReadData(order); err != nil {
		k.l.Error(agent, err.Error())
		s.Status = entity.SwapFailed
		s.FailedDesc = err.Error()
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
}
