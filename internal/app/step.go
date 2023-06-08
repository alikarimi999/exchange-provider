package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u OrderUseCase) GetMultiStep(o entity.Order, step uint) (entity.Tx, error) {
	ex, err := u.exs.getByNID(o.ExchangeNid())
	if err != nil {
		return nil, err
	}

	if o.Expire() {
		return nil, errors.Wrap(errors.ErrForbidden, errors.NewMesssage("order has expired"))
	}

	return ex.(entity.EVMDex).CreateTx(o)

}
