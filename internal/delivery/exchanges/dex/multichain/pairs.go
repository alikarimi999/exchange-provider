package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"sync"
)

type supportedPairs struct {
	mux   *sync.Mutex
	pairs map[string]*Pair
}

func newSupportedPairs() *supportedPairs {
	return &supportedPairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]*Pair),
	}
}

func (s *supportedPairs) add(p *Pair) {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, exist := s.pairs[id(p.t1, p.t2)]
	if !exist {
		_, exist = s.pairs[id(p.t2, p.t1)]
	}
	if !exist {
		s.pairs[id(p.t1, p.t2)] = p
	}
}

func (s *supportedPairs) get(t1, t2 *token) (*Pair, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if p, exist := s.pairs[id(t1, t2)]; exist {
		return p, nil
	} else if p, exist := s.pairs[id(t2, t1)]; exist {
		return p, nil
	}
	return nil, errors.Wrap(errors.ErrNotFound)
}

func (s *supportedPairs) getAll() []*Pair {
	s.mux.Lock()
	defer s.mux.Unlock()

	pairs := []*Pair{}
	for _, pair := range s.pairs {
		pairs = append(pairs, pair)
	}

	return pairs
}

func (s *supportedPairs) exist(t1, t2 *token) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, exist := s.pairs[id(t1, t2)]
	if !exist {
		_, exist = s.pairs[id(t2, t1)]
	}
	return exist
}

func (s *supportedPairs) remove(t1, t2 *token) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, exist := s.pairs[id(t1, t2)]; exist {
		delete(s.pairs, id(t1, t2))
		return nil
	} else if _, exist := s.pairs[id(t2, t1)]; exist {
		delete(s.pairs, id(t1, t2))
		return nil
	}

	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func id(t1, t2 *token) string {
	return t1.Symbol + "-" + t1.Chain + types.Delimiter + t2.Symbol + "-" + t2.Chain
}
