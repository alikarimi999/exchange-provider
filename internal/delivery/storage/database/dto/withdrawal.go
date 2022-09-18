package dto

import (
	"exchange-provider/internal/entity"
)

type Withdrawal struct {
	Id      uint64 `gorm:"primary_key"`
	WId     string
	UserId  int64
	OrderId int64

	Status  string
	Address string
	Tag     string

	Coin     string
	Chain    string
	Exchange string

	Total       string
	Fee         string
	ExchangeFee string
	Executed    string

	TxId       string
	FailedDesc string
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

		Coin:     w.CoinId,
		Chain:    w.ChainId,
		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:       w.TxId,
		Status:     string(w.Status),
		FailedDesc: w.FailedDesc,
	}
}

func (w *Withdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:      w.Id,
		WId:     w.WId,
		UserId:  w.UserId,
		OrderId: w.OrderId,

		Address: &entity.Address{Addr: w.Address, Tag: w.Tag},

		Coin: &entity.Coin{
			CoinId:  w.Coin,
			ChainId: w.Chain,
		},
		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:       w.TxId,
		Status:     entity.WithdrawalStatus(w.Status),
		FailedDesc: w.FailedDesc,
	}
}
