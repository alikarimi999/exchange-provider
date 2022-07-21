package dto

import (
	"encoding/json"
	"fmt"
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

func (w *Withdrawal) FixTxId() string {
	if w.WalletTxId == "" {
		return ""
	}

	if strings.Contains(w.WalletTxId, "@") {
		return strings.Split(w.WalletTxId, "@")[0]
	}

	return w.WalletTxId
}

func (w *Withdrawal) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}

func (w *Withdrawal) String() string {
	return fmt.Sprintf("%+v", *w)
}
