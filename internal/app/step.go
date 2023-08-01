package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u OrderUseCase) GetMultiStep(o entity.Order, step uint) (entity.Tx, error) {
	ex, err := u.exs.GetByNID(o.ExchangeNid())
	if err != nil {
		return nil, err
	}
	if o.Expire() {
		return nil, errors.Wrap(errors.ErrForbidden, errors.NewMesssage("order has expired"))
	}

	switch ex.Type() {
	case entity.EvmDEX:
		return ex.(entity.EVMDex).CreateTx(o)
	case entity.CrossDex:
		return ex.(entity.CrossDEX).CreateTx(o, int(step))
	default:
		return nil, errors.Wrap(errors.ErrBadRequest)
	}
}
