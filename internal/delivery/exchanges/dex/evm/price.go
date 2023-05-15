package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
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

	es := &entity.EstimateAmount{P: p, FeeCurrency: in}
	var (
		In, Out *types.Token
	)

	if p.T1.String() == in.String() {
		In = types.TokenFromEntity(p.T1)
		Out = types.TokenFromEntity(p.T2)
		es.FeeRate = p.FeeRate1
	} else {
		In = types.TokenFromEntity(p.T2)
		Out = types.TokenFromEntity(p.T1)
		es.FeeRate = p.FeeRate2
	}

	exchangeFeeAmount, err := d.exchangeFeeAmount(in, p)
	if err != nil {
		return nil, err
	}
	es.ExchangeFee = p.ExchangeFee
	es.ExchangeFeeAmount = exchangeFeeAmount

	amount = amount - exchangeFeeAmount
	es.FeeAmount = amount * es.FeeRate
	amount = amount - es.FeeAmount
	amountOut, _, err := d.dex.EstimateAmountOut(In, Out, amount)
	es.AmountOut = amountOut
	return es, err
}

func (d *evmDex) exchangeFeeAmount(in entity.TokenId, p *entity.Pair) (float64, error) {

	var (
		stAmount float64
		St, In   *types.Token
	)

	if p.T1.String() == in.String() {
		In = types.TokenFromEntity(p.T1)
		St = p.T1.ET.(*types.Token)
		stAmount = p.T1.Min
	} else {
		In = types.TokenFromEntity(p.T2)
		St = p.T2.ET.(*types.Token)
		stAmount = p.T2.Min
	}

	stOut, _, err := d.dex.EstimateAmountOut(In, St, stAmount)
	if err != nil {
		return 0, err
	}

	if stOut == 0 {
		return 0, fmt.Errorf("unable to calculate exchangeFeeAmount")
	}
	return (stOut / stAmount) * p.ExchangeFee, nil
}
