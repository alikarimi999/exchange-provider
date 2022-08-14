package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (o *OrderUseCase) SetTxId(userId, orderId, depositeId int64, txId string) error {
	const op = errors.Op("OrderUseCase.SetTxId")

	ord := &entity.UserOrder{
		UserId: userId,
		Id:     orderId,
	}
	if err := o.read(ord); err != nil {
		if errors.ErrorCode(err) == errors.ErrNotFound {
			return err
		}
		o.l.Error(string(op), err.Error())
		return errors.Wrap(errors.ErrInternal)
	}

	if err := o.DS.SetTxId(userId, orderId, depositeId, txId); err != nil {
		o.l.Error(string(op), err.Error())
		return errors.Wrap(errors.ErrInternal)
	}
	return nil
}
