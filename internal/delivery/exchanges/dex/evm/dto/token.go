package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
)

type EToken struct {
	types.Token
	Min float64       `json:"min"`
	Max float64       `json:"max"`
	ET  *types.EToken `json:"exchangeToken"`
}
