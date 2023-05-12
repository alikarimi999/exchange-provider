package dto

import (
	"strings"
	"time"
)

type Withdrawal struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	WalletTxId   string `json:"walletTxId"`
	IsInner      bool   `json:"isInner"`
	Status       string `json:"status"`
	Remark       string `json:"remark"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DownloadedAt time.Time
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
