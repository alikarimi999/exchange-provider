package swapspace

import "exchange-provider/internal/entity"

type pair struct {
	t1 *token
	t2 *token
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		T1: &entity.PairToken{
			Token: p.t1.toEntity(),
		},
		T2: &entity.PairToken{
			Token: p.t2.toEntity(),
		},
	}
}

func (ex *exchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {
	return nil, nil
}

func (ex *exchange) RemovePair(t1, t2 *entity.Token) error {
	return nil
}
