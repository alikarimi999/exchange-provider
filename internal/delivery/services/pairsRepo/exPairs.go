package pairsRepo

import (
	"exchange-provider/internal/entity"
	"sort"
	"sync"
)

type exPairs struct {
	mux   *sync.RWMutex
	exId  uint
	exNID string
	pairs map[string]*entity.Pair
}

func newExPairs(ex entity.Exchange) *exPairs {
	return &exPairs{
		mux:   &sync.RWMutex{},
		exId:  ex.Id(),
		exNID: ex.NID(),
		pairs: make(map[string]*entity.Pair),
	}
}

func (ep *exPairs) get(t1, t2 string) (*entity.Pair, bool) {
	ep.mux.RLock()
	defer ep.mux.RUnlock()
	p, ok := ep.pairs[pairId(t1, t2)]
	if ok {
		return p.Snapshot(), true
	}
	return nil, false
}

func (ep *exPairs) getAll() []*entity.Pair {
	ep.mux.RLock()
	defer ep.mux.RUnlock()
	ps := []*entity.Pair{}
	for _, p := range ep.pairs {
		ps = append(ps, p.Snapshot())
	}
	return ps
}

func (ep *exPairs) exists(t1, t2 string) bool {
	ep.mux.RLock()
	defer ep.mux.RUnlock()
	_, ok := ep.pairs[pairId(t1, t2)]
	return ok
}

func (ep *exPairs) add(ps ...*entity.Pair) {
	ep.mux.Lock()
	defer ep.mux.Unlock()
	for _, p := range ps {
		ep.pairs[pairId(p.T1.String(), p.T2.String())] = p.Snapshot()
	}
	ep.sortPairs()
}

func (ep *exPairs) update(p *entity.Pair) {
	ep.mux.Lock()
	defer ep.mux.Unlock()
	for id, p0 := range ep.pairs {
		if p0.Equal(p) {
			ep.pairs[id] = p.Snapshot()
		}
	}
}

func (ep *exPairs) remove(t1, t2 string) {
	ep.mux.Lock()
	defer ep.mux.Unlock()
	delete(ep.pairs, pairId(t1, t2))
}

func (ep *exPairs) sortPairs() {
	m1 := ep.pairs
	m2 := make(map[string]*entity.Pair)

	keys := []string{}
	for k := range m1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		m2[k] = m1[k]
	}

	ep.pairs = m2
}
