package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) EstimateAmountOut(in, out *entity.Token,
	amount float64) (float64, float64, error) {
	p, ok := k.pairs.Get(k.Id(), in.String(), out.String())
	if !ok {
		return 0, 0, errors.Wrap(errors.ErrNotFound)
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	if p.T1.Equal(in) {
		min := p.T1.Min
		max := p.T1.Max
		if (min != 0 && amount < min) || (max != 0 && amount > max) {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	} else {
		min := p.T2.Min
		max := p.T2.Max
		if (min != 0 && amount < min) || (max != 0 && amount > max) {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	}

	price, err := k.price(bc, qc)
	if err != nil {
		return 0, 0, errors.Wrap(errors.ErrInternal)
	}
	if p.T1.Equal(in) {
		return price * amount, 0, nil
	} else {
		return (1 / price) * amount, 0, nil
	}

}

func (k *kucoinExchange) price(bc, qc *Token) (float64, error) {
	res, err := k.readApi.TickerLevel1(symbol(bc, qc))
	if err != nil {
		return 0, err
	}
	tl := &kucoin.TickerLevel1Model{}
	if err := res.ReadData(tl); err != nil {
		return 0, err
	}

	return strconv.ParseFloat(tl.Price, 64)
}

func (k *kucoinExchange) RemovePair(t1, t2 *entity.Token) error {
	k.pairs.Remove(k.Id(), t1.String(), t2.String())
	return nil
}

func symbol(bc, qc *Token) string {
	return bc.TokenId + "-" + qc.TokenId
}
