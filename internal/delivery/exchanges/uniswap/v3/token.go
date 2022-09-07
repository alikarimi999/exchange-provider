package uniswapv3

import (
	"order_service/pkg/errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type token struct {
	symbol   string
	address  common.Address
	isNative bool
	decimals int
}

func (t *token) String() string {
	return t.symbol
}

type supportedTokens struct {
	mux    *sync.Mutex
	tokens map[string]token
}

func newSupportedTokens() *supportedTokens {
	return &supportedTokens{
		mux:    &sync.Mutex{},
		tokens: make(map[string]token),
	}
}

func (s *supportedTokens) get(symbol string) (token, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	t, ok := s.tokens[symbol]
	if ok {
		return t, nil
	}
	return token{}, errors.Wrap(errors.ErrNotFound, "Token not found")
}
