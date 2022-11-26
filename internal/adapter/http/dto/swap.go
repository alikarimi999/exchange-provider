package dto

import (
	"exchange-provider/internal/entity"
)

type Swap struct {
	Id       uint64
	Ex_Id    string `json:"exchange_id"`
	UserId   int64  `json:"user_id,omitempty"`
	OrderId  int64  `json:"order_id,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	Input    string `json:"input,omitempty"`
	Output   string `json:"output,omitempty"`

	FilledPrice string `json:"filled_price,omitempty"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	Status      string `json:"status"`
}

func SwapFromEntity(e *entity.Swap) *Swap {
	ex := &Swap{
		Id:    e.Id,
		Ex_Id: e.ExId,

		Input:       e.InAmount,
		Output:      e.OutAmount,
		Fee:         e.Fee,
		FeeCurrency: e.FeeCurrency,
		Status:      string(e.Status),
	}

	// if ex.Input != "" && ex.Output != "" {
	// 	s, _ := strconv.ParseFloat(ex.Size, 64)
	// 	f, _ := strconv.ParseFloat(ex.Funds, 64)
	// 	ex.FilledPrice = strconv.FormatFloat(f/s, 'f', 8, 64)
	// }

	return ex
}
