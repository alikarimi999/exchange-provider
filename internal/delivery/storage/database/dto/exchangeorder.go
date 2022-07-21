package dto

import "order_service/internal/entity"

type ExchangeOrder struct {
	Id          string `gorm:"primary_key"`
	UserId      int64
	OrderId     int64
	Symbol      string
	Exchange    string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string
	Status      string
}

func EToDto(eo *entity.ExchangeOrder) *ExchangeOrder {
	return &ExchangeOrder{
		Id:          eo.Id,
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
	}
}

func (eo *ExchangeOrder) ToEntity() *entity.ExchangeOrder {
	return &entity.ExchangeOrder{
		Id:          eo.Id,
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
	}
}
