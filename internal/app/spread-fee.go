package app

import "exchange-provider/internal/entity"

func (o *orderHandler) applySpreadAndFee(ord *entity.Order, route *entity.Route) error {
	aVol, sVol, rate, err := o.pc.ApplySpread(route.In, route.Out, ord.Swaps[len(ord.Swaps)-1].OutAmount)
	if err != nil {
		ord.Status = entity.OSFailed
		ord.FailedCode = entity.FCInternalError
		ord.FailedDesc = err.Error()
		return err
	}

	ord.SpreadCurrency = route.Out.String()
	ord.SpreadVol = sVol
	ord.SpreadRate = rate

	r, f, err := o.fee.ApplyFee(ord.UserId, aVol)
	if err != nil {
		ord.Status = entity.OSFailed
		ord.FailedCode = entity.FCInternalError
		ord.FailedDesc = err.Error()
		return err
	}

	ord.Withdrawal.Token = route.Out
	ord.Withdrawal.Volume = r
	ord.Fee = f
	ord.FeeCurrency = route.Out.String()
	return nil
}
