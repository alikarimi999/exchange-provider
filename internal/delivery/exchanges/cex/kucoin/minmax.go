package kucoin

import (
	"exchange-provider/internal/entity"
	"math/big"
	"strconv"
)

func (k *exchange) minAndMax(p *entity.Pair) error {
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
	amountOut = amountOut + (amountOut * p.FeeRate2)
	efa, _, err := k.exchangeFeeAmount(t2, p)
	if err != nil {
		return err
	}
	amountOut = amountOut + efa
	min0 := amountOut / (price - (price * spread))

	amountOut = t1.ET.(*Token).MinWithdrawalFee + t1.ET.(*Token).MinWithdrawalSize
	amountOut = amountOut + (amountOut * p.FeeRate1)
	efa, _, err = k.exchangeFeeAmount(t1, p)
	if err != nil {
		return err
	}
	amountOut = amountOut + efa
	min1 := amountOut * (price + (price * spread))
	p.T1.Min, p.T2.Min = min(p, min0, min1)
	return nil
}

func min(p *entity.Pair, min0, min1 float64) (float64, float64) {
	m0, _ := strconv.ParseFloat(trim(big.NewFloat(min0+(min0*0.5)).Text('f', 12), p.T1.ET.(*Token).OrderPrecision), 64)
	m1, _ := strconv.ParseFloat(trim(big.NewFloat(min1+(min1*0.5)).Text('f', 12), p.T2.ET.(*Token).OrderPrecision), 64)
	return m0, m1
}
