package evm

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (d *evmDex) EstimateAmountOut(in, out entity.TokenId,
	amount float64, lvl uint) (*entity.EstimateAmount, error) {
	p, err := d.pairs.Get(d.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}

	if !p.Enable {
		return nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage("pair is not enable right now"))
	}

	es := &entity.EstimateAmount{P: p}

	var In, Out *entity.Token
	if p.T1.String() == in.String() {
		In = p.T1
		Out = p.T2
	} else {
		In = p.T2
		Out = p.T1
	}

	amountOut, _, err := d.dex.EstimateAmountOut(In, Out, amount)
	es.AmountOut = amountOut
	return es, err
}
