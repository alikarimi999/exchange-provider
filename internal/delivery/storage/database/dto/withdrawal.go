package dto

import (
	"order_service/internal/entity"
)

type Withdrawal struct {
	Id      uint64 `gorm:"primary_key"`
	WId     string
	UserId  int64
	OrderId int64

	Address string
	Tag     string

	Coin     string
	Chain    string
	Exchange string

	Total       string
	Fee         string
	ExchangeFee string
	Executed    string

	TxId   string
	Status string
}

func WToDto(w *entity.Withdrawal) *Withdrawal {

	if w == nil {
		return &Withdrawal{}
	}

	return &Withdrawal{
		Id:      w.Id,
		WId:     w.WId,
		UserId:  w.UserId,
		OrderId: w.OrderId,

		Address: w.Addr,
		Tag:     w.Tag,

		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:   w.TxId,
		Status: string(w.Status),
	}
}

func (w *Withdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:      w.Id,
		WId:     w.WId,
		UserId:  w.UserId,
		OrderId: w.OrderId,

		Address: &entity.Address{Addr: w.Address, Tag: w.Tag},

		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:   w.TxId,
		Status: entity.WithdrawalStatus(w.Status),
	}
}
