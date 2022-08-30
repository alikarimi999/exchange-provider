package dto

import "time"

type Coin struct {
	CoinId  string
	ChainId string

	BlockTime           time.Duration
	WithdrawalPrecision int
}

type Pair struct {
	BC *Coin
	QC *Coin
}
type AddPairsRequest struct {
	Pairs []*Pair
}
