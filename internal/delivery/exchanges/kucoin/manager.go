package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
)

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.Id())
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

		k.v.Set(fmt.Sprintf("%s.pairs.%s", k.Id(), p.Id()), p)
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
	eps := []*entity.Pair{}
	for _, p := range aps {
		ep := p.toEntity()
		ep.FeeRate = k.orderFeeRate(p)
		eps = append(eps, ep)
	}

	pPrice, err := k.Price(eps...)
	if err != nil {
		return nil, err
	}
	k.pairs.Add(k, pPrice...)
	return res, nil

}

func (k *kucoinExchange) Price(ps ...*entity.Pair) ([]*entity.Pair, error) {
	if err := k.getAllPrices(ps); err != nil {
		return nil, err
	}

	return ps, nil
}

func (k *kucoinExchange) RemovePair(bc, qc *entity.Token) error {
	p, err := k.exchangePairs.get(bc, qc)
	if err != nil {
		return err
	}
	delete(k.v.Get(fmt.Sprintf("%s.pairs", k.Id())).(map[string]interface{}),
		strings.ToLower(p.Id()))
	if err := k.v.WriteConfig(); err != nil {
		return err
	}

	k.exchangePairs.remove(p.Id())
	k.pairs.Remove(k.Id(), bc, qc)
	return nil
}

func (k *kucoinExchange) Support(t1, t2 *entity.Token) bool {
	return k.pairs.Exists(k.Id(), t1, t2)
}
