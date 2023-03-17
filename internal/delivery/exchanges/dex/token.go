package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"sync"
)

type supportedTokens struct {
	mux    *sync.RWMutex
	tokens map[string]types.Token
}

func newSupportedTokens() *supportedTokens {
	return &supportedTokens{
		mux:    &sync.RWMutex{},
		tokens: make(map[string]types.Token),
	}
}

func (s *supportedTokens) add(ts ...types.Token) {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, t := range ts {
		s.tokens[t.Symbol] = t
	}
}
func (s *supportedTokens) get(symbol string) (types.Token, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	t, ok := s.tokens[symbol]
	if ok {
		return t, nil
	}
	return types.Token{}, errors.Wrap(errors.ErrNotFound, "Token not found")
}

func (s *supportedTokens) getAll() []*types.Token {
	s.mux.RLock()
	defer s.mux.RUnlock()
	ts := []*types.Token{}
	for _, t := range s.tokens {
		ts = append(ts, t.SnapShot())
	}
	return ts
}
