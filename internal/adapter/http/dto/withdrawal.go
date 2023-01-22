package dto

import "exchange-provider/internal/entity"

type Withdrawal struct {
	Id      string `json:"id"`
	OrderId string `json:"orderId,omitempty"`

	Status string `json:"status"`

	TxId string `json:"txId"`

	Address string `json:"address"`
	Tag     string `json:"tag"`

	Token string `json:"token"`

	Volume      string `json:"volume"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"feeCurrency"`
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
