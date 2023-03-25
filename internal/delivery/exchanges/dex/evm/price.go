package evm

import (
	"exchange-provider/internal/entity"
)

func (d *EvmDex) EstimateAmountOut(in, out *entity.Token, amount float64) (float64, float64, error) {
	t1, err := d.ts.get(in.String())
	if err != nil {
		return 0, 0, err
	}

	t2, err := d.ts.get(out.String())
	if err != nil {
		return 0, 0, err
	}

	amountOut, _, err := d.dex.EstimateAmountOut(t1, t2, amount)
	return amountOut, 0, err
}
