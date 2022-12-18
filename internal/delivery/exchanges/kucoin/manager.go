package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"time"
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
		s1, ok := k.cfg.Chains[chainId(p.T1.ChainId)]
		if !ok {
			if p.T1.Standard == "" || p.T1.BlockTime == 0 {
				res.Failed = append(res.Failed, &entity.PairsErr{
					Pair: p.String(),
					Err:  fmt.Errorf("set 'standard' and 'block_time' for %s", p.T1.String()),
				})
				continue
			}
			k.cfg.Chains[chainId(p.T1.ChainId)] = struct {
				standard
				time.Duration
			}{standard(p.T1.Standard), p.T1.BlockTime}
		} else if s1.standard != standard(p.T1.Standard) {
			p.T1.Standard = string(s1.standard)
		} else if s1.Duration != p.T1.BlockTime {
			p.T1.BlockTime = s1.Duration
		}

		s2, ok := k.cfg.Chains[chainId(p.T2.ChainId)]
		if !ok {
			if p.T2.Standard == "" || p.T2.BlockTime == 0 {
				res.Failed = append(res.Failed, &entity.PairsErr{
					Pair: p.String(),
					Err:  fmt.Errorf("set 'standard' and 'block_time' for %s", p.T2.String()),
				})
				continue
			}
			k.cfg.Chains[chainId(p.T2.ChainId)] = struct {
				standard
				time.Duration
			}{standard(p.T2.Standard), p.T2.BlockTime}
		} else if s2.standard != standard(p.T2.Standard) {
			p.T2.Standard = string(s2.standard)
		} else if s2.Duration != p.T2.BlockTime {
			p.T2.BlockTime = s2.Duration
		}

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
			T1: &entity.PairCoin{
				Token: &entity.Token{TokenId: p.BC.TokenId, ChainId: string(p.BC.ChainId)},
			},
			T2: &entity.PairCoin{
				Token: &entity.Token{TokenId: p.QC.TokenId, ChainId: string(p.QC.ChainId)},
			},
		})
		cs[p.BC.TokenId+string(p.BC.ChainId)] = p.BC

		cs[p.QC.TokenId+string(p.QC.ChainId)] = p.QC
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

		price1, price2, err := k.getPrice(p)
		if err != nil {
			k.l.Error(agent, err.Error())
			continue
		}
		pe.Price1 = price1
		pe.Price2 = price2

		f := k.orderFeeRate(p)
		if f != "" {
			pe.FeeRate = f
			pairs = append(pairs, pe)
		}
	}
	return pairs
}

func (k *kucoinExchange) GetPair(bc, qc *entity.Token) (*entity.Pair, error) {
	agent := fmt.Sprintf("%s.GetPair", k.Id())
	p, err := k.exchangePairs.get(bc, qc)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
	}
	pe := p.toEntity()
	price1, price2, err := k.getPrice(p)
	if err != nil {
		k.l.Error(agent, err.Error())
		return nil, err
	}

	pe.Price1 = price1
	pe.Price2 = price2

	f := k.orderFeeRate(p)
	if f != "" {
		pe.FeeRate = f
		return pe, nil
	}

	return nil, errors.New("")
}

func (k *kucoinExchange) RemovePair(bc, qc *entity.Token) error {
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

func (k *kucoinExchange) Support(c1, c2 *entity.Token) bool {
	if _, err := k.exchangePairs.get(c1, c2); err != nil {
		return false
	}
	return true
}
