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

func (w *Withdrawal) SnapShot() *Withdrawal {
	return &Withdrawal{
		Id:           w.Id,
		Address:      w.Address,
		Currency:     w.Currency,
		Amount:       w.Amount,
		Fee:          w.Fee,
		WalletTxId:   w.WalletTxId,
		IsInner:      w.IsInner,
		Status:       w.Status,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
		DownloadedAt: w.DownloadedAt,
	}
}
