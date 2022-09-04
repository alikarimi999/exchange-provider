package dto

import (
	"order_service/internal/entity"
)

type ExchangeOrder struct {
	Id      uint64 `gorm:"primary_key"`
	ExId    string
	UserId  int64
	OrderId int64
	Status  string

	Symbol      string
	Exchange    string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string

	FailedDesc string
}

func EToDto(eo *entity.ExchangeOrder) *ExchangeOrder {
	if eo == nil {
		return &ExchangeOrder{}
	}

	return &ExchangeOrder{
		Id:          uint64(eo.Id),
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

func (eo *ExchangeOrder) ToEntity() *entity.ExchangeOrder {
	if eo == nil {
		return &entity.ExchangeOrder{}
	}
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

		FailedDesc: eo.FailedDesc,
	}

}
