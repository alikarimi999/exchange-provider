package app

import (
	"order_service/internal/entity"
)

func (o *OrderUseCase) AddPairs(ex *Exchange, pairs []*entity.Pair) (*entity.AddPairsResult, error) {

	return ex.AddPairs(pairs)
}

func (o *OrderUseCase) GetAllPairs(ex *Exchange) ([]*entity.Pair, error) {
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
