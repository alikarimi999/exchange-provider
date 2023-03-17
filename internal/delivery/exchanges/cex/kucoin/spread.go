package kucoin

import "exchange-provider/internal/entity"

func (k *kucoinExchange) applySpreadAndFee(ord *entity.CexOrder, route *entity.Route) {
	aVol, sVol, rate, _ := k.pc.ApplySpread(route.In, route.Out, ord.Swaps[len(ord.Swaps)-1].OutAmount)

	ord.SpreadCurrency = route.Out.String()
	ord.SpreadVol = sVol
	ord.SpreadRate = rate

	r, f, _ := k.fee.ApplyFee(ord.UserId, aVol)

	ord.Withdrawal.Token = route.Out
	ord.Withdrawal.Volume = r
	ord.Fee = f
	ord.FeeCurrency = route.Out.String()
	return
}
