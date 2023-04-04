package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"time"
)

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.NID())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	cs := []*entity.Token{}
	ps := []*entity.Pair{}
	ts, err := k.retreiveTokens()
	if err != nil {
		return nil, err
	}

	for _, p := range req.Pairs {
		for _, t := range ts.Data {
			if t.Currency == p.BC.ET.Currency && t.ChainName == p.BC.ET.ChainName {
				p.BC.ET.Chain = t.Chain
				p.BC.Decimals, _ = strconv.Atoi(t.WalletPrecision)
			} else if t.Currency == p.QC.ET.Currency && t.ChainName == p.QC.ET.ChainName {
				p.QC.ET.Chain = t.Chain
				p.QC.Decimals, _ = strconv.Atoi(t.WalletPrecision)
			}
			if p.BC.ET.Chain != "" && p.QC.ET.Chain != "" {
				break
			}
		}
		ep, err := p.ToEntity(func(t dto.Token) (entity.ExchangeToken, error) {
			bt, err := time.ParseDuration(t.BlockTime)
			if err != nil {
				return nil, err
			}

			return &Token{
				Currency:  t.Currency,
				ChainName: t.ChainName,
				Chain:     t.Chain,

				DepositAddress: t.DepositAddress,
				DepositTag:     t.DepositTag,

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
		p.Exchange = k.NID()
		res.Added = append(res.Added, *p)
		ps2 = append(ps2, p)
		bc := p.T1
		qc := p.T2
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
