package dto

import (
	"exchange-provider/internal/entity"
)

type Swap struct {
	Id      uint64 `json:"id"`
	OrderId int64  `json:"order_id,omitempty"`
	Status  string `json:"status"`
	TxId    string `json:"tx_id,omitempty"`

	Index    int    `json:"index,omitempty"`
	Input    string `json:"input,omitempty"`
	InAmount string `json:"inAmount,omitempty"`

	Output    string `json:"output,omitempty"`
	OutAmount string `json:"outAmount,omitempty"`

	FilledPrice     string `json:"filled_price,omitempty"`
	Fee             string `json:"fee"`
	FeeCurrency     string `json:"fee_currency"`
	entity.MetaData `json:"meta_data,omitempty"`
}

func SwapFromEntity(e *entity.Swap, r *entity.Route, index int) *Swap {
	return &Swap{
		Id:      e.Id,
		OrderId: e.OrderId,
		TxId:    e.TxId,

		Index:    index,
		Input:    r.In.String(),
		InAmount: e.InAmount,

		Output:    r.Out.String(),
		OutAmount: e.OutAmount,

		Fee:         e.Fee,
		FeeCurrency: e.FeeCurrency,
		Status:      string(e.Status),
		MetaData:    e.MetaData,
	}
}
