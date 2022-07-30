package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) AddPairs(ex entity.Exchange, pairs []*entity.Pair) (*entity.AddPairsResult, error) {
	return ex.AddPairs(pairs)
}

func (o *OrderUseCase) GetAllPairs(exchange string) ([]*entity.Pair, error) {
	ex, err := o.exs.get(exchange)
	if err != nil {
		return nil, errors.Wrap(errors.ErrNotFound, "exchange not found")
	}

	return o.setFee(ex.GetAllPairs()...), nil
}

func (o *OrderUseCase) GetPair(ex entity.Exchange, bc, qc *entity.Coin) (*entity.Pair, error) {
	p, err := ex.GetPair(bc, qc)
	if err != nil {
		return nil, err
	}

	return o.setFee(p)[0], nil
}

func (o *OrderUseCase) setFee(ps ...*entity.Pair) []*entity.Pair {
	f := o.fs.GetFee()
	for _, p := range ps {
		p.Fee = f
	}
	return ps
}

func (o *OrderUseCase) RemovePair(ex entity.Exchange, bc, qc *entity.Coin) error {
	return ex.RemovePair(bc, qc)
}
