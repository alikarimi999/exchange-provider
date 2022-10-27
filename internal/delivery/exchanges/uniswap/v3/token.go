package uniswapv3

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	Address  common.Address `json:"address"`
	Decimals int            `json:"decimals"`
	ChainId  int64          `json:"chainId"`
	Native   bool           `json:"native"`
}

func (t *Token) isNative() bool {
	return t.Native
}

func (t *Token) wrappSymbol() string {
	return fmt.Sprintf("W%s", t.Symbol)
}

func (t *Token) String() string {
	return t.Symbol
}

func (t *Token) ToEntity(u *dex) *entity.PairCoin {
	return &entity.PairCoin{
		Coin: &entity.Coin{
			CoinId:  t.Symbol,
			ChainId: u.cfg.TokenStandard,
		},
		BlockTime:       u.blockTime,
		ContractAddress: t.Address.String(),
	}
}

type supportedTokens struct {
	mux    *sync.Mutex
	tokens map[string]Token
}

func newSupportedTokens() *supportedTokens {
	return &supportedTokens{
		mux:    &sync.Mutex{},
		tokens: make(map[string]Token),
	}
}

func (s *supportedTokens) add(ts ...Token) {
	s.mux.Lock()
	defer s.mux.Unlock()
	for _, t := range ts {
		s.tokens[t.Symbol] = t
	}
}
func (s *supportedTokens) get(symbol string) (Token, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	t, ok := s.tokens[symbol]
	if ok {
		return t, nil
	}
	return Token{}, errors.Wrap(errors.ErrNotFound, "Token not found")
}
