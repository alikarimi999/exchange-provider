package multichain

import (
	"exchange-provider/internal/entity"
)

type Token struct {
	TokenId  string `json:"tokenId"`
	ChainId  string `json:"chainId"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Native   bool   `json:"native"`
	Data     *Data  `json:"data"`
}

type Data struct {
	TokenId  string `json:"tokenId"`
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
	return t.TokenId + "-" + t.ChainId
}

func (t *Token) toCoin() *entity.Token {
	return &entity.Token{
		TokenId: t.TokenId,
		ChainId: t.ChainId,
		Address: t.Address,
	}
}

func c2T(c *entity.Token) *Token {
	return &Token{
		TokenId: c.TokenId,
		ChainId: c.ChainId,
	}
}
