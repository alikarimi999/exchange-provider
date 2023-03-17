package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.Name())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	cs := map[string]*kuToken{}
	ps := []*pair{}
	for _, p := range req.Pairs {
		ps = append(ps, fromDto(p))
	}

	aps := []*pair{}
	for _, p := range ps {
		if _, err := k.exchangePairs.get(p.BC.toEntityCoin().Token,
			p.QC.toEntityCoin().Token); err == nil {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		ok, err := k.pls.support(p)
		if err != nil {
			return nil, err
		}

		if !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  errors.Wrap(errors.ErrBadRequest, errors.New("pair not supported")),
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

		k.v.Set(fmt.Sprintf("%s.pairs.%s", k.Name(), p.Id()), p)
		if err := k.v.WriteConfig(); err != nil {
			k.l.Error(agent, err.Error())
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		aps = append(aps, p)
		res.Added = append(res.Added, entity.Pair{
			T1: &entity.PairToken{
				Token: &entity.Token{TokenId: p.BC.TokenId, ChainId: string(p.BC.ChainId)},
			},
			T2: &entity.PairToken{
				Token: &entity.Token{TokenId: p.QC.TokenId, ChainId: string(p.QC.ChainId)},
			},
		})
		cs[p.BC.TokenId+string(p.BC.ChainId)] = p.BC
		cs[p.QC.TokenId+string(p.QC.ChainId)] = p.QC
	}

	k.exchangePairs.add(aps...)
	k.supportedCoins.add(cs)
	for _, p := range aps {
		ep := p.toEntity()
		ep.FeeRate = k.orderFeeRate(p)
	}

	return res, nil

}

func (k *kucoinExchange) Prices(ps ...*entity.Pair) ([]*entity.Pair, error) {
	if err := k.getAllPrices(ps); err != nil {
		return nil, err
	}

	return ps, nil
}

func (k *kucoinExchange) EstimateAmountOut(t1, t2 *entity.Token, amount float64) (float64, float64, error) {
	p, err := k.exchangePairs.get(t1, t2)
	if err != nil {
		return 0, 0, err
	}

	if p.BC.TokenId == t1.TokenId {
		min, _ := strconv.ParseFloat(p.BC.minOrderSize, 64)
		max, _ := strconv.ParseFloat(p.BC.maxOrderSize, 64)
		if amount < min || amount > max {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	} else {
		min, _ := strconv.ParseFloat(p.QC.minOrderSize, 64)
		max, _ := strconv.ParseFloat(p.QC.maxOrderSize, 64)
		if amount < min || amount > max {
			return 0, min, errors.Wrap(errors.ErrBadRequest)
		}
	}

	res, err := k.readApi.TickerLevel1(p.Symbol())
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

	if p.BC.TokenId == t1.TokenId {
		af, _ := big.NewFloat(0).Mul(f, big.NewFloat(amount)).Float64()
		return af, 0, nil
	} else {
		price := big.NewFloat(0).Quo(big.NewFloat(1), f)
		af, _ := big.NewFloat(0).Mul(price, big.NewFloat(amount)).Float64()
		return af, 0, nil
	}

}

func (k *kucoinExchange) RemovePair(bc, qc *entity.Token) error {
	p, err := k.exchangePairs.get(bc, qc)
	if err != nil {
		return err
	}
	delete(k.v.Get(fmt.Sprintf("%s.pairs", k.Name())).(map[string]interface{}),
		strings.ToLower(p.Id()))
	if err := k.v.WriteConfig(); err != nil {
		return err
	}

	k.exchangePairs.remove(p.Id())
	return nil
}
