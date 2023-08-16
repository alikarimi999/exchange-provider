package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

type EstimateOpts struct {
	CustomizeFee   bool
	ExchangeFee    float64
	FeeRate        float64
	RevertEstimate bool
}

func (d *exchange) EstimateAmountOut(in, out entity.TokenId,
	amountIn float64, lvl uint, opts interface{}) ([]*entity.EstimateAmount, error) {
	p, err := d.pairs.Get(d.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}
	if !p.Enable {
		return nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage("pair is not enable right now"))
	}
	opt := &EstimateOpts{
		RevertEstimate: true,
	}
	if opts != nil {
		opt = opts.(*EstimateOpts)
	}
	es0 := &entity.EstimateAmount{P: p, FeeCurrency: in, AmountIn: amountIn}
	ess := []*entity.EstimateAmount{es0}
	var (
		In, Out *types.Token
	)

	if p.T1.String() == in.String() {
		if !opt.CustomizeFee {
			min := p.T1.Min
			max := p.T1.Max
			if (min != 0 && amountIn < min) || (max != 0 && amountIn > max) {
				return ess, errors.Wrap(errors.ErrBadRequest,
					errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
			}
		}
		In = types.TokenFromEntity(p.T1)
		Out = types.TokenFromEntity(p.T2)
		es0.FeeRate = p.FeeRate1
	} else {
		if !opt.CustomizeFee {
			min := p.T2.Min
			max := p.T2.Max
			if (min != 0 && amountIn < min) || (max != 0 && amountIn > max) {
				return ess, errors.Wrap(errors.ErrBadRequest,
					errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))

			}
		}
		In = types.TokenFromEntity(p.T2)
		Out = types.TokenFromEntity(p.T1)
		es0.FeeRate = p.FeeRate2
	}

	var ef, fr float64
	if opt.CustomizeFee {
		ef = opt.ExchangeFee
		fr = opt.FeeRate
	} else {
		ef = p.ExchangeFee
		fr = es0.FeeRate
	}
	inEFA, priceIn, err := d.ExchangeFeeAmount(in, p, ef)
	if err != nil {
		return nil, err
	}

	outEFA, priceOut, err := d.ExchangeFeeAmount(out, p, ef)
	if err != nil {
		return nil, err
	}
	es0.InUsd = priceIn
	es0.OutUsd = priceOut
	es0.ExchangeFee = ef
	es0.ExchangeFeeAmount = inEFA
	amount := amountIn - inEFA
	es0.FeeRate = fr
	es0.FeeAmount = amount * fr
	amount = amount - es0.FeeAmount

	if amount <= 0 {
		if !opt.CustomizeFee {
			if err := d.minAndMax(p); err != nil {
				return nil, errors.Wrap(errors.ErrInternal)
			}
			if err := d.pairs.Update(d.Id(), p, false); err != nil {
				return nil, errors.Wrap(errors.ErrInternal)
			}
		}

		if p.T1.String() == in.String() {
			return ess, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min amount updated to %f", p.T1.Min)))
		} else {
			return ess, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min amount updated to %f", p.T2.Min)))
		}
	}

	amountOut, _, err := d.dex.EstimateAmountOut(In, Out, amount)
	if err != nil {
		return nil, err
	}

	es0.AmountOut = amountOut

	if opt.RevertEstimate {

		amIn, _, err := d.dex.EstimateAmountOut(Out, In, amountOut)
		if err != nil {
			return nil, err
		}
		amIn = (amountIn * amountOut) / amIn
		amInPlusFee := (amIn / (1 - fr))
		feeAmount := amInPlusFee - amIn
		amIn = amInPlusFee + outEFA

		es1 := &entity.EstimateAmount{
			AmountIn:          amIn,
			AmountOut:         amountIn,
			FeeRate:           fr,
			FeeAmount:         feeAmount,
			ExchangeFee:       ef,
			ExchangeFeeAmount: outEFA,
			FeeCurrency:       out,
		}
		ess = append(ess, es1)
	}
	return ess, nil
}

func (d *exchange) ExchangeFeeAmount(in entity.TokenId, p *entity.Pair,
	exchangeFee float64) (float64, float64, error) {

	var (
		stAmount, stOut, price float64
		St, In                 *types.Token
		err                    error
	)

	if p.T1.String() == in.String() {
		In = types.TokenFromEntity(p.T1)
		St = &p.T1.ET.(*types.EToken).StableToken
		stAmount = p.T1.Min
	} else {
		In = types.TokenFromEntity(p.T2)
		St = &p.T2.ET.(*types.EToken).StableToken
		stAmount = p.T2.Min
	}

	if In.ContractAddress != St.ContractAddress {
		stOut, _, err = d.dex.EstimateAmountOut(In, St, stAmount)
		if err != nil {
			return 0, 0, err
		}
		if stOut == 0 {
			return 0, 0, fmt.Errorf("unable to calculate exchangeFeeAmount")
		}
		price = stOut / stAmount
	} else {
		price = 1
	}

	return exchangeFee / price, price, nil
}
