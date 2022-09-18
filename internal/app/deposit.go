package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
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

	if ord.Deposit.TxId != "" {
		return errors.Wrap(errors.NewMesssage("order already has tx id"))
	}

	exist, err := o.repo.CheckTxId(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("tx id used before"), op, errors.ErrBadRequest)
	}

	ord.Deposit.TxId = txId
	ord.Deposit.Status = entity.DepositTxIdSet
	ord.Status = entity.OSTxIdSetted
	if err := o.write(ord); err != nil {
		return err
	}

	o.dh.dCh <- ord.Deposit

	return nil
}
