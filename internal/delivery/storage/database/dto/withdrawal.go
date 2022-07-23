package dto

import (
	"order_service/internal/entity"
)

type Withdrawal struct {
	Id      string `gorm:"primary_key"`
	UserId  int64
	OrderId int64
	Address string

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
	return &Withdrawal{
		Id:      w.Id,
		UserId:  w.UserId,
		OrderId: w.OrderId,
		Address: w.Address,

		Coin:     w.Coin.Id,
		Chain:    w.Coin.Chain.Id,
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
		UserId:  w.UserId,
		OrderId: w.OrderId,
		Address: w.Address,

		Coin: &entity.Coin{
			Id:    w.Coin,
			Chain: &entity.Chain{Id: w.Chain},
		},
		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:   w.TxId,
		Status: entity.WithdrawalStatus(w.Status),
	}
}
