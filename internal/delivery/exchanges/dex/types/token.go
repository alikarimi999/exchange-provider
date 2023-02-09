package types

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Address  common.Address `json:"address"`
	Decimals int            `json:"decimals"`
	ChainId  int64          `json:"chainId"`
	Native   bool           `json:"native"`
}

func (t *Token) IsNative() bool {
	return t.Native
}

func (t *Token) String() string {
	return t.Symbol
}

func (t *Token) ToEntity(standard string) *entity.PairToken {
	return &entity.PairToken{
		Token: &entity.Token{
			TokenId:  t.Symbol,
			ChainId:  standard,
			Address:  t.Address.String(),
			Decimals: uint64(t.Decimals),
		},
	}
}
