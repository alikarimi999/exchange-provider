package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"
	"time"

	"order_service/pkg/errors"
)

type chainTicker struct {
	chain       *entity.Chain
	cache       entity.WithdrawalCache
	ticker      *time.Ticker
	tracker     *withdrawalTracker
	windowsSize time.Duration
	stopCh      chan struct{}
	l           logger.Logger
}

func (h *withdrawalHandler) newChainTicker(c *entity.Chain) *chainTicker {
	return &chainTicker{
		chain:       c,
		cache:       h.tracker.wc,
		ticker:      time.NewTicker(c.BlockTime),
		tracker:     h.tracker,
		windowsSize: time.Minute * 5,
		stopCh:      make(chan struct{}),
		l:           h.l,
	}
}

func (ti *chainTicker) tick(wg *sync.WaitGroup) {
	const op = errors.Op("chainTicker.tick")
	ti.l.Debug(string(op), fmt.Sprintf("Started ticker for chain: '%s'", ti.chain.Id))

	defer wg.Done()
	for {
		select {
		case t := <-ti.ticker.C:
			ws, err := ti.cache.GetPendingWithdrawals(ti.chain.Id, t.Add(-ti.windowsSize))
			if err != nil {
				ti.l.Error(string(op), errors.Wrap(err, op, "pending withdrawals").Error())
				continue
			}
			for _, w := range ws {
				ti.tracker.track(w)
			}
		case <-ti.stopCh:
			ti.l.Debug(string(op), fmt.Sprintf("Stopped ticker for chain: '%s'", ti.chain.Id))
			return
		}
	}
}

func (ti *chainTicker) stop() {
	ti.ticker.Stop()
	close(ti.stopCh)
}
