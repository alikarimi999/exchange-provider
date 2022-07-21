package dto

import (
	"encoding/json"
	"fmt"
	"order_service/internal/entity"
	"strings"
)

type Withdrawal struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	WalletTxId string `json:"walletTxId"`
	IsInner    bool   `json:"isInner"`
	Status     string `json:"status"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

func (w *Withdrawal) ToEntity() *entity.Withdrawal {
	ew := &entity.Withdrawal{
		Id:       w.Id,
		Exchange: "kucoin",
		TxId:     strings.Split(w.WalletTxId, "@")[0],
	}

	switch w.Status {
	case "SUCCESS":
		ew.Status = entity.WithdrawalSucceed
	case "FAILURE":
		ew.Status = entity.WithdrawalFailed
	default:
		ew.Status = entity.WithdrawalPending
	}

	return ew
}

func (w *Withdrawal) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}

func (w *Withdrawal) String() string {
	return fmt.Sprintf("%+v", *w)
}
