package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"time"
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

	if ord.Status == entity.OExpired {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order expired"))
	}

	if ord.Deposit.ExpireAt > 0 && time.Now().Unix() >= ord.Deposit.ExpireAt {
		ord.Status = entity.OExpired
		o.write(ord)
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order expired"))
	}

	exist, err := o.repo.TxIdExists(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("tx id used before"), op, errors.ErrBadRequest)
	}
	ex, err := o.exs.getByName(ord.Routes[0].Exchange)
	if err != nil {
		return err
	}
	ord.Deposit.TxId = txId
	ord.Deposit.Status = entity.DepositTxIdSet
	ord.Status = entity.OConfimDeposit
	ord.UpdatedAt = time.Now().Unix()
	if err := o.write(ord); err != nil {
		return err
	}

	cex := ex.(entity.Cex)
	go cex.TxIdSetted(ord)
	return nil
}
