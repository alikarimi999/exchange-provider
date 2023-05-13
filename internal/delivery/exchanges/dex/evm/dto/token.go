package dto

import "exchange-provider/internal/entity"

type EToken struct {
	entity.TokenId

	ContractAddress string `json:"contractAddress"`
	Decimals        uint64 `json:"decimals"`
	Native          bool   `json:"native"`
	ET              Token  `json:"exchangeToken"`
}

type Token struct {
	Name            string
	ContractAddress string `json:"contractAddress"`
	Decimals        uint64 `json:"decimals"`
}
