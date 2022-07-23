package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) selectExchange(c1, c2 *entity.Coin) (string, error) {
	for _, ex := range o.exs {
		if ex.Support(c1, c2) {
			return ex.ID(), nil
		}
	}
	return "", errors.Wrap(errors.ErrNotFound)
}
