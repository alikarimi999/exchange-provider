package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.Name())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	cs := []*Token{}
	ps := []*entity.Pair{}
	for _, p := range req.Pairs {
		ps = append(ps, p.ToEntity(func(t dto.Token) entity.ExchangeToken {
			return &Token{
				TokenId:             t.TokenId,
				ChainId:             t.ChainId,
				BlockTime:           t.BlockTime,
				WithdrawalPrecision: t.WithdrawalPrecision,
			}
		}))
	}

	for _, p := range ps {
		if k.pairs.Exists(k.Id(), p.T1.String(), p.T2.String()) {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		if err := k.pls.support(p, false); err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}

		if err := k.setInfos(p); err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		res.Added = append(res.Added, *p)
		bc := p.T1.ET.(*Token)
		qc := p.T2.ET.(*Token)
		cs = append(cs, bc)
		cs = append(cs, qc)

	}

	k.supportedCoins.add(cs)
	return res, nil

}

func (k *kucoinExchange) EstimateAmountOut(in, out *entity.Token,
	amount float64) (float64, float64, error) {
	p, ok := k.pairs.Get(k.Id(), in.String(), out.String())
	if !ok {
		return 0, 0, errors.Wrap(errors.ErrNotFound)
	}

	bc := p.T1.ET.(*Token)
	qc := p.T2.ET.(*Token)
	if p.T1.Equal(in) {
		min := bc.MinOrderSize
		max := bc.MaxOrderSize
		if amount < min || amount > max {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	} else {
		min := qc.MinOrderSize
		max := qc.MaxOrderSize
		if amount < min || amount > max {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	}

	res, err := k.readApi.TickerLevel1(symbol(bc, qc))
	if err != nil {
		return 0, 0, err
	}
	tl := kucoin.TickerLevel1Model{}
	if err := res.ReadData(tl); err != nil {
		return 0, 0, err
	}
	f, err := numbers.StringToBigFloat(tl.Price)
	if err != nil {
		return 0, 0, err
	}

	if p.T1.Equal(in) {
		af, _ := big.NewFloat(0).Mul(f, big.NewFloat(amount)).Float64()
		return af, 0, nil
	} else {
		price := big.NewFloat(0).Quo(big.NewFloat(1), f)
		af, _ := big.NewFloat(0).Mul(price, big.NewFloat(amount)).Float64()
		return af, 0, nil
	}

}

func (k *kucoinExchange) RemovePair(t1, t2 *entity.Token) error {
	k.pairs.Remove(k.Id(), t1.String(), t2.String())
	return nil
}

func symbol(bc, qc *Token) string {
	return bc.TokenId + "-" + qc.TokenId
}
