package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
)

func (d *evmDex) AddPairs(data interface{}) (*entity.AddPairsResult, error) {

	req := data.(*dto.AddPairsRequest)
	ps := []*entity.Pair{}
	res := &entity.AddPairsResult{}
	for _, p := range req.Pairs {
		ep, err := p.ToEntity(d.TokenStandard, d.Network)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}
		ps = append(ps, ep)
	}

	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	add := []*entity.Pair{}

	for _, p := range ps {
		if p.T1.Id.Network != d.Network || p.T2.Id.Network != d.Network {
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
			p.Exchange = d.NID()
			add = append(add, p)
			res.Added = append(res.Added, p.String())
			mux.Unlock()
		}(p)
	}

	wg.Wait()

	if len(add) > 0 {
		if err := d.pairs.Add(d, add...); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (d *evmDex) checkPair(t1, t2 *entity.Token) error {
	T1 := types.TokenFromEntity(t1)
	T2 := types.TokenFromEntity(t2)

	amOut, _, err := d.dex.EstimateAmountOut(T1, T2, 1)
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

func (d *evmDex) RemovePair(t1, t2 entity.TokenId) error {
	return d.pairs.Remove(d.Id(), t1.String(), t2.String(), true)
}
