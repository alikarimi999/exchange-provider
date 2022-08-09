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

	MinDeposit          float64
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

	MinDeposit   float64
	BestAsk      string
	BestBid      string
	FeeCurrency  string
	OrderFeeRate string
	SpreadRate   string
	FeeRate      string
}

func (p *PairCoin) String() string {
	return p.CoinId + "-" + p.ChainId
}

func (p *Pair) String() string {
	return p.BC.String() + "/" + p.QC.String()
}
