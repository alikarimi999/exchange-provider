package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"sync"
)

const (
	pairDelimiter = "-"
)

type kuCoin struct {
	*entity.Coin
	minSize        string
	maxSize        string
	orderPrecision int
}

type pair struct {
	id     string  // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	symbol string  // base.id + pairDelimiter + quote.id
	bc     *kuCoin // base coin
	qc     *kuCoin // quote coin
}

func fromEntity(ep *entity.Pair) *pair {
	return &pair{
		id:     ep.BC.CoinId + ep.BC.ChainId + ep.QC.CoinId + ep.QC.ChainId,
		symbol: ep.BC.CoinId + pairDelimiter + ep.QC.CoinId,
		bc: &kuCoin{
			Coin:           ep.BC.Coin,
			minSize:        ep.BC.MinOrderSize,
			maxSize:        ep.BC.MaxOrderSize,
			orderPrecision: ep.BC.OrderPrecision,
		},
		qc: &kuCoin{
			Coin:           ep.QC.Coin,
			minSize:        ep.QC.MinOrderSize,
			maxSize:        ep.QC.MaxOrderSize,
			orderPrecision: ep.QC.OrderPrecision,
		},
	}
}

func (p *pair) toEntity() *entity.Pair {
	return &entity.Pair{
		BC: &entity.PairCoin{
			Coin:           p.bc.Coin,
			MinOrderSize:   p.bc.minSize,
			MaxOrderSize:   p.bc.maxSize,
			OrderPrecision: p.bc.orderPrecision,
		},
		QC: &entity.PairCoin{
			Coin:           p.qc.Coin,
			MinOrderSize:   p.qc.minSize,
			MaxOrderSize:   p.qc.maxSize,
			OrderPrecision: p.qc.orderPrecision,
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

func (sp *exPairs) exists(bc, qc *entity.Coin) bool {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	_, exist := sp.pairs[bc.CoinId+bc.ChainId+qc.CoinId+qc.ChainId]
	if exist {
		return true
	}

	return exist
}

func (sp *exPairs) get(bc, qc *entity.Coin) (*pair, error) {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	p, exist := sp.pairs[bc.CoinId+bc.ChainId+qc.CoinId+qc.ChainId]
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
