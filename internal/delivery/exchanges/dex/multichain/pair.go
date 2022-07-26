package multichain

import "exchange-provider/internal/entity"

type Pair struct {
	T1 *Token `json:"t1"`
	T2 *Token `json:"t2"`
}

func (p *Pair) toEntiy() *entity.Pair {
	return &entity.Pair{
		T1: &entity.PairCoin{Token: p.T1.toCoin()},
		T2: &entity.PairCoin{Token: p.T2.toCoin()},
	}
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}
