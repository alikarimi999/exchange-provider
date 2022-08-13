package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type Deposit struct {
	Id         int64  `json:"deposit_id"`
	UserId     int64  `json:"user_id,omitempty"`
	OrderId    int64  `json:"order_id,omitempty"`
	Exchange   string `json:"exchange,omitempty"`
	Volume     string `json:"volume"`
	Fullfilled bool   `json:"fullfilled"`
	Address    string `json:"address"`
}

func DFromEntity(d *entity.Deposit) *Deposit {
	return &Deposit{
		Id: d.Id,

		Volume:     d.Volume,
		Fullfilled: d.Fullfilled,
		Address:    d.Address,
	}
}

type SetTxIdRequest struct {
	OrderId   int64  `json:"order_id"`
	DepositId int64  `json:"deposit_id"`
	TxId      string `json:"tx_id"`
}

func (r *SetTxIdRequest) Validate() error {
	if r.OrderId == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order_id is required"))
	}
	if r.DepositId == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("deposit_id is required"))
	}
	if r.TxId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tx_id is required"))
	}
	return nil
}
