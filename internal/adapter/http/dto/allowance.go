package dto

import (
	"exchange-provider/internal/entity"
)

type AllowanceReq struct {
	Token entity.TokenId `json:"token"`
	Owner string         `json:"owner"`
}

type AToken struct {
	entity.TokenId
	Address  string `json:"address"`
	Decimals uint64 `json:"decimals"`
}

type AllowanceRes struct {
	Token   AToken `json:"token"`
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
	Amount  string `json:"amount"`
}
