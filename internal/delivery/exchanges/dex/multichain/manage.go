package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"sync"
)

func (*Multichain) Type() entity.ExType {
	return entity.DEX
}

func (*Multichain) Stop() {}

func (*Multichain) StartAgain() (*entity.StartAgainResult, error) {
	return nil, nil
}

func (*Multichain) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}

func (*Multichain) Run(wg *sync.WaitGroup) {}

func (m *Multichain) Configs() interface{} {
	return m.cfg
}

func (m *Multichain) Support(in, out *entity.Coin) bool {
	p, err := m.pairs.get(c2T(in), c2T(out))
	if err != nil {
		return false
	}

	if p.T1.ChainId == in.ChainId {
		return p.T1.Data.RouterABI != ""
	} else {
		return p.T2.Data.RouterABI != ""
	}
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
			Coin:            p.T1.toCoin(),
			ContractAddress: p.T1.Address,
		},
		C2: &entity.PairCoin{
			Coin:            p.T2.toCoin(),
			ContractAddress: p.T2.Address,
		},
	}
}

func (m *Multichain) Pair(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt}, nil
}

func (m *Multichain) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt, Price: "1"}, nil

}
