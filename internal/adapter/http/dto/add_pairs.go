package dto

import (
	"exchange-provider/internal/entity"
)

type PairsErr struct {
	Pair string `json:"pair"`
	Err  string `json:"error"`
}
type AddPairsResult struct {
	Addedd []string    `json:"addedPairs"`
	Exs    []string    `json:"existedPairs"`
	Failed []*PairsErr `json:"failedPairs"`
}

func FromEntity(r *entity.AddPairsResult) *AddPairsResult {
	res := &AddPairsResult{
		Addedd: []string{},
		Exs:    []string{},
		Failed: []*PairsErr{},
	}
	res.Addedd = append(res.Addedd, r.Added...)
	res.Exs = append(res.Exs, r.Existed...)

	for _, p := range r.Failed {
		res.Failed = append(res.Failed, &PairsErr{
			Pair: p.Pair,
			Err:  p.Err.Error(),
		})
	}
	return res

}
