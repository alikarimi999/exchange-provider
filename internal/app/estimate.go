package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"sync"
)

type estimate struct {
	exs []entity.Exchange
	ess []*entity.EstimateAmount
}

func (o *OrderUseCase) EstimateAmountOut(in, out entity.TokenId,
	amount float64, lp, lvl uint) ([]*entity.EstimateAmount, []entity.Exchange, error) {
	ess, exs, _, err := o.estimateAmountOut(in, out, amount, lp, lvl, nil)
	return ess, exs, err
}

func (o *OrderUseCase) estimateAmountOut(in, out entity.TokenId, amount float64,
	lp, lvl uint, excludeLp []uint) ([]*entity.EstimateAmount, []entity.Exchange, []uint, error) {

	if lp > 0 {
		ex, err := o.exs.Get(lp)
		if err != nil {
			return nil, nil, nil, err
		}
		if !ex.IsEnable() {
			return nil, nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("lp is disable"))
		}
		es, err := ex.EstimateAmountOut(in, out, amount, lvl, nil)
		if err != nil {
			return nil, nil, nil, err
		}
		return es, []entity.Exchange{ex, ex}, nil, nil
	}

	exs := o.exs.GetAll()
	wg := &sync.WaitGroup{}
	pMux := &sync.Mutex{}
	minAndMax := []*entity.Pair{}
	estimates := []*estimate{}

start:
	for _, ex := range exs {
		if !ex.IsEnable() {
			continue
		}
		for _, el := range excludeLp {
			if ex.Id() == el {
				continue start
			}
		}

		wg.Add(1)
		go func(ex entity.Exchange) {
			defer wg.Done()
			ess, err := ex.EstimateAmountOut(in, out, amount, lvl, nil)
			pMux.Lock()
			if err == nil && ess[0].AmountOut > 0 {
				es := &estimate{}
				es.ess = append(es.ess, ess[0])
				es.exs = append(es.exs, ex)
				if len(ess) == 2 && ess[1].AmountOut > 0 {
					es.ess = append(es.ess, ess[1])
				}
				estimates = append(estimates, es)
			} else if err != nil && strings.Contains(err.Error(), "min") && ess[0].P != nil {
				minAndMax = append(minAndMax, ess[0].P)
			}
			pMux.Unlock()
		}(ex)
	}
	wg.Wait()

	var max, min float64
	if len(estimates) == 0 {
		if len(minAndMax) > 0 {
			if minAndMax[0].T1.String() == in.String() {
				min = minAndMax[0].T1.Min
				max = minAndMax[0].T1.Max
				for _, p := range minAndMax {
					if p.T1.Min < min {
						min = p.T1.Min
					}
					if p.T1.Max == 0 || p.T1.Max > max {
						max = p.T1.Max
					}
				}
			} else {
				min = minAndMax[0].T2.Min
				max = minAndMax[0].T2.Max
				for _, p := range minAndMax {
					if p.T2.Min < min {
						min = p.T2.Min
					}
					if p.T2.Max == 0 || p.T2.Max > max {
						max = p.T2.Max
					}
				}
			}

			return nil, nil, nil, errors.Wrap(errors.ErrNotFound,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
		}
		return nil, nil, nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("pair '%s/%s' not found", in.String(), out.String())))
	}
	aLPs := []uint{}
	es := &estimate{}
	for _, est := range estimates {
		if len(est.ess) > 0 {
			aLPs = append(aLPs, est.exs[0].Id())
			if len(es.ess) == 0 || est.ess[0].AmountOut > es.ess[0].AmountOut {
				if len(es.ess) == 0 {
					es.ess = append(es.ess, est.ess[0])
					es.exs = append(es.exs, est.exs[0])
				} else if len(es.ess) == 1 {
					es.ess[0] = est.ess[0]
					es.exs[0] = est.exs[0]
				}
			}
		}
	}

	for _, est := range estimates {
		if len(est.ess) == 2 && (len(es.ess) == 1 || est.ess[1].AmountIn < es.ess[1].AmountIn) {
			if len(es.ess) == 1 {
				es.ess = append(es.ess, est.ess[1])
				es.exs = append(es.exs, est.exs[0])
			} else if len(es.ess) == 2 {
				es.ess[1] = est.ess[1]
				es.exs[1] = est.exs[0]
			}
		}
	}
	return es.ess, es.exs, aLPs, nil
}
