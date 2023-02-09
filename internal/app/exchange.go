package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/rand"
	"time"
)

func (o *OrderUseCase) AddExchange(ex entity.Exchange) error {
	if exists := o.exs.exists(ex.Id()); !exists {
		return o.exs.AddExchange0(ex)
	}
	return fmt.Errorf("exchange %s  already exists", ex.Id())
}

func (o *OrderUseCase) GetExchange(id string) (entity.Exchange, error) {
	return o.exs.get(id)
}

func (o *OrderUseCase) AllExchanges(names ...string) []entity.Exchange {
	return o.exs.getByNames(names...)
}

func (o *OrderUseCase) SelectExchangeByPair(in, out *entity.Token) (entity.Exchange, error) {

	exs := o.exs.getAll()
	sCexs := []entity.Exchange{}
	sDexs := []entity.Exchange{}
	for _, ex := range exs {
		if ex.Support(in, out) {
			if ex.Type() == entity.EvmDEX {
				sDexs = append(sDexs, ex)
			} else {
				sCexs = append(sCexs, ex)
			}
		}
	}

	if len(sDexs) > 0 {
		return sDexs[randInt(len(sDexs))], nil
	}

	if len(sCexs) > 0 {
		return sCexs[randInt(len(sCexs))], nil
	}

	return nil, errors.Wrap(errors.ErrNotFound)
}

func randInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
