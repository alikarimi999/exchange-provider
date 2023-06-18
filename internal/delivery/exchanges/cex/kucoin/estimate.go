package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/big"
	"strconv"
)

func (ex *exchange) EstimateAmountOut(in, out entity.TokenId,
	amount float64, lvl uint) (*entity.EstimateAmount, error) {
	p, err := ex.pairs.Get(ex.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}

	if !p.Enable {
		return nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage("pair is not enable right now"))
	}

	return ex.estimateAmountOut(p, in, out, amount, lvl)
}

func (ex *exchange) estimateAmountOut(p *entity.Pair, in, out entity.TokenId,
	amount float64, lvl uint) (*entity.EstimateAmount, error) {
	es := &entity.EstimateAmount{
		P: p,
	}

	var In, Out *Token
	var eOut *entity.Token
	if p.T1.String() == in.String() {
		min := p.T1.Min
		max := p.T1.Max
		if (min != 0 && amount < min) || (max != 0 && amount > max) {
			return es, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
		}
		In = p.T1.ET.(*Token)
		Out = p.T2.ET.(*Token)
		eOut = p.T2
	} else {
		min := p.T2.Min
		max := p.T2.Max
		if (min != 0 && amount < min) || (max != 0 && amount > max) {
			return es, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))

		}
		In = p.T2.ET.(*Token)
		Out = p.T1.ET.(*Token)
		eOut = p.T1
	}
	amount, _ = strconv.ParseFloat(trim(big.NewFloat(amount).Text('f', 12), In.OrderPrecision), 64)
	depositEnable, _, err := ex.isDipositAndWithdrawEnable(In)
	if err != nil {
		return nil, err
	}

	if !depositEnable {
		return nil, fmt.Errorf("pair is not enable")
	}

	_, withdrawEnable, err := ex.isDipositAndWithdrawEnable(Out)
	if err != nil {
		return nil, err
	}
	if !withdrawEnable {
		return nil, fmt.Errorf("pair is not enable")
	}

	if !ex.isPairEnabled(p) {
		return nil, fmt.Errorf("pair is not enable")
	}

	if err := ex.setOrderFeeRate(p); err != nil {
		return nil, err
	}

	price, err := ex.price(p)
	if err != nil {
		return nil, err
	}
	spread, err := ex.spread(lvl, p, price)
	if err != nil {
		return nil, err
	}

	var amountOut float64
	if p.T1.String() == in.String() {
		amountOut = (price - (price * spread)) * amount
		es.FeeRate = p.FeeRate2
	} else {
		amountOut = (1 / (price + (price * spread))) * amount
		es.FeeRate = p.FeeRate1
	}
	exchangeFeeAmount, err := ex.exchangeFeeAmount(eOut, p)
	if err != nil {
		return nil, err
	}

	amountOut = amountOut - exchangeFeeAmount
	feeAmount := amountOut * es.FeeRate
	amountOut = amountOut - feeAmount - Out.MinWithdrawalFee
	if amountOut < Out.MinWithdrawalSize+Out.MinWithdrawalFee {
		if err := ex.minAndMax(p); err != nil {
			return nil, errors.Wrap(errors.ErrInternal)
		}
		if err := ex.pairs.Update(ex.Id(), p); err != nil {
			return nil, errors.Wrap(errors.ErrInternal)
		}

		if p.T1.String() == in.String() {
			return es, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min amount updated to %f", p.T1.Min)))
		} else {
			return es, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min amount updated to %f", p.T2.Min)))
		}
	}

	es.AmountIn = amount
	es.FeeAmount = feeAmount
	es.ExchangeFee = p.ExchangeFee
	es.ExchangeFeeAmount = exchangeFeeAmount
	es.FeeCurrency = out
	es.AmountOut = amountOut
	es.SpreadRate = spread
	es.Price = price
	return es, nil
}

func (ex *exchange) price(p *entity.Pair) (float64, error) {
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {

		bc := p.T1.ET.(*Token).Currency
		qc := ep.IC1.Currency
		p0, err := ex.ticker(bc, qc)
		if err != nil {
			return 0, err
		}

		bc = p.T2.ET.(*Token).Currency
		qc = ep.IC2.Currency
		p1, err := ex.ticker(bc, qc)
		if err != nil {
			return 0, err
		}
		return applyFee(applyFee(p0, ep.KucoinFeeRate1), ep.KucoinFeeRate2) / p1, nil
	}
	price, err := ex.ticker(p.T1.ET.(*Token).Currency, p.T2.ET.(*Token).Currency)
	if err != nil {
		return 0, err
	}
	price = applyFee(price, ep.KucoinFeeRate1)
	return price, nil
}

func (ex *exchange) ticker(bc, qc string) (float64, error) {
	return ex.si.getPrice(bc, qc)
}
func applyFee(price, fee float64) float64 {
	return price - (price * fee)
}
