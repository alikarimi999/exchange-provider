package app

import (
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"
	"time"

	"order_service/pkg/errors"
)

type chainTicker struct {
	chain       entity.Chain
	cache       entity.WithdrawalCache
	ticker      *time.Ticker
	tracker     *withdrawalTracker
	windowsSize time.Duration

	l logger.Logger
}

func (ti *chainTicker) tick(wg *sync.WaitGroup) {
	const op = errors.Op("chainTicker.tick")

	defer wg.Done()
	for {
		select {
		case t := <-ti.ticker.C:
			ws, err := ti.cache.GetPendingWithdrawals(ti.chain, t.Add(-ti.windowsSize))
			if err != nil {
				ti.l.Error(string(op), errors.Wrap(err, op, "pending withdrawals").Error())
				continue
			}
			for _, w := range ws {
				ti.tracker.track(w)
			}
		}
	}
}
