package app

import (
	"order_service/internal/entity"
	"sync"
)

type supportedCoins struct {
	mux   *sync.Mutex
	coins map[string]*entity.Coin // map[coinId+chainId]*appCoin
}

func newSupportedCoins() *supportedCoins {
	return &supportedCoins{
		mux:   &sync.Mutex{},
		coins: make(map[string]*entity.Coin),
	}
}

func (s *supportedCoins) add(coins map[string]*entity.Coin) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id, coin := range coins {
		s.coins[id] = coin
	}

}

func (s *supportedCoins) remove(coinId, chainId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	delete(s.coins, coinId+chainId)
}

func (s *supportedCoins) exist(coinId, chainId string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, exist := s.coins[coinId+chainId]
	return exist
}

func (s *supportedCoins) snapshots() []*entity.Coin {
	s.mux.Lock()
	defer s.mux.Unlock()

	coins := make([]*entity.Coin, 0)
	for _, coin := range s.coins {
		coins = append(coins, coin)
	}
	return coins
}
