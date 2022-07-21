package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

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

func OWToDTO(w *entity.Withdrawal) *Withdrawal {
	return &Withdrawal{
		Id:          w.Id,
		OrderId:     w.OrderId,
		UserId:      w.UserId,
		Address:     w.Address,
		Coin:        w.Coin.Symbol,
		Chain:       string(w.Coin.Chain),
		Exchange:    w.Exchange,
		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,
		TxId:        w.TxId,
		Status:      string(w.Status),
	}
}

func (w *Withdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:          w.Id,
		OrderId:     w.OrderId,
		UserId:      w.UserId,
		Address:     w.Address,
		Coin:        entity.Coin{Symbol: w.Coin, Chain: entity.Chain(w.Chain)},
		Exchange:    w.Exchange,
		Total:       w.Total,
		Fee:         w.Fee,
		ExchangeFee: w.ExchangeFee,
		Executed:    w.Executed,
		TxId:        w.TxId,
		Status:      entity.WithdrawalStatus(w.Status),
	}
}

func (w *Withdrawal) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w *Withdrawal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}

type PendingWithdrawal struct {
	Id       string
	OrderId  int64
	UserId   int64
	Coin     string
	Chain    string
	Exchange string
}

func WToDTO(w *entity.Withdrawal) *PendingWithdrawal {
	return &PendingWithdrawal{
		Id:       w.Id,
		OrderId:  w.OrderId,
		UserId:   w.UserId,
		Coin:     w.Coin.Symbol,
		Chain:    string(w.Coin.Chain),
		Exchange: w.Exchange,
	}
}

func (w *PendingWithdrawal) ToEntity() *entity.Withdrawal {
	return &entity.Withdrawal{
		Id:       w.Id,
		OrderId:  w.OrderId,
		UserId:   w.UserId,
		Coin:     entity.Coin{Symbol: w.Coin, Chain: entity.Chain(w.Chain)},
		Exchange: w.Exchange,
	}
}

func (w *PendingWithdrawal) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

func (w *PendingWithdrawal) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}
