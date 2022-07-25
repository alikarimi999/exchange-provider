package dto

import "order_service/internal/entity"

type ExchangeOrder struct {
	Id          string `json:"id"`
	UserId      int64  `json:"user_id"`
	OrderId     int64  `json:"order_id"`
	Symbol      string `json:"symbol"`
	Exchange    string `json:"exchange"`
	Side        string `json:"side"`
	Funds       string `json:"funds"`
	Size        string `json:"size"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	Status      string `json:"status"`
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
