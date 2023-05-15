package kucoin

import (
	"exchange-provider/internal/entity"
)

func (k *kucoinExchange) minAndMax(p *entity.Pair) error {
	price, err := k.price(p)
	if err != nil {
		return err
	}
	spread, err := k.spread(0, p, price)
	if err != nil {
		return err
	}

	t1 := p.T1
	t2 := p.T2

	amountOut := t2.ET.(*Token).MinWithdrawalFee + t2.ET.(*Token).MinWithdrawalSize
	amountOut = amountOut + (amountOut * 0.2)
	amountOut = amountOut + (amountOut * p.FeeRate2)
	efa, err := k.exchangeFeeAmount(t2, p)
	if err != nil {
		return err
	}
	amountOut = amountOut + efa
	p.T1.Min = amountOut / (price - (price * spread))

	amountOut = t1.ET.(*Token).MinWithdrawalFee + t1.ET.(*Token).MinWithdrawalSize
	amountOut = amountOut + (amountOut * 0.2)
	amountOut = amountOut + (amountOut * p.FeeRate1)
	efa, err = k.exchangeFeeAmount(t1, p)
	if err != nil {
		return err
	}
	amountOut = amountOut + efa
	p.T2.Min = amountOut * (price + (price * spread))
	return nil
}
