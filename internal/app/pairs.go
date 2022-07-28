package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddPairs(ex entity.Exchange, pairs []*entity.Pair) (*entity.AddPairsResult, error) {
	return ex.AddPairs(pairs)
}

func (o *OrderUseCase) SupportedPairs(exchange string) ([]*entity.Pair, error) {
	ex, exists := o.exs[exchange]
	if exists {
		return ex.GetPairs(), nil
	}
	return nil, errors.Wrap(errors.ErrNotFound, "exchange not found")
}

// check if the pair is supported by the exchange
func (o *OrderUseCase) Support(exchange string, c1, c2 *entity.Coin) (bool, error) {
	ex, exists := o.exs[exchange]
	if exists {
		return ex.Support(c1, c2), nil
	}
	return false, errors.Wrap(errors.ErrNotFound, "exchange not found")
}
