package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
)

type EToken struct {
	Symbol string
	types.Token
	Min float64      `json:"min"`
	Max float64      `json:"max"`
	Et  types.EToken `json:"exchangeToken"`
}
