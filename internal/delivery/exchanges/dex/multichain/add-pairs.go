package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (m *Multichain) AddPairs(data interface{}) (*entity.AddPairsResult, error) {

	req, ok := data.(*dto.AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}
	// ps := []*Pair{}
	for _, p := range req.Pairs {
		c1, err := parseCoin(p.C1)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}
		c2, err := parseCoin(p.C2)
		if err != nil {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(), Err: err})
			continue
		}

		if exist := m.pairs.exist(c2T(c1), c2T(c2)); exist {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		t1, t2 := getinfos(m.f, c1.CoinId, c1.ChainId, c2.CoinId, c2.ChainId)
		if len(t1.cs) == 0 || len(t2.cs) == 0 {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("not found")})
			continue
		}

		if _, ok := m.cs[chainId(t1.Chain)]; !ok {
			c, err := m.newChain(t1.Chain)
			if err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
					Err: err})
				continue
			}
			m.cs[chainId(t1.Chain)] = c
		}

		if _, ok := m.cs[chainId(t2.Chain)]; !ok {
			c, err := m.newChain(t2.Chain)
			if err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
					Err: err})
				continue
			}
			m.cs[chainId(t1.Chain)] = c
		}

		m.pairs.add(&Pair{t1: t1, t2: t2})

	}
	return res, nil
}
