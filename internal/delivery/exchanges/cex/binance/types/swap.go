package types

import (
	"exchange-provider/internal/entity"

	"github.com/adshao/go-binance/v2"
)

type BinanceFee struct {
	Coin   string
	Amount float64
}

type Swap struct {
	In                entity.TokenId
	Out               entity.TokenId
	Side              binance.SideType
	InAmountRequested float64
	InAmountExecuted  float64
	OutAmount         float64
	BinanceFees       []BinanceFee
	FeeCurrency       string
}
