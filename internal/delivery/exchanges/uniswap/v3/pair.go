package uniswapv3

import (
	"fmt"
	"order_service/pkg/errors"
	"sync"
)

type pair struct {
	BT *token
	QT *token
}

func (p *pair) String() string {
	return fmt.Sprintf("%s/%s", p.BT.String(), p.QT.String())
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
	id := fmt.Sprintf("%s/%s", bt, qt)

	if p, exist := s.pairs[id]; exist {
		return p, nil
	}
	return pair{}, errors.Wrap(errors.ErrNotFound)
}
