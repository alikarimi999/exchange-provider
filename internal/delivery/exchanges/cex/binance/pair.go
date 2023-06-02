package binance

import "exchange-provider/internal/entity"

type ExchangePair struct {
	HasIntermediaryCoin bool   `json:"hasIntermediaryCoin"`
	IC1                 *Token `json:"ic1"`
	IC2                 *Token `json:"ic2"`

	BinanceFeeRate1 float64 `json:"binanceFeeRate1"`
	BinanceFeeRate2 float64 `json:"binanceFeeRate2"`
}

func (t *Token) toId() entity.TokenId {
	return entity.TokenId{
		Symbol: t.Coin,
	}
}

func (p *ExchangePair) Snapshot() entity.ExchangePair {
	var ic0, ic1 *Token
	if p.IC1 != nil && p.IC2 != nil {
		ic0 = p.IC1.Snapshot().(*Token)
		ic1 = p.IC2.Snapshot().(*Token)
	}
	return &ExchangePair{
		HasIntermediaryCoin: p.HasIntermediaryCoin,
		IC1:                 ic0,
		IC2:                 ic1,
		BinanceFeeRate1:     p.BinanceFeeRate1,
		BinanceFeeRate2:     p.BinanceFeeRate2,
	}
}
