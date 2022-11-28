package multichain

import (
	"exchange-provider/internal/entity"
)

type token struct {
	Address  string
	Chain    string
	Name     string
	Symbol   string
	Decimals int
	Native   bool
	cs       []*pairInfo
}

func (t *token) toCoin() *entity.Coin {
	return &entity.Coin{
		CoinId:  t.Symbol,
		ChainId: t.Chain,
	}
}

func c2T(c *entity.Coin) *token {
	return &token{
		Symbol: c.CoinId,
		Chain:  c.ChainId,
	}
}
