package dto

import "time"

type Token struct {
	TokenId string
	ChainId string

	BlockTime           time.Duration
	WithdrawalPrecision int
}

func (c *Token) String() string {
	return c.TokenId + "-" + c.ChainId
}

type Pair struct {
	T1 *Token
	T2 *Token
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}

type AddPairsRequest struct {
	Pairs []*Pair
}
