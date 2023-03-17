package kucoin

import (
	"exchange-provider/internal/entity"
	"time"
)

type depositeRecord struct {
	// txId is the transaction id of the deposit and is used as the key in the cache.
	TxId         string
	Currency     string
	Volume       string
	Status       string
	DownloadedAt time.Time
}

func (d *depositeRecord) MatchCurrency(de *entity.Deposit) bool {
	return d.Currency == string(de.TokenId)
}

func (d *depositeRecord) snapShot() *depositeRecord {
	return &depositeRecord{
		TxId:         d.TxId,
		Currency:     d.Currency,
		Volume:       d.Volume,
		Status:       d.Status,
		DownloadedAt: d.DownloadedAt,
	}
}
