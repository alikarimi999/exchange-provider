package dto

import "time"

type Coin struct {
	CoinId  string
	ChainId string

	BlockTime           time.Duration
	WithdrawalPrecision int
}

type Pair struct {
	C1 *Coin
	C2 *Coin
}
type AddPairsRequest struct {
	Pairs []*Pair
}
