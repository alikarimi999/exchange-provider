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

	_, exist := s.pairs[pairId(p.BT.Symbol, p.QT.Symbol)]
	if !exist {
		_, exist = s.pairs[pairId(p.QT.Symbol, p.BT.Symbol)]
	}
	if !exist {
		s.pairs[pairId(p.BT.Symbol, p.QT.Symbol)] = p
	}
}

func (s *supportedPairs) get(bt, qt string) (*types.Pair, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if p, exist := s.pairs[pairId(bt, qt)]; exist {
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

func (s *supportedPairs) exist(bt, qt string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, exist := s.pairs[pairId(bt, qt)]
	if !exist {
		_, exist = s.pairs[pairId(qt, bt)]
	}
	return exist
}

func (s *supportedPairs) existsExactly(bt, qt string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, exist := s.pairs[pairId(bt, qt)]
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

func pairId(bt, qt string) string {
	return fmt.Sprintf("%s%s%s", bt, types.Delimiter, qt)
}
