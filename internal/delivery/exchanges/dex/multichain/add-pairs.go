package multichain

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
)

type AddPairsRequest struct {
	Pairs []*Pair
}

func (a *AddPairsRequest) Validate() error {
	for _, p := range a.Pairs {
		if p.T1.CoinId != p.T2.CoinId {
			return fmt.Errorf("both tokens must have the same coinId")
		}
		if p.T1.ChainId == p.T2.ChainId {
			return fmt.Errorf("both tokens cannot have the same chainId")
		}
		if _, err := strconv.Atoi(p.T1.ChainId); err != nil {
			return err
		}
		if _, err := strconv.Atoi(p.T2.ChainId); err != nil {
			return err
		}
	}
	return nil
}

func (m *Multichain) AddPairs(data interface{}) (*entity.AddPairsResult, error) {

	req, ok := data.(*AddPairsRequest)
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest)
	}

	res := &entity.AddPairsResult{}

	for _, p := range req.Pairs {
		if exist := m.pairs.exist(p.T1, p.T2); exist {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		if _, ok := m.cs[ChainId(p.T1.ChainId)]; !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("chainId '%s' not supported", p.T1.ChainId)})
			continue
		}

		if _, ok := m.cs[ChainId(p.T2.ChainId)]; !ok {
			res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
				Err: fmt.Errorf("chainId '%s' not supported", p.T2.ChainId)})
			continue
		}

		m.pairs.add(p)
		res.Added = append(res.Added, *p.toEntity())

	}
	return res, nil
}
