package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"sync"
)

const (
	pairDelimiter = "-"
)

type exCoin struct {
	*entity.Coin
	minSize   string
	maxSize   string
	precision int
}

type pair struct {
	id     string  // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	symbol string  // base.id + pairDelimiter + quote.id
	bc     *exCoin // base coin
	qc     *exCoin // quote coin
}

func fromEntity(ep *entity.Pair) *pair {
	return &pair{
		id:     ep.BC.Id + ep.BC.Chain.Id + ep.QC.Id + ep.QC.Chain.Id,
		symbol: ep.BC.Id + pairDelimiter + ep.QC.Id,
		bc: &exCoin{
			Coin:      ep.BC.Coin,
			minSize:   ep.BC.MinOrderSize,
			maxSize:   ep.BC.MaxOrderSize,
			precision: ep.BC.Precision,
		},
		qc: &exCoin{
			Coin:      ep.QC.Coin,
			minSize:   ep.QC.MinOrderSize,
			maxSize:   ep.QC.MaxOrderSize,
			precision: ep.QC.Precision,
		},
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		BC: &entity.PairCoin{
			Coin:         p.bc.Coin,
			MinOrderSize: p.bc.minSize,
			MaxOrderSize: p.bc.maxSize,
			Precision:    p.bc.precision,
		},
		QC: &entity.PairCoin{
			Coin:         p.qc.Coin,
			MinOrderSize: p.qc.minSize,
			MaxOrderSize: p.qc.maxSize,
			Precision:    p.qc.precision,
		},
	}
}

type exPairs struct {
	mux   *sync.Mutex
	pairs map[string]*pair // map[id]*pair
}

func newExPairs() *exPairs {
	return &exPairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]*pair),
	}
}

func (sp *exPairs) add(pairs []*pair) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	for _, p := range pairs {
		sp.pairs[p.id] = p
	}
}

func (sp *exPairs) exists(c1, c2 *entity.Coin) bool {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	id := c1.Id + c1.Chain.Id + c2.Id + c2.Chain.Id
	_, exist := sp.pairs[id]
	if exist {
		return true
	}

	id = c2.Id + c2.Chain.Id + c1.Id + c1.Chain.Id
	_, exist = sp.pairs[id]
	return exist
}

func (sp *exPairs) get(c1, c2 *entity.Coin) (*pair, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	id := c1.Id + c1.Chain.Id + c2.Id + c2.Chain.Id
	p, exist := sp.pairs[id]
	if exist {
		return p, nil
	}

	id = c2.Id + c2.Chain.Id + c1.Id + c1.Chain.Id
	p, exist = sp.pairs[id]
	if exist {
		return p, nil
	}

	return nil, errors.New("pair not found")
}

func (sp *exPairs) snapshot() []*pair {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	pairs := make([]*pair, 0)
	for _, p := range sp.pairs {
		pairs = append(pairs, p)
	}
	return pairs
}
