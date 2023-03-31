package dto

import (
	"exchange-provider/internal/entity"
)

type Token struct {
	Symbol   string `json:"symbol"`
	Standard string `json:"standard"`
	Network  string `json:"network"`

	Address  string  `json:"address,omitempty"`
	Decimals uint64  `json:"decimals,omitempty"`
	Native   bool    `json:"native,omitempty"`
	Min      float64 `json:"min,omitempty"`
	Max      float64 `json:"max,omitempty"`
}

func tokenFromEntity(et *entity.Token, info bool) Token {
	t := Token{
		Symbol:   et.Symbol,
		Standard: et.Standard,
		Network:  et.Network,
	}
	if info {
		t.Address = et.ContractAddress
		t.Decimals = et.Decimals
		t.Native = et.Native
		t.Min = et.Min
		t.Max = et.Max

	}
	return t
}

func (t *Token) ToEntity() *entity.Token {
	return &entity.Token{
		Symbol:   t.Symbol,
		Standard: t.Standard,
		Network:  t.Network,

		ContractAddress: t.Address,
		Decimals:        t.Decimals,
		Native:          t.Native,
		Min:             t.Min,
		Max:             t.Max,
	}
}
