package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"time"
)

func (k *exchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.NID())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	ps := []*entity.Pair{}
	for _, p := range req.Pairs {
		p.BC.ToUpper()
		p.QC.ToUpper()
		if k.pairs.Exists(k.Id(), p.BC.String(), p.QC.String()) {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		fn := func(t dto.Token) (entity.ExchangeToken, error) {
			bt, err := time.ParseDuration(t.BlockTime)
			if err != nil {
				return nil, err
			}
			if t.StableToken == "" {
				return nil, fmt.Errorf("stableToken cannot be empty")
			}

			return &Token{
				Currency:            t.Currency,
				ChainName:           t.ChainName,
				StableToken:         t.StableToken,
				WithdrawalPrecision: t.WithdrawalPrecision,
				BlockTime:           bt,
			}, nil
		}

		ep, err := p.ToEntity(fn)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		ep.EP = &ExchangePair{}
		if p.IC != "" {
			ic := Token{Currency: strings.ToUpper(p.IC)}
			ep.EP = &ExchangePair{
				HasIntermediaryCoin: true,
				IC1:                 ic.snapshot(),
				IC2:                 ic.snapshot(),
			}
		}

		if err := k.setInfos(ep); err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err:  err,
			})
			continue
		}
		ep.LP = k.Id()
		ep.Exchange = k.NID()
		res.Added = append(res.Added, p.String())
		ps = append(ps, ep)
	}

	if len(ps) > 0 {
		if err := k.pairs.Add(k, ps...); err != nil {
			return nil, err
		}
	}
	return res, nil

}

func (k *exchange) RemovePair(t1, t2 entity.TokenId) error {
	return k.pairs.Remove(k.Id(), t1.String(), t2.String(), true)
}
