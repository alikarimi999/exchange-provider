package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"fmt"
)

func (*Multichain) Type() entity.ExType {
	return entity.EvmDEX
}

func (*Multichain) Stop() {}

func (m *Multichain) Run() {
	m.l.Debug(fmt.Sprintf("%s.Run", m.Id()), "started")
}

func (m *Multichain) Configs() interface{} {
	return m.cfg
}

func (m *Multichain) Support(in, out *entity.Token) bool {
	p, err := m.pairs.get(c2T(in), c2T(out))
	if err != nil {
		return false
	}

	if p.T1.ChainId == in.Standard {
		return p.T1.Data.RouterABI != ""
	} else {
		return p.T2.Data.RouterABI != ""
	}
}

func (m *Multichain) Price(ps ...*entity.Pair) ([]*entity.Pair, error) {
	// p, err := m.pairs.get(c2T(c1), c2T(c2))
	// if err != nil {
	// 	return nil, err
	// }
	// return p.toEntity(), nil
	return nil, nil
}

func (m *Multichain) GetAllPairs() []*entity.Pair {
	ps := m.pairs.getAll()

	eps := make([]*entity.Pair, 0, len(ps))
	for _, p := range ps {
		eps = append(eps, p.toEntity())
	}
	return eps
}

func (m *Multichain) RemovePair(c1, c2 *entity.Token) error {
	return m.pairs.remove(c2T(c1), c2T(c2))
}

func (p *Pair) toEntity() *entity.Pair {
	return &entity.Pair{
		T1: p.T1.toCoin(),
		T2: p.T2.toCoin(),
	}
}

func (m *Multichain) Pair(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt}, nil
}

func (m *Multichain) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt, Price1: "1"}, nil

}
