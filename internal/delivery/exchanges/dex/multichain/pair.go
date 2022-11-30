package multichain

import "exchange-provider/internal/entity"

type Pair struct {
	T1 *Token `json:"t1"`
	T2 *Token `json:"t2"`
}

func (p *Pair) toEntiy() *entity.Pair {
	return &entity.Pair{
		C1: &entity.PairCoin{Coin: p.T1.toCoin()},
		C2: &entity.PairCoin{Coin: p.T2.toCoin()},
	}
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}
