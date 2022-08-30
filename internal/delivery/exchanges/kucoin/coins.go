package kucoin

import (
	"order_service/internal/entity"
	"sync"

	"github.com/pkg/errors"
)

type supportedCoins struct {
	mux   *sync.Mutex
	coins map[string]*kuCoin // map[coin.Id+chain.Id]*withdrawalCoin
}

func newSupportedCoins() *supportedCoins {
	return &supportedCoins{
		mux:   &sync.Mutex{},
		coins: make(map[string]*kuCoin),
	}
}

func (s *supportedCoins) add(coins map[string]*kuCoin) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id, wc := range coins {
		s.coins[id] = wc
	}

}

func (s *supportedCoins) get(coinId, chainId string) (*kuCoin, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	wc, exist := s.coins[coinId+chainId]
	if exist {
		return wc, nil
	}
	return nil, errors.New("coin not found")
}

func (s *supportedCoins) needChain(coinId, chainId string) (bool, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	wc, exist := s.coins[coinId+chainId]
	if exist {
		return wc.needChain, nil
	}
	return false, errors.New("coin not found")
}

func (k *kucoinExchange) withdrawalOpts(c *entity.Coin, tag string) (map[string]string, error) {

	opts := map[string]string{}

	need, err := k.supportedCoins.needChain(c.CoinId, c.ChainId)
	if err != nil {
		return nil, err
	}

	if need {
		opts["chain"] = c.ChainId
	}
	opts["memo"] = tag
	return opts, nil

}

func (s *supportedCoins) snapshot() map[string]*kuCoin {
	s.mux.Lock()
	defer s.mux.Unlock()
	res := make(map[string]*kuCoin)
	for id, wc := range s.coins {
		res[id] = &kuCoin{
			WithdrawalPrecision: wc.WithdrawalPrecision,
			needChain:           wc.needChain,
		}
	}

	return res
}

func (s *supportedCoins) purge() {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.coins = make(map[string]*kuCoin)
}
