package app

import (
	"exchange-provider/internal/entity"
)

func (o *OrderUseCase) AddPairs(ex entity.Exchange, data interface{}) (*entity.AddPairsResult, error) {
	return ex.AddPairs(data)
}

func (o *OrderUseCase) RemovePair(exId uint, t1, t2 *entity.Token) error {
	ex, err := o.exs.get(exId)
	if err != nil {
		return err
	}
	return ex.RemovePair(t1, t2)
}
