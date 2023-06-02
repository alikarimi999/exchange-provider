package kucoin

import (
	"exchange-provider/pkg/logger"
	"sync"
	"time"
)

type cache struct {
	k *exchange

	dMux *sync.RWMutex
	ds   map[string]depositRecord
	prD  map[string]struct{ t time.Time }

	t *time.Ticker
	l logger.Logger
}

func newCache(k *exchange, l logger.Logger) *cache {
	c := &cache{
		k: k,

		dMux: &sync.RWMutex{},
		ds:   make(map[string]depositRecord),
		prD:  make(map[string]struct{ t time.Time }),

		t: time.NewTicker(2 * time.Hour),
		l: l,
	}

	go c.run(k.stopCh)
	return c
}

func (c *cache) run(stopCh chan struct{}) {
	agnet := c.k.agent("cache.run")
	for {
		select {
		case <-c.t.C:
			c.dMux.Lock()
			for id, d := range c.ds {
				if time.Now().After(d.DownloadedAt.Add(12 * time.Hour)) {
					delete(c.ds, id)
				}
			}
			c.dMux.Unlock()
			c.dMux.Lock()
			for id, d := range c.prD {
				if time.Now().After(d.t.Add(2 * time.Hour)) {
					delete(c.prD, id)
				}
			}
			c.dMux.Unlock()

		case <-stopCh:
			c.l.Debug(agnet, "stopped")
			return
		}
	}
}

func (c *cache) saveD(de *depositRecord) {
	c.dMux.Lock()
	defer c.dMux.Unlock()
	c.ds[de.TxId] = *de
}

func (c *cache) getD(txid string) (*depositRecord, bool) {
	c.dMux.Lock()
	defer c.dMux.Unlock()
	d, ok := c.ds[txid]
	if !ok {
		return nil, false
	}
	return d.snapshot(), true
}

func (c *cache) removeD(txid string) {
	c.dMux.Lock()
	defer c.dMux.Unlock()
	delete(c.ds, txid)
}

func (c *cache) proccessedD(txid string) {
	c.dMux.Lock()
	defer c.dMux.Unlock()
	c.prD[txid] = struct{ t time.Time }{t: time.Now()}
}
func (c *cache) existsOrProccessedD(txid string) bool {
	c.dMux.Lock()
	defer c.dMux.Unlock()
	_, ok := c.ds[txid]
	if ok {
		return ok
	}
	_, ok = c.prD[txid]
	return ok
}
