package dto

import "order_service/internal/entity"

type ExchangeOrder struct {
	Id          string
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

func EoFromEntity(e *entity.ExchangeOrder) *ExchangeOrder {
	return &ExchangeOrder{
		Id:          e.Id,
		UserId:      e.UserId,
		OrderId:     e.OrderId,
		Symbol:      e.Symbol,
		Exchange:    e.Exchange,
		Side:        e.Side,
		Funds:       e.Funds,
		Size:        e.Size,
		Fee:         e.Fee,
		FeeCurrency: e.FeeCurrency,
		Status:      string(e.Status),
	}
}
