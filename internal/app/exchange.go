package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/rand"
	"time"
)

func (o *OrderUseCase) AddExchange(ex entity.Exchange) error {

	exists, status := o.exs.exists(ex.NID())
	if !exists {
		return o.exs.add(&Exchange{
			Exchange:       ex,
			CurrentStatus:  ExchangeStatusActive,
			LastChangeTime: time.Now(),
		})
	}
	return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("exchange %s  already exists and it's status is %s", ex.NID(), status)))

}

func (o *OrderUseCase) GetExchange(nid string) (*Exchange, error) {
	ex, err := o.exs.get(nid)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("exchange %s not found", nid)))
	}
	return ex, nil
}

func (o *OrderUseCase) GetAllExchangesList() []string {
	return o.exs.getAllList()
}

func (o *OrderUseCase) AllExchanges(names ...string) []*Exchange {
	return o.exs.all(names...)
}

func (o *OrderUseCase) GetAllActivesExchanges() []*Exchange {
	return o.exs.getActives()
}

func (o *OrderUseCase) GetAllDeactivesExchanges() []*Exchange {
	return o.exs.getDeactives()
}

func (o *OrderUseCase) SelectExchangeByPair(in, out *entity.Coin) (entity.Exchange, error) {
	exs := []entity.Exchange{}
	for _, ex := range o.exs.getActives() {
		if ex.Support(in, out) {
			exs = append(exs, ex.Exchange)
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
