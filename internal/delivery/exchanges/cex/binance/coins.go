package binance

import (
	"context"
	"exchange-provider/pkg/logger"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
)

type coinsList struct {
	mux  *sync.RWMutex
	list []*binance.CoinInfo
	c    *binance.Client
	t    *time.Ticker
	l    logger.Logger
}

func newCoinsLIst(c *binance.Client, l logger.Logger, stopCh <-chan struct{}) (*coinsList, error) {
	cl := &coinsList{
		mux: &sync.RWMutex{},
		c:   c,
		t:   time.NewTicker(time.Minute),
		l:   l,
	}
	if err := cl.downloadCoins(); err != nil {
		return nil, err
	}
	go cl.run(stopCh)
	return cl, nil
}

func (cl *coinsList) run(stopCh <-chan struct{}) {
	select {
	case <-cl.t.C:
		if err := cl.downloadCoins(); err != nil {
			cl.l.Debug("coinList.run", err.Error())
		}
	case <-stopCh:
		return
	}
}

func (cl *coinsList) downloadCoins() error {
	coins, err := cl.c.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		return err
	}

	cl.mux.Lock()
	cl.list = coins
	cl.mux.Unlock()
	return nil
}

func (cl *coinsList) getCoin(coin, network string) (binance.Network, bool) {
	cl.mux.RLock()
	defer cl.mux.RUnlock()

	for _, c := range cl.list {
		if c.Coin == coin {
			for _, n := range c.NetworkList {
				if n.Network == network {
					return n, true
				}
			}
		}
	}
	return binance.Network{}, false
}
