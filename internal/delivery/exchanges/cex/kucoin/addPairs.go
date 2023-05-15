package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const max_conccurrent_jobs = 20

func (k *kucoinExchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	agent := fmt.Sprintf("%s.AddPairs", k.NID())
	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(agent, errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	ps := []*entity.Pair{}
	ts, err := k.retreiveTokens()
	if err != nil {
		return nil, err
	}
	for _, p := range req.Pairs {
		p.BC.ToUpper()
		p.QC.ToUpper()
		if k.pairs.Exists(k.Id(), p.BC.String(), p.QC.String()) {
			res.Existed = append(res.Existed, p.String())
			continue
		}
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

		if p.BC.ET.Chain == "" {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err: fmt.Errorf("token with currency '%s' and chainName '%s' does not exists in kucoin",
					p.BC.ET.Currency, p.BC.ET.ChainName),
			})
			continue
		}
		if p.QC.ET.Chain == "" {
			res.Failed = append(res.Failed, &entity.PairsErr{
				Pair: p.String(),
				Err: fmt.Errorf("token with currency '%s' and chainName '%s' does not exists in kucoin",
					p.QC.ET.Currency, p.QC.ET.ChainName),
			})
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
				Chain:               t.Chain,
				StableToken:         t.StableToken,
				BlockTime:           bt,
				WithdrawalPrecision: t.WithdrawalPrecision,
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
		ps = append(ps, ep)
	}

	ps2 := []*entity.Pair{}
	if len(ps) > 0 {
		if err := k.pls.downloadList(); err != nil {
			return nil, err
		}
		wg := &sync.WaitGroup{}
		mux := sync.Mutex{}
		waitChan := make(chan struct{}, max_conccurrent_jobs)

		for _, p := range ps {
			waitChan <- struct{}{}
			wg.Add(1)
			go func(p *entity.Pair) {
				defer func() {
					<-waitChan
					wg.Done()
				}()
				if err := k.support(p); err != nil {
					mux.Lock()
					res.Failed = append(res.Failed, &entity.PairsErr{
						Pair: p.String(),
						Err:  err,
					})
					mux.Unlock()
					return
				}

				if err := k.setInfos(p); err != nil {
					mux.Lock()
					res.Failed = append(res.Failed, &entity.PairsErr{
						Pair: p.String(),
						Err:  err,
					})
					mux.Unlock()
					return
				}
				p.LP = k.Id()
				p.Exchange = k.NID()
				mux.Lock()
				res.Added = append(res.Added, p.String())
				ps2 = append(ps2, p)
				mux.Unlock()
			}(p)
		}
		wg.Wait()
	}

	if len(ps2) > 0 {
		if err := k.pairs.Add(k, ps2...); err != nil {
			return nil, err
		}
	}
	return res, nil

}

func (k *kucoinExchange) RemovePair(t1, t2 entity.TokenId) error {
	return k.pairs.Remove(k.Id(), t1.String(), t2.String(), true)
}
