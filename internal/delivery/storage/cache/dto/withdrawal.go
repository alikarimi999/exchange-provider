package dto

import (
	"encoding/json"
	"exchange-provider/internal/entity"
)

type Withdrawal struct {
	Id      uint64
	WId     string
	OrderId int64
	UserId  int64

	Status  string
	Address string
	Tag     string

	Coin      string
	Chain     string
	Unwrapped bool

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

		Coin:      w.CoinId,
		Chain:     w.ChainId,
		Unwrapped: w.Unwrapped,

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

		Coin:      &entity.Coin{CoinId: w.Coin, ChainId: w.Chain},
		Unwrapped: w.Unwrapped,

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
