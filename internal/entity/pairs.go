package entity

type Chain struct {
	Id string
}

type Coin struct {
	CoinId  string
	ChainId string
}

type PairCoin struct {
	*Coin

	MinOrderSize        string
	MaxOrderSize        string
	MinWithdrawalSize   string
	WithdrawalMinFee    string
	OrderPrecision      int
	WithdrawalPrecision int
	SetChain            bool
}

type Pair struct {
	BC *PairCoin
	QC *PairCoin

	Price        string
	FeeCurrency  string
	OrderFeeRate string
}
