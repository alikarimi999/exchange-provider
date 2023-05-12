package app

import (
	"exchange-provider/internal/entity"
)

func (u OrderUseCase) GetMultiStep(o entity.Order, step uint) (entity.Tx, error) {
	ex, err := u.exs.getByNID(o.ExchangeNid())
	if err != nil {
		return nil, err
	}

	return ex.(entity.EVMDex).CreateTx(o)

}
