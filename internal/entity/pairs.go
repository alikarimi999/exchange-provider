package entity

import (
	"math/big"
	"time"
)

type Token struct {
	TokenId string
	ChainId string
}

func (c *Token) String() string {
	return c.TokenId + "-" + c.ChainId
}

type PairCoin struct {
	*Token

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
	T1 *PairCoin
	T2 *PairCoin

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
	return p.Token.String()
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}
