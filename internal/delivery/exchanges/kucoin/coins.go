package kucoin

import (
	"order_service/internal/entity"
	"sync"

	"github.com/pkg/errors"
)

type withdrawalCoin struct {
	needChain bool
}

type withdrawalCoins struct {
	mux   *sync.Mutex
	coins map[string]*withdrawalCoin // map[coin.Id+chain.Id]*withdrawalCoin
}

func newWithdrawalCoins() *withdrawalCoins {
	return &withdrawalCoins{
		mux:   &sync.Mutex{},
		coins: make(map[string]*withdrawalCoin),
	}
}

func (s *withdrawalCoins) add(coins map[string]*withdrawalCoin) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id, wc := range coins {
		s.coins[id] = wc
	}

}

// func (s *withdrawalCoins) get(coinId, chainId string) (*withdrawalCoin, error) {
// 	s.mux.Lock()
// 	defer s.mux.Unlock()
// 	wc, exist := s.coins[coinId+chainId]
// 	if exist {
// 		return wc, nil
// 	}
// 	return nil, errors.New("coin not found")
// }

func (s *withdrawalCoins) needChain(coinId, chainId string) (bool, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	wc, exist := s.coins[coinId+chainId]
	if exist {
		return wc.needChain, nil
	}
	return false, errors.New("coin not found")
}

func (k *kucoinExchange) withdrawalOpts(c *entity.Coin) (map[string]string, error) {

	opts := map[string]string{}

	need, err := k.withdrawalCoins.needChain(c.Id, c.Chain.Id)
	if err != nil {
		return nil, err
	}

	if need {
		opts["chain"] = c.Chain.Id
	}

	return opts, nil

}