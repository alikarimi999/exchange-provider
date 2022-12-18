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
		return o.exs.add(ex)
	}
	return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("exchange %s  already exists!", ex.Id())))
}

func (o *OrderUseCase) GetExchange(id string) (entity.Exchange, error) {
	ex, err := o.exs.get(id)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("exchange %s not found", id)))
	}
	return ex, nil
}

func (o *OrderUseCase) AllExchanges(names ...string) []entity.Exchange {
	return o.exs.all(names...)
}

func (o *OrderUseCase) SelectExchangeByPair(in, out *entity.Token) (entity.Exchange, error) {
	exs := []entity.Exchange{}
	for _, ex := range o.exs.getAll() {
		if ex.Support(in, out) {
			exs = append(exs, ex)
		}
	}

	if len(exs) == 0 {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("no exchange support %s/%s", in.String(), out.String())))
	}

	// pick one randomly
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return exs[r.Intn(len(exs))], nil

}
