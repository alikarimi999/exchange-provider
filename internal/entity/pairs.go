package entity

type Chain struct {
	Id string
}

type Coin struct {
	CoinId  string
	ChainId string
}

func (c *Coin) String() string {
	return c.CoinId + "-" + c.ChainId
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

	BestAsk      string
	BestBid      string
	FeeCurrency  string
	OrderFeeRate string
	Fee          string
}

func (p *PairCoin) String() string {
	return p.CoinId + "-" + p.ChainId
}

func (p *Pair) String() string {
	return p.BC.String() + "/" + p.QC.String()
}
