package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"sync"
)

const (
	pairDelimiter = "-"
)

type exchangeCoin struct {
	*entity.Coin
	precision int
}

type pair struct {
	id     string        // base.id + base.Chain.Id + quote.id + quote.Chain.Id
	symbol string        // base.id + pairDelimiter + quote.id
	b      *exchangeCoin // base coin
	q      *exchangeCoin // quote coin
}

func newPair(ep *entity.ExchangePair) *pair {
	return &pair{
		id:     ep.BC.Id + ep.BC.Chain.Id + ep.QC.Id + ep.QC.Chain.Id,
		symbol: ep.BC.Id + pairDelimiter + ep.QC.Id,
		b:      &exchangeCoin{ep.BC.Coin, ep.BC.Precision},
		q:      &exchangeCoin{ep.QC.Coin, ep.QC.Precision},
	}
}

type exchangePairs struct {
	mux   *sync.Mutex
	pairs map[string]*pair // map[id]*pair
}

func newExchangePairs() *exchangePairs {
	return &exchangePairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]*pair),
	}
}

func (sp *exchangePairs) add(pairs []*pair) {
	sp.mux.Lock()
	defer sp.mux.Unlock()
	for _, p := range pairs {
		sp.pairs[p.id] = p
	}
}

func (sp *exchangePairs) exists(c1, c2 *entity.Coin) bool {
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

func (sp *exchangePairs) get(c1, c2 *entity.Coin) (*pair, error) {
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

func (sp *exchangePairs) snapshot() []*pair {
	sp.mux.Lock()
	defer sp.mux.Unlock()

	pairs := make([]*pair, 0)
	for _, p := range sp.pairs {
		pairs = append(pairs, p)
	}
	return pairs
}
