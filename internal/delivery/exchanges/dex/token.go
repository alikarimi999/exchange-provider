package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"sync"
)

type supportedTokens struct {
	mux    *sync.Mutex
	tokens map[string]types.Token
}

func newSupportedTokens() *supportedTokens {
	return &supportedTokens{
		mux:    &sync.Mutex{},
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
	s.mux.Lock()
	defer s.mux.Unlock()

	t, ok := s.tokens[symbol]
	if ok {
		return t, nil
	}
	return types.Token{}, errors.Wrap(errors.ErrNotFound, "Token not found")
}
