package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
)

type EToken struct {
	entity.TokenId
	types.Token
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
