package uniswapv3

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type token struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Address  common.Address `json:"address"`
	Decimals int            `json:"decimals"`
	ChainId  int64          `json:"chainId"`
}

func (t *token) isNative() bool {
	return t.Symbol == ether
}

func (t *token) String() string {
	return t.Symbol
}

func (t *token) ToEntity() *entity.PairCoin {
	return &entity.PairCoin{
		Coin: &entity.Coin{
			CoinId:  t.Symbol,
			ChainId: chainId,
		},
	}
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
