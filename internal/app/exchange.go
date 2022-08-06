package app

import (
	"fmt"
	"math/rand"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"time"
)

func (o *OrderUseCase) AddExchange(exc entity.Exchange) error {
	const op = errors.Op("OrderUseCase.AddExchange")
	exists, status := o.exs.exists(exc.NID())
	if !exists {
		ex, err := exc.Setup(o.rc, o.l)
		if err != nil {
			return errors.Wrap(op, errors.New(fmt.Sprintf("exchange %s (err: %s )", exc.NID(), err.Error())))
		}

		o.exs.add(ex)
		return nil
	}
	return errors.Wrap(errors.ErrBadRequest, errors.New(fmt.Sprintf("exchange %s with accountId %s already exists and it's status is %s", exc.Name(), exc.AccountId(), status)))

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

func (o *OrderUseCase) SelectExchangeByPair(bc, qc *entity.Coin) (string, error) {
	exs := []string{}
	for _, ex := range o.exs.getActives() {
		if ex.Support(bc, qc) {
			exs = append(exs, ex.NID())
		}
	}

	if len(exs) == 0 {
		return "", errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("no exchange support %s/%s", bc.String(), qc.String())))
	}

	// pick one randomly
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return exs[r.Intn(len(exs))], nil

}
