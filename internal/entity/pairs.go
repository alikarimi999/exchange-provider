package entity

import "time"

type Chain struct {
	Id        string
	BlockTime time.Duration
}

type Coin struct {
	Id    string
	Chain *Chain
}

type Pair struct {
	BaseCoin  *Coin
	QuoteCoin *Coin
}
