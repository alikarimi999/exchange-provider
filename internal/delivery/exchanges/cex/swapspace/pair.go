package swapspace

import (
	"exchange-provider/internal/delivery/exchanges/cex/swapspace/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"sync"
)

func (ex *exchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	req := data.(*dto.AddPairsRequest)

	ps := []*entity.Pair{}
	for _, p := range req.Pairs {
		ps = append(ps, p.ToEntity(func(t dto.Token) entity.ExchangeToken {
			return &Token{
				Code:       t.Code,
				Network:    t.Network,
				HasExtraId: t.HasExtraId,
			}
		}))
	}

	res := &entity.AddPairsResult{}

	mux := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	add := []*entity.Pair{}

	for _, p := range ps {
		wg.Add(1)
		go func(p *entity.Pair) {
			defer wg.Done()
			if ex.pairs.Exists(ex.Id(), p.T1.String(), p.T2.String()) {
				mux.Lock()
				res.Existed = append(res.Existed, p.String())
				mux.Unlock()
				return
			}

			if err := ex.checkPair(p); err != nil {
				mux.Lock()
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
				mux.Unlock()
				return
			}

			p.LP = ex.Id()
			mux.Lock()
			add = append(add, p)
			res.Added = append(res.Added, *p)
			mux.Unlock()
		}(p)
	}
	wg.Wait()
	if err := ex.pairs.Add(ex, add...); err != nil {
		return nil, err
	}
	return res, nil
}

func (ex *exchange) RemovePair(t1, t2 *entity.Token) error {
	if !ex.pairs.Exists(ex.Id(), t1.String(), t2.String()) {
		return errors.Wrap(errors.ErrNotFound)
	}
	ex.pairs.Remove(ex.Id(), t1.String(), t2.String())
	return nil
}
