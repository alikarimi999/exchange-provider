package entity

import (
	"math/big"
	"time"
)

type Token struct {
	TokenId  string
	ChainId  string
	Address  string
	Decimals uint64
	Native   bool
}

func (c *Token) String() string {
	return c.TokenId + "-" + c.ChainId
}

func (t *Token) Equal(t2 *Token) bool {
	return t.TokenId == t2.TokenId && t.ChainId == t2.ChainId
}

func (t *Token) Snapshot() *Token {
	return &Token{TokenId: t.TokenId, ChainId: t.ChainId, Address: t.Address,
		Decimals: t.Decimals, Native: t.Native}
}

type PairToken struct {
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

func (t *PairToken) Snapshot() *PairToken {
	return &PairToken{
		Token:             t.Token.Snapshot(),
		BlockTime:         t.BlockTime,
		ContractAddress:   t.ContractAddress,
		Address:           t.Address,
		Tag:               t.Tag,
		MinDeposit:        t.MinDeposit,
		MinOrderSize:      t.MinOrderSize,
		MaxOrderSize:      t.MaxOrderSize,
		MinWithdrawalSize: t.MinWithdrawalSize,
		WithdrawalMinFee:  t.WithdrawalMinFee,
		OrderPrecision:    t.OrderPrecision,
		SetChain:          t.SetChain,
	}
}

type Pair struct {
	T1 *PairToken
	T2 *PairToken

	Exchange        string
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

func (p *PairToken) String() string {
	return p.Token.String()
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}
func (p *Pair) Snapshot() *Pair {
	return &Pair{
		T1:              p.T1.Snapshot(),
		T2:              p.T2.Snapshot(),
		Exchange:        p.Exchange,
		ContractAddress: p.ContractAddress,
		FeeTier:         p.FeeTier,
		Liquidity:       p.Liquidity,
		MinDeposit:      p.MinDeposit,
		Price1:          p.Price1,
		Price2:          p.Price2,
		FeeCurrency:     p.FeeCurrency,
		OrderFeeRate:    p.OrderFeeRate,
		SpreadRate:      p.SpreadRate,
		FeeRate:         p.FeeRate,
	}
}
