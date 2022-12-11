package entity

import (
	"math/big"
	"time"
)

type Coin struct {
	CoinId  string
	ChainId string
}

func (c *Coin) String() string {
	return c.CoinId + "-" + c.ChainId
}

type PairCoin struct {
	*Coin

	BlockTime           time.Duration
	ContractAddress     string
	Address             string
	Tag                 string
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
	C1 *PairCoin
	C2 *PairCoin

	ContractAddress string
	FeeTier         int64
	Liquidity       *big.Int
	MinDeposit      float64
	Price1          string
	Price2          string
	FeeCurrency     string
	OrderFeeRate    string
	SpreadRate      string
	FeeRate         string
}

func (p *PairCoin) String() string {
	return p.CoinId + "-" + p.ChainId
}

func (p *Pair) String() string {
	return p.C1.String() + "/" + p.C2.String()
}
