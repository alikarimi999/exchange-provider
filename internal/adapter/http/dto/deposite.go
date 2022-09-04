package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type Deposit struct {
	Id int64 `json:"id"`

	Status   string `json:"status"`
	Exchange string `json:"exchange,omitempty"`
	Coin     string `json:"coin"`

	TxId   string `json:"tx_d"`
	Volume string `json:"volume"`

	Address string `json:"address"`
	Tag     string `json:"tag"`

	FailedDesc string `json:"failed_descritpion"`
}

func DFromEntity(d *entity.Deposit) *Deposit {
	return &Deposit{
		Id: d.Id,

		Status:   d.Status,
		Exchange: d.Exchange,
		Coin:     d.Coin.String(),

		TxId:   d.TxId,
		Volume: d.Volume,

		Address: d.Addr,
		Tag:     d.Tag,

		FailedDesc: d.FailedDesc,
	}
}

type SetTxIdRequest struct {
	Seq  int64  `json:"order_id"`
	TxId string `json:"tx_id"`
	Msg  string `json:"message"`
}

func (r *SetTxIdRequest) Validate() error {
	if r.Seq == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order_id is required"))
	}

	if r.TxId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tx_id is required"))
	}
	return nil
}
