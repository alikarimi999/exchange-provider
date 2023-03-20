package entity

import (
	"math/big"
)

type ExchangePair interface {
	Snapshot() ExchangePair
}

type Pair struct {
	T1 *Token
	T2 *Token

	LP              uint
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

	EP ExchangePair
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}

func (p *Pair) Equal(p1 *Pair) bool {
	return (p.T1.Equal(p1.T1) && p.T2.Equal(p1.T2))
}

func (p *Pair) Snapshot() *Pair {
	var ep ExchangePair
	if p.EP != nil {
		ep = p.EP.Snapshot()
	} else {
		ep = nil
	}
	return &Pair{
		T1: p.T1.Snapshot(),
		T2: p.T2.Snapshot(),
		LP: p.LP,

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
		EP:              ep,
	}
}
