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

type PairCoin struct {
	*Coin

	MinOrderSize      string
	MaxOrderSize      string
	MinWithdrawalSize string
	WithdrawalMinFee  string
	Precision         int
	SetChain          bool
}

type Pair struct {
	BC *PairCoin
	QC *PairCoin

	Price        string
	FeeCurrency  string
	OrderFeeRate string
}
