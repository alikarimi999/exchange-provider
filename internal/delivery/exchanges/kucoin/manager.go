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
	cs := map[string]*kuCoin{}
	ps := []*pair{}
	for _, p := range req.Pairs {
		ps = append(ps, fromDto(p))
	}

	aps := []*pair{}
	for _, p := range ps {
		if _, err := k.exchangePairs.get(p.BC.toEntityCoin().Coin, p.QC.toEntityCoin().Coin); err == nil {
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
			C1: &entity.PairCoin{
				Coin: &entity.Coin{CoinId: p.BC.CoinId, ChainId: p.BC.ChainId},
			},
			C2: &entity.PairCoin{
				Coin: &entity.Coin{CoinId: p.QC.CoinId, ChainId: p.QC.ChainId},
			},
		})
		cs[p.BC.CoinId+p.BC.ChainId] = p.BC.snapshot()

		cs[p.QC.CoinId+p.QC.ChainId] = p.QC.snapshot()
	}

	k.exchangePairs.add(aps...)
	k.supportedCoins.add(cs)

	return res, nil

}

func (k *kucoinExchange) GetAllPairs() []*entity.Pair {
	agent := fmt.Sprintf("%s.GetAllPairs", k.Id())

	pairs := []*entity.Pair{}
	ps := k.exchangePairs.snapshot()

	for _, p := range ps {
		pe := p.toEntity()
		if err := k.setPrice(pe); err != nil {
			k.l.Error(agent, err.Error())
			continue
		}
		if err := k.setOrderFeeRate(pe); err != nil {
			k.l.Error(agent, err.Error())
			continue
		}

		pairs = append(pairs, pe)
	}
	return pairs
}

func (k *kucoinExchange) GetPair(bc, qc *entity.Coin) (*entity.Pair, error) {
	p, err := k.exchangePairs.get(bc, qc)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	pe := p.toEntity()
	if err := k.setPrice(pe); err != nil {
		return nil, err
	}
	if err := k.setOrderFeeRate(pe); err != nil {
		return nil, err
	}

	return pe, nil
}

func (k *kucoinExchange) RemovePair(bc, qc *entity.Coin) error {
	if p, err := k.exchangePairs.get(bc, qc); err != nil {
		delete(k.v.Get(fmt.Sprintf("%s.pairs", k.Id())).(map[string]interface{}),
			strings.ToLower(p.Id()))
		if err := k.v.WriteConfig(); err != nil {
			return err
		}

		k.exchangePairs.remove(p.Id())
		return nil
	}
	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func (k *kucoinExchange) Support(c1, c2 *entity.Coin) bool {
	if _, err := k.exchangePairs.get(c1, c2); err != nil {
		return false
	}
	return true
}
