package kucoin

import (
	"exchange-provider/internal/entity"
	"math/big"
	"strconv"
)

func (k *kucoinExchange) applySpreadAndFee(ord *entity.CexOrder, route *entity.Route, p *entity.Pair) {
	total, _ := strconv.ParseFloat(ord.Swaps[0].OutAmount, 64)
	fee := total * (p.FeeRate / 100)
	ord.Withdrawal.Volume = big.NewFloat(total - fee).String()
	ord.Fee = big.NewFloat(fee).String()
	ord.FeeCurrency = route.Out.String()
	ord.Withdrawal.Token = route.Out

}
