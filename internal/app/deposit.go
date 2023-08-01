package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u *OrderUseCase) SetTxId(oId *entity.ObjectId, txId string) error {
	const op = errors.Op("OrderUseCase.SetTxId")

	ord, err := u.repo.Get(oId)
	if err != nil {
		return err
	}

	ex, err := u.exs.GetByNID(ord.ExchangeNid())
	if err != nil {
		return err
	}

	exist, err := u.repo.TxIdExists(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("txId used before"), op, errors.ErrBadRequest)
	}

	return ex.SetTxId(ord, txId)
}
