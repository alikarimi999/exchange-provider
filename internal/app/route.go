package app

import (
	"exchange-provider/internal/entity"
)

func (o *OrderUseCase) routing(in, out *entity.Token, amount float64) (map[int]*entity.Route, error) {
	ex, _, err := o.estimateAmountOut(in, out, amount, 0)
	if err != nil {
		return nil, err
	}
	routes := make(map[int]*entity.Route)
	routes[0] = &entity.Route{In: in, Out: out, Exchange: ex.Name(), ExType: ex.Type()}
	return routes, nil
}
