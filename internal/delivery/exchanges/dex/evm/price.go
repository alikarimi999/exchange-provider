package evm

import (
	"exchange-provider/internal/entity"
)

func (d *EvmDex) Price(ps ...*entity.Pair) ([]*entity.Pair, error) {
	return d.price(ps...)
}

func (d *EvmDex) price(ps ...*entity.Pair) ([]*entity.Pair, error) {
	input := []*entity.Pair{}
	for _, p := range ps {
		if p.T1.TokenId == d.NativeToken || p.T2.TokenId == d.NativeToken {
			continue
		}
		input = append(input, p)
	}
	if err := d.Prices(input); err != nil {
		return nil, err
	}

	output := []*entity.Pair{}
	for _, p := range input {
		if p.Price1 == "" {
			continue
		}
		output = append(output, p)
		if p.T1.TokenId == d.WrappedNativeToken {
			p1 := p.Snapshot()
			p1.T1.TokenId = d.NativeToken
			p1.T1.Native = true
			output = append(output, p1)
		} else if p.T2.TokenId == d.WrappedNativeToken {
			p1 := p.Snapshot()
			p1.T2.TokenId = d.NativeToken
			p1.T2.Native = true
			output = append(output, p1)
		}
	}
	return output, nil
}
