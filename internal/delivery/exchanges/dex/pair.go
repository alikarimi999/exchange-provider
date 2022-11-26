package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"
)

type supportedPairs struct {
	mux   *sync.Mutex
	pairs map[string]types.Pair
}

func newSupportedPairs() *supportedPairs {
	return &supportedPairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]types.Pair),
	}
}

func (s *supportedPairs) add(p types.Pair) {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, exist := s.pairs[pairId(p.T1.Symbol, p.T2.Symbol)]
	if !exist {
		_, exist = s.pairs[pairId(p.T2.Symbol, p.T1.Symbol)]
	}
	if !exist {
		s.pairs[pairId(p.T1.Symbol, p.T2.Symbol)] = p
	}
}

func (s *supportedPairs) get(t1, t2 string) (*types.Pair, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if p, exist := s.pairs[pairId(t1, t2)]; exist {
		return &p, nil
	} else if p, exist := s.pairs[pairId(t2, t1)]; exist {
		return &p, nil
	}
	return &types.Pair{}, errors.Wrap(errors.ErrNotFound)
}

func (s *supportedPairs) getAll() []types.Pair {
	s.mux.Lock()
	defer s.mux.Unlock()

	pairs := []types.Pair{}
	for _, pair := range s.pairs {
		pairs = append(pairs, pair)
	}

	return pairs
}

func (s *supportedPairs) exist(c1, c2 string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, exist := s.pairs[pairId(c1, c2)]
	if !exist {
		_, exist = s.pairs[pairId(c2, c1)]
	}
	return exist
}

func (s *supportedPairs) remove(bt, qt string) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, exist := s.pairs[pairId(bt, qt)]; exist {
		delete(s.pairs, pairId(bt, qt))
		return nil
	}

	return errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair not found"))
}

func pairId(c1, c2 string) string {
	return fmt.Sprintf("%s%s%s", c1, types.Delimiter, c2)
}
