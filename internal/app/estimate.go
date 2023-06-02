package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strings"
	"sync"
)

type estimate struct {
	ex entity.Exchange
	*entity.EstimateAmount
}

func (o *OrderUseCase) EstimateAmountOut(in, out entity.TokenId,
	amount float64, lp, lvl uint) (*entity.EstimateAmount, entity.Exchange, error) {
	return o.estimateAmountOut(in, out, amount, lp, lvl)
}

func (o *OrderUseCase) estimateAmountOut(in, out entity.TokenId, amount float64,
	lp, lvl uint) (*entity.EstimateAmount, entity.Exchange, error) {

	if lp > 0 {
		ex, err := o.exs.get(lp)
		if err != nil {
			return nil, nil, err
		}
		if !ex.IsEnable() {
			return nil, nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("lp is disable"))
		}
		es, err := ex.EstimateAmountOut(in, out, amount, lvl)
		if err != nil {
			return nil, nil, err
		}
		return es, ex, nil
	}

	exs := o.exs.getAll()
	wg := &sync.WaitGroup{}
	pMux := &sync.Mutex{}
	minAndMax := []*entity.Pair{}
	estimates := []*estimate{}
	for _, ex := range exs {
		if !ex.IsEnable() {
			continue
		}
		wg.Add(1)
		go func(ex entity.Exchange) {
			defer wg.Done()
			es, err := ex.EstimateAmountOut(in, out, amount, lvl)
			pMux.Lock()
			if err == nil && es.AmountOut > 0 {
				estimates = append(estimates, &estimate{
					ex:             ex,
					EstimateAmount: es,
				})
			} else if err != nil && es != nil && strings.Contains(err.Error(), "min") && es.P != nil {
				minAndMax = append(minAndMax, es.P)
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

			return nil, nil, errors.Wrap(errors.ErrNotFound,
				errors.NewMesssage(fmt.Sprintf("min is %f and max is %f", min, max)))
		}
		return nil, nil, errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("pair '%s/%s' not found", in.String(), out.String())))
	}

	es := &estimate{
		EstimateAmount: &entity.EstimateAmount{},
	}

	for _, est := range estimates {
		fmt.Println(est.ex.NID(), est.AmountIn, est.AmountOut)
		if est.EstimateAmount != nil && est.AmountOut > es.AmountOut {
			es = est
		}
	}

	return es.EstimateAmount, es.ex, nil
}
