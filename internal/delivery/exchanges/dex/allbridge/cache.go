package allbridge

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"fmt"
	"strings"
	"sync"
	"time"
)

type cache struct {
	ex        *exchange
	mux       *sync.RWMutex
	logs      map[string]*types.TokensReceivedLog
	lastBlock map[string]uint64
	limitLog  uint64
}

func newCache(ex *exchange, fromDB bool) (*cache, error) {
	c := &cache{
		ex:        ex,
		mux:       &sync.RWMutex{},
		logs:      make(map[string]*types.TokensReceivedLog),
		lastBlock: make(map[string]uint64),
		limitLog:  50000,
	}
	if fromDB {
		if err := c.downloadPreviousLogs(); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *cache) downloadPreviousLogs() error {

	wg := &sync.WaitGroup{}
	var err0 error
	for id, n := range c.ex.ns {
		wg.Add(1)
		go func(id string, n types.Network) {
			defer wg.Done()
			lb, err := c.ex.cfg.Networks.network(id).client.BlockNumber(context.Background())
			if err != nil {
				err0 = err
				return
			}
			c.lastBlock[id] = lb
			for i := 1; i <= 2; i++ {
				fromBlock := lb - (c.limitLog * uint64(i))
				toBlock := lb - (c.limitLog * uint64(i-1))
				ls, _, err := n.DownloadLogs(fromBlock, toBlock)
				if err != nil {
					err0 = err
					return
				}
				c.mux.Lock()
				for _, l := range ls {
					lId := logId(id, l.Nonce.String())
					if _, ok := c.logs[lId]; ok {
						continue
					}
					c.logs[lId] = l
				}
				c.mux.Unlock()
			}
		}(id, n)
	}
	wg.Wait()
	return err0
}

func (c *cache) run(stopCh <-chan struct{}) {
	agent := c.ex.agent("cache.run")
	t := time.NewTicker(3 * time.Second)
	dt := time.NewTicker(30 * time.Hour)
	for {
		select {
		case <-t.C:
			for id, n := range c.ex.ns {
				ls, lb, err := n.DownloadLogs(c.lastBlock[id], 0)
				if err == nil {
					for _, l := range ls {
						lId := logId(id, l.Nonce.String())
						if _, ok := c.logs[lId]; ok {
							continue
						}

						c.logs[lId] = l
					}
					if lb > 0 {
						c.lastBlock[id] = lb + 1
					}
				} else {
					if !strings.Contains(err.Error(), "Gateway Time") {
						c.ex.l.Debug(agent, fmt.Sprintf("%s: %s", id, err.Error()))
					}
				}
			}

		case <-dt.C:
			c.mux.Lock()
			threshold := time.Now().Add(-3 * time.Hour)
			for i, l := range c.logs {
				if l.DownloadAt.After(threshold) {
					delete(c.logs, i)
				}
			}
			c.mux.Unlock()

		case <-stopCh:
			return
		}
	}
}
func logId(network, nonce string) string {
	return network + "-" + nonce
}

func (c *cache) getRecievedLog(network, nonce string) (*types.TokensReceivedLog, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	l, ok := c.logs[logId(network, nonce)]
	return l, ok
}
