package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/dto"
	"exchange-provider/internal/entity"
)

func (d *exchange) AddPairs(data interface{}) (*entity.AddPairsResult, error) {

	req := data.(*dto.AddPairsRequest)
	ps := []*entity.Pair{}
	res := &entity.AddPairsResult{}
	for _, p := range req.Pairs {

		ep, err := p.ToEntity()
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}
		_, err = d.network(p.T1.TokenId)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}
		_, err = d.network(p.T2.TokenId)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}
		ps = append(ps, ep)
	}

	if len(ps) > 0 {
		if err := d.pairs.Add(d, ps...); err != nil {
			return nil, err
		}
	}
	return res, nil
}
func (d *exchange) RemovePair(t1, t2 entity.TokenId) error {
	return d.pairs.Remove(d.Id(), t1.String(), t2.String(), true)
}
