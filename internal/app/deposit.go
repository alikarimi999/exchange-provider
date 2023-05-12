package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (u *OrderUseCase) SetTxId(oId *entity.ObjectId, txId string) error {
	const op = errors.Op("OrderUseCase.SetTxId")

	ord, err := u.repo.Get(oId)
	if err != nil {
		return err
	}

	if ord.STATUS() != entity.OCreated {
		return errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage(fmt.Sprintf("unable to set txId for order in '%s' status", ord.STATUS())))
	}
	ex, err := u.exs.getByNID(ord.ExchangeNid())
	if err != nil {
		return err
	}
	if ex.Type() != entity.CEX {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("unable to set transaction id for this order"))
	}
	exist, err := u.repo.TxIdExists(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("txId used before"), op, errors.ErrBadRequest)
	}

	return ex.(entity.Cex).TxIdSetted(ord, txId)
}
