package types

import (
	"exchange-provider/internal/entity"
	"fmt"
)

type Token struct {
	ContractAddress string `json:"contractAddress"`
	Decimals        uint64 `json:"decimals"`
	Native          bool   `json:"native"`
}

func (t *Token) Check() error {
	if t.ContractAddress == "" {
		return fmt.Errorf("contractAddress cannot be empty")
	}
	if t.Decimals == 0 {
		return fmt.Errorf("decimals cannot be zero")
	}
	return nil
}

type EToken struct {
	StableToken Token `json:"stableToken"`
}

func (t *EToken) Snapshot() entity.ExchangeToken {
	return &EToken{
		StableToken: Token{
			ContractAddress: t.StableToken.ContractAddress,
			Decimals:        t.StableToken.Decimals,
			Native:          t.StableToken.Native,
		},
	}
}

func TokenFromEntity(t *entity.Token) *Token {
	return &Token{
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
	}
}
