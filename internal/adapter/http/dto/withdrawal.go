package dto

import "order_service/internal/entity"

type Withdrawal struct {
	Id       uint64 `json:"id"`
	WId      string `json:"exchange_withdrawal_id"`
	OrderId  int64
	UserId   int64
	Exchange string

	Address string

	Coin  string
	Chain string

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
		WId:     w.WId,
		OrderId: w.OrderId,
		UserId:  w.UserId,
		Address: w.Address,

		Exchange: w.Exchange,

		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,

		TxId:   w.TxId,
		Status: string(w.Status),
	}
}
