package types

import "exchange-provider/internal/entity"

type Swap struct {
	In                entity.TokenId
	Out               entity.TokenId
	Side              string
	InAmountRequested float64
	InAmountExecuted  float64
	OutAmount         float64
	KucoinFee         float64
	FeeCurrency       string
}
