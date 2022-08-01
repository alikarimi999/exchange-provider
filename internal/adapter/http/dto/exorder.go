package dto

import (
	"order_service/internal/entity"
	"strconv"
)

type ExchangeOrder struct {
	Id          uint64
	Ex_Id       string `json:"exchange_id"`
	UserId      int64  `json:"user_id"`
	OrderId     int64  `json:"order_id"`
	Symbol      string `json:"symbol"`
	Exchange    string `json:"exchange"`
	Side        string `json:"side"`
	Funds       string `json:"funds"`
	Size        string `json:"size"`
	FilledPrice string `json:"filled_price"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	Status      string `json:"status"`
}

func EoFromEntity(e *entity.ExchangeOrder) *ExchangeOrder {
	ex := &ExchangeOrder{
		Id:          e.Id,
		Ex_Id:       e.ExId,
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

	if ex.Funds != "" && ex.Size != "" {
		s, _ := strconv.ParseFloat(ex.Size, 64)
		f, _ := strconv.ParseFloat(ex.Funds, 64)
		ex.FilledPrice = strconv.FormatFloat(f/s, 'f', 8, 64)
	}

	return ex
}
