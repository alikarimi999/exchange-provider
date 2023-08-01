package evm

import "exchange-provider/internal/entity"

func (ex *exchange) minAndMax(p *entity.Pair) error {
	t1Efa, _, err := ex.ExchangeFeeAmount(p.T1.Id, p, p.ExchangeFee)
	if err != nil {
		return err
	}
	t2Efa, _, err := ex.ExchangeFeeAmount(p.T2.Id, p, p.ExchangeFee)
	if err != nil {
		return err
	}
	minT1 := t1Efa + (t1Efa * p.FeeRate1 * 2)
	minT1 = minT1 + (minT1 * 0.1)

	minT2 := t2Efa + (t2Efa * p.FeeRate2 * 2)
	minT2 = minT2 + (minT2 * 0.1)

	p.T1.Min = minT1
	p.T2.Min = minT2
	return nil
}
