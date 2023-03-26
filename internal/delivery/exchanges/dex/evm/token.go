package evm

import (
	"exchange-provider/internal/entity"
)

type Token struct {
}

func (t *Token) Snapshot() entity.ExchangeToken { return &Token{} }
