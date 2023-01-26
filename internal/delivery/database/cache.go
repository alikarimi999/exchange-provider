package database

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type cache struct {
	mux *sync.RWMutex
	ws  map[string]struct{ t time.Time }
	t   *time.Ticker
}

func newCache(db *mongo.Database) *cache {
	c := &cache{
		mux: &sync.RWMutex{},
		ws:  make(map[string]struct{ t time.Time }),
		t:   time.NewTicker(5 * time.Hour),
	}
	go c.run()
	return c
}

func (c *cache) run() {
	for range c.t.C {
		c.mux.Lock()
		for id, w := range c.ws {
			if time.Now().After(w.t.Add(5 * time.Hour)) {
				delete(c.ws, id)
			}
		}
		c.mux.Unlock()
	}
}

func (c *cache) addPendingWithdrawal(ids ...string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, id := range ids {
		c.ws[id] = struct{ t time.Time }{t: time.Now()}
	}
	return nil
}

func (c *cache) getPendingWithdrawals(end time.Time) ([]string, error) {
	ws := []string{}
	c.mux.RLock()
	defer c.mux.RUnlock()
	for id, w := range c.ws {
		if w.t.Before(end) {
			ws = append(ws, id)
		}
	}
	return ws, nil
}

func (c *cache) delPendingWithdrawal(orderId string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.ws, orderId)
	return nil
}
