package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
)

func (d *EvmDex) EstimateAmountOut(in, out *entity.Token, amount float64) (float64, float64, error) {
	t1, err := d.get(in.TokenId)
	if err != nil {
		return 0, 0, err
	}

	t2, err := d.get(out.TokenId)
	if err != nil {
		return 0, 0, err
	}

	amountOut, _, err := d.dex.EstimateAmountOut(t1, t2, amount)
	return amountOut, 0, err
}

func (d *EvmDex) price(in, out *types.Token) (*entity.Pair, error) {

	t1 := &types.Token{}
	t2 := &types.Token{}
	if in.Address.Hash().Big().Cmp(out.Address.Hash().Big()) == -1 {
		t1 = in
		t2 = out
	} else {
		t2 = in
		t1 = out
	}

	ps := []*entity.Pair{&entity.Pair{
		T1: t1.ToEntity(d.TokenStandard),
		T2: t2.ToEntity(d.TokenStandard),
	}}
	if err := d.dex.Prices(ps); err != nil {
		return nil, err
	}

	return ps[0], nil
}
