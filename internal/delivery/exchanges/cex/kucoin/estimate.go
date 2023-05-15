package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) EstimateAmountOut(in, out entity.TokenId,
	amount float64, lvl uint) (*entity.EstimateAmount, error) {
	p, err := k.pairs.Get(k.Id(), in.String(), out.String())
	if err != nil {
		return nil, err
	}

	if !p.Enable {
		return nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage("pair is not enable right now"))
	}

	return k.estimateAmountOut(p, in, out, amount, lvl)
}

func (k *kucoinExchange) estimateAmountOut(p *entity.Pair, in, out entity.TokenId,
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

	var (
		depositEnable, withdrawEnable               bool
		exchangeFeeAmount, price, amountOut, spread float64
		err1, err2, err3, err4                      error
	)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		depositEnable, _, err1 = k.isDipositAndWithdrawEnable(In)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, withdrawEnable, err2 = k.isDipositAndWithdrawEnable(Out)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		price, err3 = k.price(p)
		if err3 != nil {
			return
		}
		spread, err3 = k.spread(lvl, p, price)
		if err3 != nil {
			return
		}

		if p.T1.String() == in.String() {
			amountOut = (price - (price * spread)) * amount
			es.FeeRate = p.FeeRate2
		} else {
			amountOut = (1 / (price + (price * spread))) * amount
			es.FeeRate = p.FeeRate1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		exchangeFeeAmount, err4 = k.exchangeFeeAmount(eOut, p)
	}()
	wg.Wait()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return es, errors.Wrap(errors.ErrInternal)
	}

	amountOut = amountOut - exchangeFeeAmount
	feeAmount := amountOut * es.FeeRate
	amountOut = amountOut - feeAmount - Out.MinWithdrawalFee
	if amountOut < Out.MinWithdrawalSize {
		if err := k.minAndMax(p); err != nil {
			return nil, errors.Wrap(errors.ErrInternal)
		}
		if err := k.pairs.Update(k.Id(), p); err != nil {
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

	if depositEnable && withdrawEnable {
		es.FeeAmount = feeAmount
		es.ExchangeFee = p.ExchangeFee
		es.ExchangeFeeAmount = exchangeFeeAmount
		es.FeeCurrency = out
		es.AmountOut = amountOut
		es.SpreadRate = spread
		es.Price = price
		return es, nil
	}
	return es, errors.Wrap(errors.ErrNotFound)
}

func (k *kucoinExchange) price(p *entity.Pair) (float64, error) {
	ep := p.EP.(*ExchangePair)
	if ep.HasIntermediaryCoin {
		var (
			bc, qc     string
			p0, p1     float64
			err0, err1 error
		)
		wg := &sync.WaitGroup{}
		qc = ep.IC1.Currency
		wg.Add(1)
		go func() {
			defer wg.Done()
			bc = p.T1.ET.(*Token).Currency
			p0, err1 = k.ticker(bc, qc)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			bc = p.T2.ET.(*Token).Currency
			p1, err1 = k.ticker(bc, qc)
		}()
		wg.Wait()

		if err0 != nil {
			return 0, err0
		}
		if err1 != nil {
			return 0, err1
		}
		return applyFee(applyFee(p0, ep.KucoinFeeRate1), ep.KucoinFeeRate2) / p1, nil
	}
	price, err := k.ticker(p.T1.ET.(*Token).Currency, p.T2.ET.(*Token).Currency)
	if err != nil {
		return 0, err
	}
	price = applyFee(price, ep.KucoinFeeRate1)
	return price, nil
}

func (k *kucoinExchange) ticker(bc, qc string) (float64, error) {
	res, err := k.readApi.TickerLevel1(bc + "-" + qc)
	if err != nil {
		return 0, err
	}
	tl := &kucoin.TickerLevel1Model{}
	if err := res.ReadData(tl); err != nil {
		return 0, err
	}

	return strconv.ParseFloat(tl.Price, 64)
}

func applyFee(price, fee float64) float64 {
	return price - (price * fee)
}
