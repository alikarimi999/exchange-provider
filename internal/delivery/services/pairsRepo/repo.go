package pairsRepo

import (
	"exchange-provider/internal/entity"
	"sort"
	"sync"
)

type pairsRepo struct {
	mux *sync.RWMutex
	eps map[uint]*exPairs
}

func PairsRepo() entity.PairsRepo {
	return &pairsRepo{
		mux: &sync.RWMutex{},
		eps: make(map[uint]*exPairs),
	}
}

func (pr *pairsRepo) Add(exId uint, ps ...*entity.Pair) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	ep, ok := pr.eps[exId]
	if !ok {
		ep = newExPairs()
		pr.eps[exId] = ep
		pr.sortEps()
	}

	ep.add(ps...)
}

func (pr *pairsRepo) Get(exId uint, t1, t2 string) (*entity.Pair, bool) {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	ep, ok := pr.eps[exId]
	if !ok {
		return nil, false
	}
	return ep.get(t1, t2)
}

func (pr *pairsRepo) Exists(exId uint, t1, t2 string) bool {
	pr.mux.RLock()
	defer pr.mux.RUnlock()
	ep, ok := pr.eps[exId]
	if !ok {
		return false
	}
	return ep.exists(t1, t2)
}

func (pr *pairsRepo) Update(exId uint, p *entity.Pair) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	pr.eps[exId].update(p)
}

func (pr *pairsRepo) Remove(exId uint, t1, t2 string) {
	pr.mux.Lock()
	defer pr.mux.Unlock()
	ep, ok := pr.eps[exId]
	if ok {
		ep.remove(t1, t2)
	}
}

func (pr *pairsRepo) sortEps() {
	m1 := pr.eps
	m2 := make(map[uint]*exPairs)

	keys := []int{}
	for k := range m1 {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		m2[uint(k)] = m1[uint(k)]
	}

	pr.eps = m2
}
