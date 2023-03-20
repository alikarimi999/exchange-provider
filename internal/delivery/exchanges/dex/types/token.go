package types

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Network  string         `json:"network"`
	Address  common.Address `json:"address"`
	Decimals int            `json:"decimals"`
	ChainId  int64          `json:"chainId"`
	Native   bool           `json:"native"`
}

func (t *Token) SnapShot() *Token {
	return &Token{
		Name:     t.Name,
		Symbol:   t.Symbol,
		Address:  t.Address,
		Decimals: t.Decimals,
		ChainId:  t.ChainId,
		Native:   t.Native,
	}
}

func (t *Token) IsNative() bool {
	return t.Native
}

func (t *Token) String() string {
	return t.Symbol
}

func (t *Token) ToToken() *entity.Token {
	return &entity.Token{
		Symbol:   t.Symbol,
		Standard: t.Network,
		Address:  t.Address.String(),
		Decimals: uint64(t.Decimals),
	}
}
func (t *Token) ToEntity(standard string) *entity.Token {
	return &entity.Token{
		Symbol:   t.Symbol,
		Standard: standard,
		Address:  t.Address.String(),
		Decimals: uint64(t.Decimals),
	}
}
