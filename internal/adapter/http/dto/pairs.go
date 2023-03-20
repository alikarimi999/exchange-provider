package dto

import (
	"exchange-provider/internal/entity"
)

func PairFromEntity(p *entity.Pair) Pair {
	return Pair{
		T1: tokenFromEntity(p.T1, true),
		T2: tokenFromEntity(p.T2, true),
		LP: p.LP,
	}
}

type Pair struct {
	T1 Token `json:"t1"`
	T2 Token `json:"t2"`
	LP uint  `json:"lp"`
}
