package dto

import (
	"exchange-provider/internal/entity"
)

type Swap struct {
	Id     uint64 `json:"id"`
	Status string `json:"status"`
	TxId   string `json:"txId"`

	Exchange uint   `json:"exchange"`
	Input    string `json:"input"`
	Output   string `json:"output"`

	InAmount  string `json:"inAmount"`
	OutAmount string `json:"outAmount"`

	Fee             string `json:"fee"`
	FeeCurrency     string `json:"feeCurrency"`
	entity.MetaData `json:"metaData,omitempty"`
}

func swapFromEntity(e *entity.Swap, r *entity.Route) *Swap {
	return &Swap{
		Id:   e.Id,
		TxId: e.TxId,

		Exchange: r.Exchange,
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
