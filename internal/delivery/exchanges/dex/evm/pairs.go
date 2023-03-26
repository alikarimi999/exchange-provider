package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
)

func (d *EvmDex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {

	req := data.(*dto.AddPairsRequest)
	ps := []*entity.Pair{}
	for _, p := range req.Pairs {
		ps = append(ps, p.ToEntity(func(t dto.Token) entity.ExchangeToken { return &Token{} }))
	}

	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	res := &entity.AddPairsResult{}
	add := []*entity.Pair{}

	for _, p := range ps {
		if p.T1.Network != d.Network || p.T2.Network != d.Network {
			mux.Lock()
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("invalid token network")})
			mux.Unlock()
			continue
		}
		if d.pairs.Exists(d.Id(), p.T1.String(), p.T2.String()) {
			mux.Lock()
			res.Existed = append(res.Existed, p.String())
			mux.Unlock()
			continue
		}
		wg.Add(1)
		go func(p *entity.Pair) {
			defer wg.Done()
			if err := d.checkPair(p.T1, p.T2); err != nil {
				mux.Lock()
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
				mux.Unlock()
				return
			}
			mux.Lock()
			p.LP = d.Id()
			p.Exchange = d.Name()
			add = append(add, p)
			res.Added = append(res.Added, *p)
			mux.Unlock()
		}(p)
	}

	wg.Wait()

	if len(add) > 0 {
		if err := d.pairs.Add(d, add...); err != nil {
			return nil, err
		}

		for _, p := range add {
			if !d.ts.exists(p.T1.String()) {
				d.ts.add(p.T1)
			}
			if !d.ts.exists(p.T2.String()) {
				d.ts.add(p.T2)
			}
			p.LP = d.Id()
		}
	}
	return res, nil
}

func (d *EvmDex) checkPair(t1, t2 *entity.Token) error {
	amOut, _, err := d.dex.EstimateAmountOut(t1, t2, 1)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if amOut == 0 {
		return errors.Wrap(errors.ErrNotFound)
	}
	return nil
}

func (d *EvmDex) RemovePair(t1, t2 *entity.Token) error { return nil }
