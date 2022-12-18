package kucoin

import (
	"encoding/json"
	"exchange-provider/internal/entity"
)

type depositeRecord struct {
	// txId is the transaction id of the deposit and is used as the key in the cache.
	TxId     string `json:"-"`
	Currency string `json:"currency"`
	Volume   string `json:"volume"`
	Status   string `json:"status"`
}

// implement `encoding.BinaryMarshaler` for saving in redis cache
func (d *depositeRecord) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *depositeRecord) MatchCurrency(de *entity.Deposit) bool {
	return d.Currency == string(de.TokenId)
}
