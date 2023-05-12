package kucoin

import (
	"exchange-provider/internal/entity"
)

type ExchangePair struct {
	HasIntermediaryCoin bool   `json:"hasIntermediaryCoin"`
	IC1                 *Token `json:"ic1"`
	IC2                 *Token `json:"ic2"`

	KucoinFeeRate1 float64 `json:"kucoinFeeRate1"`
	KucoinFeeRate2 float64 `json:"kucoinFeeRate2"`
}

func (t *Token) toId() entity.TokenId {
	return entity.TokenId{
		Symbol: t.Currency,
	}
}

func (p *ExchangePair) Snapshot() entity.ExchangePair {
	var ic0, ic1 *Token
	if p.IC1 != nil && p.IC2 != nil {
		ic0 = p.IC1.snapshot()
		ic1 = p.IC2.snapshot()
	}
	return &ExchangePair{
		HasIntermediaryCoin: p.HasIntermediaryCoin,
		IC1:                 ic0,
		IC2:                 ic1,
		KucoinFeeRate1:      p.KucoinFeeRate1,
		KucoinFeeRate2:      p.KucoinFeeRate2,
	}
}
