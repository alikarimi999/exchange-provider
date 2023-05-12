package kucoin

import (
	"fmt"
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

func (d *depositRecord) MatchCurrency(t *Token) error {
	if !(d.Currency == string(t.Currency)) {
		return fmt.Errorf("currency mismatch,`%s`:`%s` ",
			t.Currency, d.Currency)
	}
	return nil
}

func (d *depositRecord) snapshot() *depositRecord {
	return &depositRecord{
		TxId:         d.TxId,
		Currency:     d.Currency,
		Volume:       d.Volume,
		Status:       d.Status,
		DownloadedAt: d.DownloadedAt,
	}
}
