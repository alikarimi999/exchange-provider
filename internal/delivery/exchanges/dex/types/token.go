package types

import (
	"exchange-provider/internal/entity"
	"time"

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

func (t *Token) ToEntity(standard string, blockTime time.Duration) *entity.PairCoin {
	return &entity.PairCoin{
		Coin: &entity.Coin{
			CoinId:  t.Symbol,
			ChainId: standard,
		},
		BlockTime:       blockTime,
		ContractAddress: t.Address.String(),
	}
}
