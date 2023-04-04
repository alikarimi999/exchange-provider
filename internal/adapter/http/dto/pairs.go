package dto

import (
	"exchange-provider/internal/entity"
)

func PairFromEntity(p *entity.Pair) Pair {
	return Pair{
		T1:      tokenFromEntity(p.T1, true),
		T2:      tokenFromEntity(p.T2, true),
		FeeRate: p.FeeRate,
		LP:      p.LP,
	}
}

type Pair struct {
	T1      Token   `json:"t1"`
	T2      Token   `json:"t2"`
	FeeRate float64 `json:"feeRate"`
	LP      uint    `json:"lp"`
}
