package dto

import (
	"encoding/json"
	"order_service/internal/entity"
)

type exchangeOrder struct {
	Id          uint64
	ExId        string
	UserId      int64
	OrderId     int64
	Status      string
	Exchange    string
	Symbol      string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string
	FailedDesc  string
}

func eoToDto(eo *entity.ExchangeOrder) *exchangeOrder {
	return &exchangeOrder{
		Id:          eo.Id,
		ExId:        eo.ExId,
		UserId:      eo.UserId,
		OrderId:     eo.OrderId,
		Symbol:      eo.Symbol,
		Exchange:    eo.Exchange,
		Side:        eo.Side,
		Funds:       eo.Funds,
		Size:        eo.Size,
		Fee:         eo.Fee,
		FeeCurrency: eo.FeeCurrency,
		Status:      string(eo.Status),

		FailedDesc: eo.FailedDesc,
	}
}

func (eo *exchangeOrder) ToEntity() *entity.ExchangeOrder {
	return &entity.ExchangeOrder{
		Id:          eo.Id,
		ExId:        eo.ExId,
		UserId:      eo.UserId,
		OrderId:     eo.OrderId,
		Symbol:      eo.Symbol,
		Exchange:    eo.Exchange,
		Side:        eo.Side,
		Funds:       eo.Funds,
		Size:        eo.Size,
		Fee:         eo.Fee,
		FeeCurrency: eo.FeeCurrency,
		Status:      entity.ExOrderStatus(eo.Status),
		FailedDesc:  eo.FailedDesc,
	}
}

func (e *exchangeOrder) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
