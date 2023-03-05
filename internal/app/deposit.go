package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (o *OrderUseCase) SetTxId(oId *entity.ObjectId, txId string) error {
	const op = errors.Op("OrderUseCase.SetTxId")

	ord := &entity.CexOrder{
		ObjectId: oId,
	}

	if err := o.read(ord); err != nil {
		return err
	}

	if ord.Deposit.TxId != "" {
		return errors.Wrap(errors.NewMesssage("order already has tx id"))
	}

	exist, err := o.repo.TxIdExists(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("tx id used before"), op, errors.ErrBadRequest)
	}

	ord.Deposit.TxId = txId
	ord.Deposit.Status = entity.DepositTxIdSet
	ord.Status = entity.OConfimDeposit
	if err := o.write(ord); err != nil {
		return err
	}

	go o.oh.handle(ord)
	return nil
}
