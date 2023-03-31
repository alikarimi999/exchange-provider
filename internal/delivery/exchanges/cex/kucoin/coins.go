package kucoin

import (
	"exchange-provider/internal/entity"
	"sync"

	"github.com/pkg/errors"
)

type supportedCoins struct {
	mux   *sync.RWMutex
	coins map[string]*Token // map[coin.Id+chain.Id]*withdrawalCoin
}

func newSupportedCoins() *supportedCoins {
	return &supportedCoins{
		mux:   &sync.RWMutex{},
		coins: make(map[string]*Token),
	}
}

func (s *supportedCoins) add(ts []*entity.Token) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, t := range ts {
		s.coins[t.String()] = t.ET.(*Token)
	}
}

func (s *supportedCoins) get(tokenId string) (*Token, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	t, exist := s.coins[tokenId]
	if exist {
		return t, nil
	}
	return nil, errors.New("coin not found")
}
