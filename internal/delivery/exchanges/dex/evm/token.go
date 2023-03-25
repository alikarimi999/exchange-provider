package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/dto"
	"exchange-provider/internal/entity"
)

type Token struct {
}

func (t *Token) Snapshot() entity.ExchangeToken { return &Token{} }

func convert(t *dto.Token) *Token {
	return &Token{}
}
