package dto

import "order_service/internal/entity"

type Withdrawal struct {
	Id      string
	OrderId int64
	UserId  int64
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

func WFromEntity(w *entity.Withdrawal) *Withdrawal {
	return &Withdrawal{
		Id:      w.Id,
		OrderId: w.OrderId,
		UserId:  w.UserId,
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
