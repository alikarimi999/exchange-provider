package kucoin

import (
	"exchange-provider/internal/entity"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) swap(o *entity.CexOrder, p *entity.Pair) (string, error) {
	agent := k.agent("swap")

	index := 0

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)

	var side, size, funds string
	if p.T1.Equal(o.Routes[index].In) {
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
	agent := k.agent("trackSap")

	time.Sleep(3 * time.Second)
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

	if order.Side == "sell" {
		s.InAmount = order.DealSize
		s.OutAmount = order.DealFunds
	} else {
		s.InAmount = order.DealFunds
		s.OutAmount = order.DealSize
	}

	s.Fee = order.Fee
	s.FeeCurrency = order.FeeCurrency
	s.Status = entity.SwapSucceed
}
