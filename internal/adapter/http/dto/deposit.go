package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type Deposit struct {
	Id string `json:"id"`

	Status string `json:"status"`
	Token  string `json:"token"`

	TxId   string `json:"txId"`
	Volume string `json:"volume"`

	Address string `json:"address"`
	Tag     string `json:"tag"`

	FailedDesc string `json:"failedDescritpion,omitempty"`
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
	Id   string `json:"orderId"`
	TxId string `json:"txId"`
	Msg  string `json:"message"`
}

func (r *SetTxIdRequest) Validate() error {
	if r.Id == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("orderId is required"))
	}

	if r.TxId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("txId is required"))
	}
	return nil
}
