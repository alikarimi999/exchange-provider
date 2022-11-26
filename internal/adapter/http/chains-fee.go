package http

import (
	"sync"
)

type chainsFee struct {
	mux   *sync.Mutex
	chain map[string]float64
}

func (c *chainsFee) update(chains map[string]float64) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range chains {
		c.chain[k] = v
	}
}

func (c *chainsFee) lowerEq(c1, c2 string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.chain[c1] <= c.chain[c2]
}
