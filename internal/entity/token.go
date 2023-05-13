package entity

import "strings"

type ExchangeToken interface {
	Snapshot() ExchangeToken
}

type TokenId struct {
	Symbol   string `json:"symbol"`
	Standard string `json:"standard"`
	Network  string `json:"network"`
}

func (t *TokenId) ToUpper() *TokenId {
	t.Symbol = strings.ToUpper(t.Symbol)
	t.Standard = strings.ToUpper(t.Standard)
	t.Network = strings.ToUpper(t.Network)
	return t
}

func (id *TokenId) String() string {
	return strings.ToUpper(id.Symbol + "-" + id.Standard + "-" + id.Network)
}

type Token struct {
	Id              TokenId `json:"id"`
	ContractAddress string  `json:"contractAddress"`
	Decimals        uint64  `json:"decimals"`
	Native          bool    `json:"native"`
	Min             float64 `json:"min"`
	Max             float64 `json:"max"`

	ET ExchangeToken `json:"et"`
}

func (t *Token) String() string {
	return t.Id.String()
}

func (t *Token) Snapshot() *Token {
	var et ExchangeToken
	if t.ET != nil {
		et = t.ET.Snapshot()
	} else {
		et = nil
	}

	return &Token{
		Id: TokenId{
			Symbol:   t.Id.Symbol,
			Standard: t.Id.Standard,
			Network:  t.Id.Network,
		},

		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,

		ET: et,
	}
}
