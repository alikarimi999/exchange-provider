package kucoin

import (
	"time"
)

type depositRecord struct {
	// txId is the transaction id of the deposit and is used as the key in the cache.
	TxId         string
	Currency     string
	Volume       string
	Status       string
	DownloadedAt time.Time
}

func (d *depositRecord) MatchCurrency(t *Token) bool {
	return d.Currency == string(t.Currency)
}

func (d *depositRecord) snapShot() *depositRecord {
	return &depositRecord{
		TxId:         d.TxId,
		Currency:     d.Currency,
		Volume:       d.Volume,
		Status:       d.Status,
		DownloadedAt: d.DownloadedAt,
	}
}
