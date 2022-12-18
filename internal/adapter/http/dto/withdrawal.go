package dto

import "exchange-provider/internal/entity"

type Withdrawal struct {
	Id      uint64 `json:"id"`
	OrderId int64  `json:"order_id,omitempty"`

	Status string `json:"status"`

	TxId string `json:"tx_id"`

	Address string `json:"address"`
	Tag     string `json:"tag"`

	Token string `json:"token"`

	Volume      string `json:"volume"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
}

func WFromEntity(w *entity.Withdrawal) *Withdrawal {
	return &Withdrawal{
		Id:     w.Id,
		Status: string(w.Status),

		Address: w.Addr,
		Tag:     w.Tag,

		Token:       w.Token.String(),
		Volume:      w.Volume,
		Fee:         w.Fee,
		FeeCurrency: w.FeeCurrency,

		TxId: w.TxId,
	}

}
