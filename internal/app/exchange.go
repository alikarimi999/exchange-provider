package app

import (
	"fmt"
	"order_service/internal/delivery/exchanges/kucoin"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddExchange(exchange entity.Exchange, cfg interface{}) error {
	const op = errors.Op("OrderUseCase.AddExchange")
	if !o.exs.exists(exchange.ID()) {
		ex, err := exchange.Setup(cfg, o.rc, o.l)
		if err != nil {
			return errors.Wrap(op, errors.New(fmt.Sprintf("exchange %s (err: %s )", exchange.ID(), err.Error())))
		}

		o.exs.add(ex)
		return nil
	}
	return errors.Wrap(errors.ErrBadRequest, errors.New(fmt.Sprintf("exchange %s already exists", exchange.ID())))

}

func (o *OrderUseCase) ChangeExchangeAccount(id string, cfg *kucoin.Configs) error {
	ex, err := o.exs.get(id)
	if err != nil {
		return err
	}

	if err := ex.ChangeAccount(cfg); err != nil {
		return err
	}
	o.exs.update(ex)
	return nil

}

func (o *OrderUseCase) GetExchange(exchange string) (entity.Exchange, error) {
	return o.exs.get(exchange)
}

func (o *OrderUseCase) GetAllExchanges() []string {
	return o.exs.getAll()
}

func (o *OrderUseCase) ExchangeExists(exchange string) bool {
	return o.exs.exists(exchange)
}

func (o *OrderUseCase) selectExchange(bc, qc *entity.Coin) (string, error) {
	for _, ex := range o.exs.exs {
		if ex.Support(bc, qc) {
			return ex.ID(), nil
		}
	}
	return "", errors.Wrap(errors.ErrNotFound)
}
