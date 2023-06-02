package binance

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/adshao/go-binance/v2"
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
	var eIn, eOut *entity.Token
	if p.T1.String() == in.String() {
		min := p.T1.Min
		max := p.T1.Max
		if (min != 0 && amount < min) || (max != 0 && amount > max) {
			return es, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
		}
		In = p.T1.ET.(*Token)
		Out = p.T2.ET.(*Token)
		eIn = p.T1
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
		eIn = p.T2
		eOut = p.T1
	}

	var (
		outEFA, price, amountOut, spread float64
		errs                             error
		p0, p1                           float64
	)

	depositEnable, _, err := ex.isDipositAndWithdrawEnable(In)
	if err != nil {
		return nil, err
	}
	if !depositEnable {
		return nil, errors.Wrap(errors.ErrInternal)
	}
	_, withdrawEnable, err := ex.isDipositAndWithdrawEnable(Out)
	if err != nil {
		return nil, err
	}
	if !withdrawEnable {
		return nil, errors.Wrap(errors.ErrInternal)
	}

	amount, _ = strconv.ParseFloat(trim(big.NewFloat(amount).Text('f', 18), In.OrderPrecision), 64)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		errs = ex.setOrderFeeRate(p)
		if errs != nil {
			return
		}

		p0, p1, err = ex.price(p)
		if err != nil {
			errs = err
			return
		}
		price = ex.calcPrice(p0, p1, in, out, p)
		spread, err = ex.spread(lvl, p, price)
		if err != nil {
			errs = err
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
		outEFA, err = ex.exchangeFeeAmount(eOut, p)
		if err != nil {
			errs = err
			return
		}
	}()
	wg.Wait()

	if errs != nil {
		return es, errors.Wrap(errors.ErrInternal)
	}

	amountOut = amountOut - outEFA
	feeAmount := amountOut * es.FeeRate
	amountOut = amountOut - feeAmount - Out.MinWithdrawalFee

	if amountOut < Out.MinWithdrawalSize+Out.MinWithdrawalFee {

		var (
			bcEFA, qcEFA float64
			s0, s1       *binance.Symbol
			err0, err1   error
		)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			inEFA, err := ex.exchangeFeeAmount(eIn, p)
			if err != nil {
				err0 = err
			}
			if eIn.String() == p.T1.String() {
				bcEFA = inEFA
				qcEFA = outEFA
			} else {
				bcEFA = outEFA
				qcEFA = inEFA
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			s0, s1, err1 = ex.getPairSymbols(p)
		}()
		wg.Wait()
		if err0 != nil {
			return nil, err0
		}
		if err1 != nil {
			return nil, err1
		}

		if err := ex.minAndMax(p, p0, p1, bcEFA, qcEFA, s0, s1); err != nil {
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

	if depositEnable && withdrawEnable {
		es.AmountIn = amount
		es.FeeAmount = feeAmount
		es.ExchangeFee = p.ExchangeFee
		es.ExchangeFeeAmount = outEFA
		es.FeeCurrency = out
		es.AmountOut = amountOut
		es.SpreadRate = spread
		es.Price = price
		return es, nil
	}
	return es, errors.Wrap(errors.ErrNotFound)
}
