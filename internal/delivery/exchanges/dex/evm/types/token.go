package types

import "exchange-provider/internal/entity"

type Token struct {
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	Decimals        uint64 `json:"decimals"`
	Native          bool   `json:"native"`
}

func (t *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		Name:            t.Name,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
	}
}

func TokenFromEntity(t *entity.Token) *Token {
	return &Token{
		Name:            t.Id.Symbol,
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
	}
}
