package types

import (
	"exchange-provider/internal/entity"
	"fmt"
)

type Token struct {
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	Decimals        uint64 `json:"decimals"`
	Native          bool   `json:"native"`
}

func (t *Token) Check() error {
	if t.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
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
			Name:            t.StableToken.Name,
			ContractAddress: t.StableToken.ContractAddress,
			Decimals:        t.StableToken.Decimals,
			Native:          t.StableToken.Native,
		},
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
