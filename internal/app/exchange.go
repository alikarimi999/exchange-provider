package app

import (
	"fmt"
	"order_service/internal/delivery/exchanges/kucoin"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddExchange(exchange entity.Exchange, cfg interface{}) error {
	const op = errors.Op("OrderUseCase.AddExchange")
	if !o.exStore.exists(exchange.ID()) {
		ex, err := exchange.Setup(cfg, o.rc, o.l)
		if err != nil {
			return errors.Wrap(op, errors.New(fmt.Sprintf("exchange %s (err: %s )", exchange.ID(), err.Error())))
		}

		o.exStore.add(ex)
		return nil
	}
	return errors.Wrap(errors.ErrBadRequest, errors.New(fmt.Sprintf("exchange %s already exists", exchange.ID())))

}

func (o *OrderUseCase) ChangeExchangeAccount(id string, cfg *kucoin.Configs) error {
	ex, err := o.exStore.get(id)
	if err != nil {
		return err
	}

	if err := ex.ChangeAccount(cfg); err != nil {
		return err
	}
	o.exStore.update(ex)
	return nil

}

func (o *OrderUseCase) GetExchange(exchange string) (entity.Exchange, error) {
	return o.exStore.get(exchange)
}

func (o *OrderUseCase) ExchangeExists(exchange string) bool {
	return o.exStore.exists(exchange)
}

func (o *OrderUseCase) selectExchange(c1, c2 *entity.Coin) (string, error) {
	for _, ex := range o.exs {
		if ex.Support(c1, c2) {
			return ex.ID(), nil
		}
	}
	return "", errors.Wrap(errors.ErrNotFound)
}
