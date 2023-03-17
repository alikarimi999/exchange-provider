package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
)

type estimate struct {
	ex        entity.Exchange
	amountOut float64
}

func (o *OrderUseCase) EstimateAmountOut(in, out *entity.Token,
	amount float64, lp uint) (entity.Exchange, float64, error) {
	return o.estimateAmountOut(in, out, amount, lp)
}

func (o *OrderUseCase) estimateAmountOut(in, out *entity.Token,
	amount float64, lp uint) (entity.Exchange, float64, error) {

	if lp > 0 {
		ex, err := o.exs.get(lp)
		if err != nil {
			return nil, 0, err
		}
		amOut, _, err := ex.EstimateAmountOut(in, out, amount)
		if err != nil {
			return nil, 0, err
		}
		return ex, amOut, nil
	}

	exs := o.exs.getAll()
	wg := &sync.WaitGroup{}
	pMux := &sync.Mutex{}
	mins := []float64{}
	estimates := []*estimate{}
	for _, ex := range exs {
		wg.Add(1)
		go func(ex entity.Exchange) {
			defer wg.Done()
			pr, min, err := ex.EstimateAmountOut(in, out, amount)
			pMux.Lock()
			if err == nil && pr > 0 {
				estimates = append(estimates, &estimate{
					ex:        ex,
					amountOut: pr,
				})
			} else if min > 0 {
				mins = append(mins, min)
			}
			pMux.Unlock()
		}(ex)
	}
	wg.Wait()

	var min float64

	if len(estimates) == 0 {
		if len(mins) > 0 {
			min = mins[0]
			for _, m := range mins {
				if m < min {
					min = m
				}
			}
			return nil, 0, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("min amount is %f", min)))
		}
		return nil, 0, errors.Wrap(errors.ErrNotFound)
	}

	es := &estimate{}
	for _, est := range estimates {
		if est.amountOut > es.amountOut {
			es = est
		}
	}

	return es.ex, es.amountOut, nil
}
