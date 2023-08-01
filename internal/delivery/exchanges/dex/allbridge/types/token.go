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

func TokenFromEntity(t *entity.Token) *Token {
	return &Token{
		ContractAddress: t.ContractAddress,
		Decimals:        t.Decimals,
		Native:          t.Native,
	}
}

func (t *Token) Check() error {
	if t.ContractAddress == "" {
		return fmt.Errorf("contractAddress cannot be empty")
	}

	if t.Decimals == 0 {
		return fmt.Errorf("decimals cannot be 0")
	}
	return nil
}

type EToken struct {
	TransferTime map[string]TransferTime
}

func (t *EToken) Snapshot() entity.ExchangeToken {
	tt := make(map[string]TransferTime)
	if t.TransferTime != nil {
		for n, t := range t.TransferTime {
			tt[n] = TransferTime{
				Allbridge: t.Allbridge,
				Wormhole:  t.Wormhole,
			}
		}
	}
	return &EToken{TransferTime: tt}
}
