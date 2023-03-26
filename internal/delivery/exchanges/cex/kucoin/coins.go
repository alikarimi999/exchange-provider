package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"
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

func (s *supportedCoins) add(ts []*Token) {
	s.mux.Lock()
	defer s.mux.Unlock()

	for _, t := range ts {
		s.coins[s.tId(t.TokenId, t.Network)] = t
	}
}

func (s *supportedCoins) get(coinId, chainId string) (*Token, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	t, exist := s.coins[s.tId(coinId, chainId)]
	if exist {
		return t, nil
	}
	return nil, errors.New("coin not found")
}

func (s *supportedCoins) needChain(coinId, chainId string) (bool, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	wc, exist := s.coins[s.tId(coinId, chainId)]
	if exist {
		return wc.NeedChain, nil
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

func (*supportedCoins) tId(coindId, chainId string) string {
	return fmt.Sprintf("%s-%s", coindId, chainId)
}
