package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddPairs(ex entity.Exchange, pairs []*entity.Pair) (*entity.AddPairsResult, error) {
	return ex.AddPairs(pairs)
}

func (o *OrderUseCase) SupportedPairs(exchange string) ([]*entity.Pair, error) {
	ex, err := o.exs.get(exchange)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, "exchange not found")
	}

	return ex.GetPairs(), nil
}

// check if the pair is supported by the exchange
func (o *OrderUseCase) Support(exchange string, bc, qc *entity.Coin) (bool, error) {
	ex, err := o.exs.get(exchange)
	if err != nil {
		return false, errors.Wrap(errors.ErrNotFound, "exchange not found")
	}

	return ex.Support(bc, qc), nil
}
