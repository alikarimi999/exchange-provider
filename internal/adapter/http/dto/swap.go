package dto

import (
	"exchange-provider/internal/entity"
)

type Swap struct {
	Id      uint64 `json:"id"`
	OrderId int64  `json:"order_id"`
	Status  string `json:"status"`
	TxId    string `json:"tx_id"`

	Index  int    `json:"index"`
	Input  string `json:"input"`
	Output string `json:"output"`

	InAmount  string `json:"inAmount"`
	OutAmount string `json:"outAmount"`

	FilledPrice     string `json:"filled_price"`
	Fee             string `json:"fee"`
	FeeCurrency     string `json:"fee_currency"`
	entity.MetaData `json:"meta_data,omitempty"`
}

func swapFromEntity(e *entity.Swap, r *entity.Route) *Swap {
	return &Swap{
		Id:      e.Id,
		OrderId: e.OrderId,
		TxId:    e.TxId,

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
