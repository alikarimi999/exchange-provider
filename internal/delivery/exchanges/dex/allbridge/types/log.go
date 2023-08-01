package types

import (
	"math/big"
	"time"
)

type TokensReceivedLog struct {
	TxId       string
	Recipient  string
	Amount     *big.Int
	Nonce      *big.Int
	DownloadAt time.Time
}
