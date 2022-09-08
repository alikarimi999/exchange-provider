package uniswapv3

import (
	"fmt"
	"order_service/pkg/errors"
	"sync"
)

var pairDelimiter string = "/"

type pair struct {
	BT *token
	QT *token
}

func (p *pair) String() string {
	return fmt.Sprintf("%s%s%s", p.BT.String(), pairDelimiter, p.QT.String())
}

type supportedPairs struct {
	mux   *sync.Mutex
	pairs map[string]pair
}

func newSupportedPairs() *supportedPairs {
	return &supportedPairs{
		mux:   &sync.Mutex{},
		pairs: make(map[string]pair),
	}
}

func (s *supportedPairs) get(bt, qt string) (pair, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if p, exist := s.pairs[pairId(bt, qt)]; exist {
		return p, nil
	}
	return pair{}, errors.Wrap(errors.ErrNotFound)
}

func (s *supportedPairs) remove(bt, qt string) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, exist := s.pairs[pairId(bt, qt)]; exist {
		delete(s.pairs, pairId(bt, qt))
		return nil
	}

	return errors.Wrap(errors.ErrNotFound)
}

func pairId(bt, qt string) string {
	return fmt.Sprintf("%s%s%s", bt, pairDelimiter, qt)
}
