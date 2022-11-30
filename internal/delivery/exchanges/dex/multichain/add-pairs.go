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
	// ps := []*Pair{}
	for _, p := range req.Pairs {
		if exist := m.pairs.exist(p.T1, p.T2); exist {
			res.Existed = append(res.Existed, p.String())
			continue
		}

		if _, ok := m.cs[chainId(p.T1.ChainId)]; !ok {
			c, err := m.newChain(p.T1.ChainId)
			if err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
					Err: err})
				continue
			}
			m.cs[chainId(p.T1.ChainId)] = c
		}

		if _, ok := m.cs[chainId(p.T2.ChainId)]; !ok {
			c, err := m.newChain(p.T2.ChainId)
			if err != nil {
				res.Failed = append(res.Failed, &entity.PairsErr{Pair: p.String(),
					Err: err})
				continue
			}
			m.cs[chainId(p.T2.ChainId)] = c
		}

		m.pairs.add(p)
		res.Added = append(res.Added, *p.toEntity())

	}
	return res, nil
}
