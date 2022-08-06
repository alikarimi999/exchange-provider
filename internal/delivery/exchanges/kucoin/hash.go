package kucoin

import (
	"crypto/sha1"
	"encoding/hex"
)

func hash(apiKey, apiSecret, passphrase string) string {

	s := apiKey + apiSecret + passphrase
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))

}
