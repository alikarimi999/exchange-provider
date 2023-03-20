package kucoin

import (
	"exchange-provider/internal/entity"
	"sync"

	"github.com/pkg/errors"
)

type supportedCoins struct {
	mux   *sync.Mutex
	coins map[string]*kuToken // map[coin.Id+chain.Id]*withdrawalCoin
}

func newSupportedCoins() *supportedCoins {
	return &supportedCoins{
		mux:   &sync.Mutex{},
		coins: make(map[string]*kuToken),
	}
}

func (s *supportedCoins) add(coins map[string]*kuToken) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for id, wc := range coins {
		s.coins[id] = wc
	}

}

func (s *supportedCoins) get(coinId, chainId string) (*kuToken, error) {
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

func (k *kucoinExchange) withdrawalOpts(c *entity.Token, tag string) (map[string]string, error) {

	opts := map[string]string{}

	need, err := k.supportedCoins.needChain(c.Symbol, c.Standard)
	if err != nil {
		return nil, err
	}

	if need {
		opts["chain"] = c.Standard
	}
	opts["memo"] = tag
	return opts, nil

}
