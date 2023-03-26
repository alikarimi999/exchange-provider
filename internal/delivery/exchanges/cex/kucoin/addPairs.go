package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"time"
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
		ep, err := p.ToEntity(func(t dto.Token) (entity.ExchangeToken, error) {
			bt, err := time.ParseDuration(t.BlockTime)
			if err != nil {
				return nil, err
			}
			return &Token{
				TokenId:             t.TokenId,
				Network:             t.Network,
				BlockTime:           bt,
				WithdrawalPrecision: t.WithdrawalPrecision,
			}, nil
		})
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		ps = append(ps, ep)
	}

	ps2 := []*entity.Pair{}
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
		p.LP = k.Id()
		p.Exchange = k.Name()

		res.Added = append(res.Added, *p)
		ps2 = append(ps2, p)
		bc := p.T1.ET.(*Token)
		qc := p.T2.ET.(*Token)
		cs = append(cs, bc)
		cs = append(cs, qc)
	}

	if len(ps2) > 0 {
		if err := k.pairs.Add(k, ps2...); err != nil {
			return nil, err
		}
		k.supportedCoins.add(cs)
	}
	return res, nil

}
