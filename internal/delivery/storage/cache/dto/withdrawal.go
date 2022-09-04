package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type Withdrawal struct {
	Id      uint64
	WId     string
	OrderId int64
	UserId  int64

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

func OWToDTO(w *entity.Withdrawal) *Withdrawal {
	return &Withdrawal{
		Id:      w.Id,
		WId:     w.WId,
		OrderId: w.OrderId,
		UserId:  w.UserId,

		Address: w.Addr,
		Tag:     w.Tag,

		Exchange:    w.Exchange,
		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,
		TxId:        w.TxId,
		Status:      string(w.Status),
		FailedDesc:  w.FailedDesc,
	}
}

func (w *Withdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:      w.Id,
		WId:     w.WId,
		OrderId: w.OrderId,
		UserId:  w.UserId,

		Address: &entity.Address{Addr: w.Address, Tag: w.Tag},

		Exchange:    w.Exchange,
		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,
		TxId:        w.TxId,
		Status:      entity.WithdrawalStatus(w.Status),
		FailedDesc:  w.FailedDesc,
	}
}

func (w *Withdrawal) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w *Withdrawal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}

type PendingWithdrawal struct {
	Id       uint64
	WId      string
	OrderId  int64
	UserId   int64
	Coin     string
	Chain    string
	Exchange string
}

func WToDTO(w *entity.Withdrawal) *PendingWithdrawal {
	return &PendingWithdrawal{
		Id:      w.Id,
		WId:     w.WId,
		OrderId: w.OrderId,
		UserId:  w.UserId,

		Exchange: w.Exchange,
	}
}

func (w *PendingWithdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:       w.Id,
		WId:      w.WId,
		OrderId:  w.OrderId,
		UserId:   w.UserId,
		Exchange: w.Exchange,
	}
}

func (w *PendingWithdrawal) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w *PendingWithdrawal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}
