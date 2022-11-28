package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
)

type Pair struct {
	t1 *token
	t2 *token
}

func (m *Multichain) GetAddress(c *entity.Coin) (*entity.Address, error) {
	a, err := m.cs[chainId(c.ChainId)].w.RandAddress()
	if err != nil {
		return nil, err
	}
	return &entity.Address{Addr: a.String()}, nil
}

func (m *Multichain) Support(c1, c2 *entity.Coin) bool {
	return m.pairs.exist(c2T(c1), c2T(c2))
}

func (m *Multichain) GetPair(c1, c2 *entity.Coin) (*entity.Pair, error) {
	p, err := m.pairs.get(c2T(c1), c2T(c2))
	if err != nil {
		return nil, err
	}
	return p.toEntity(), nil
}

func (m *Multichain) GetAllPairs() []*entity.Pair {
	ps := m.pairs.getAll()

	eps := make([]*entity.Pair, 0, len(ps))
	for _, p := range ps {
		eps = append(eps, p.toEntity())
	}
	return eps
}

func (m *Multichain) RemovePair(c1, c2 *entity.Coin) error {
	return m.pairs.remove(c2T(c1), c2T(c2))
}

func (p *Pair) toEntity() *entity.Pair {
	return &entity.Pair{
		C1: &entity.PairCoin{
			Coin:            p.t1.toCoin(),
			ContractAddress: p.t1.Address,
		},
		C2: &entity.PairCoin{
			Coin:            p.t2.toCoin(),
			ContractAddress: p.t2.Address,
		},
	}
}

func (m *Multichain) Pair(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt}, nil
}

func (m *Multichain) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt, Price: "1"}, nil

}
