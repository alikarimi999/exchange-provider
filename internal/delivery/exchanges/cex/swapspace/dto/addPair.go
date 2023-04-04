package dto

import (
	"exchange-provider/internal/entity"
)

type AddPairsRequest struct {
	Pairs []*Pair `json:"pairs"`
}

type Pair struct {
	T1      *EToken `json:"t1"`
	T2      *EToken `json:"t2"`
	FeeRate float64 `json:"feeRate"`
}

type EToken struct {
	Symbol   string `json:"symbol"`
	Standard string `json:"standard"`
	Network  string `json:"network"`

	Address  string `json:"address"`
	Decimals uint64 `json:"decimals"`
	Native   bool   `json:"native"`
	ET       Token  `json:"exchangeToken"`
}

func (t *EToken) toEntity(fn func(Token) entity.ExchangeToken) *entity.Token {
	return &entity.Token{
		Symbol:   t.Symbol,
		Standard: t.Standard,
		Network:  t.Network,

		ContractAddress: t.Address,
		Decimals:        t.Decimals,
		Native:          t.Native,
		ET:              fn(t.ET),
	}
}

type Token struct {
	Code       string `json:"code"`
	Network    string `json:"network"`
	HasExtraId bool   `json:"hasExtraId"`
}

func (p *Pair) ToEntity(fn func(Token) entity.ExchangeToken) *entity.Pair {
	return &entity.Pair{
		T1:      p.T1.toEntity(fn),
		T2:      p.T2.toEntity(fn),
		FeeRate: p.FeeRate,
	}
}
