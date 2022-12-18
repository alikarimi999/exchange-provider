package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type Deposit struct {
	Id int64 `json:"id"`

	Status string `json:"status"`
	Token  string `json:"token"`

	TxId   string `json:"tx_id"`
	Volume string `json:"volume"`

	Address string `json:"address"`
	Tag     string `json:"tag"`

	FailedDesc string `json:"failed_descritpion,omitempty"`
}

func DFromEntity(d *entity.Deposit) *Deposit {
	return &Deposit{
		Id: d.Id,

		Status: d.Status,
		Token:  d.Token.String(),

		TxId:   d.TxId,
		Volume: d.Volume,

		Address: d.Addr,
		Tag:     d.Tag,

		FailedDesc: d.FailedDesc,
	}
}

type SetTxIdRequest struct {
	Id   int64  `json:"order_id"`
	TxId string `json:"tx_id"`
	Msg  string `json:"message"`
}

func (r *SetTxIdRequest) Validate() error {
	if r.Id == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order_id is required"))
	}

	if r.TxId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tx_id is required"))
	}
	return nil
}
