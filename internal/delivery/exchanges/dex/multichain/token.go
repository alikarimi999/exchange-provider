package multichain

import (
	"exchange-provider/internal/entity"
)

type Token struct {
	CoinId   string `json:"coinId"`
	ChainId  string `json:"chainId"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Native   bool   `json:"native"`
	Data     *Data  `json:"data"`
}

type Data struct {
	CoinId   string `json:"coinId"`
	ChainId  string `json:"chainId"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Native   bool   `json:"native"`

	AnyToken       *Token `json:"anytoken"`
	FromAnyToken   *Token `json:"fromanytoken"`
	Router         string `json:"router"`
	RouterABI      string `json:"routerABI"`
	RouterName     string `json:"routerName"`
	DepositAddress string `json:"DepositAddress"`
	IsApprove      bool   `json:"isApprove"`
}

func (t *Token) String() string {
	return t.CoinId + "-" + t.ChainId
}

func (t *Token) toCoin() *entity.Coin {
	return &entity.Coin{
		CoinId:  t.CoinId,
		ChainId: t.ChainId,
	}
}

func c2T(c *entity.Coin) *Token {
	return &Token{
		CoinId:  c.CoinId,
		ChainId: c.ChainId,
	}
}
