package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) SetTxId(userId, seq int64, txId string) error {
	const op = errors.Op("OrderUseCase.SetTxId")

	ord := &entity.UserOrder{
		UserId: userId,
		Seq:    seq,
	}
	if err := o.read(ord); err != nil {
		if errors.ErrorCode(err) == errors.ErrNotFound {
			return err
		}
		o.l.Error(string(op), err.Error())
		return errors.Wrap(errors.ErrInternal)
	}

	if ord.Deposite.TxId != "" {
		return errors.Wrap(errors.NewMesssage("order already has tx id"))
	}

	exist, err := o.repo.CheckTxId(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("tx_id used before"), op, errors.ErrBadRequest)
	}

	ord.Deposite.TxId = txId
	if err := o.write(ord); err != nil {
		return err
	}

	if err := o.DS.SetTxId(ord.Deposite.Id, txId); err != nil {
		return err
	}

	return nil
}
