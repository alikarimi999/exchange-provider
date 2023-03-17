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

	exist, err := o.repo.TxIdExists(txId)
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	if exist {
		return errors.Wrap(errors.NewMesssage("tx id used before"), op, errors.ErrBadRequest)
	}
	ex, err := o.GetExchange(ord.Routes[0].Exchange)
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
