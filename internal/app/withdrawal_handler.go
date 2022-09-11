package app

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"order_service/pkg/logger"
	"sync"
	"time"
)

type withdrawalHandler struct {
	tracker     *withdrawalTracker
	wg          *sync.WaitGroup
	ticker      *time.Ticker
	cache       entity.WithdrawalCache
	windowsSize time.Duration
	l           logger.Logger
}

func newWithdrawalHandler(ouc *OrderUseCase, repo entity.OrderRepo, oc entity.OrderCache,
	wc entity.WithdrawalCache, exs *exStore, l logger.Logger) *withdrawalHandler {

	w := &withdrawalHandler{

		tracker:     newWithdrawalTracker(ouc, repo, oc, wc, exs, l),
		ticker:      time.NewTicker(time.Minute * 1),
		windowsSize: time.Minute * 1,
		cache:       wc,
		wg:          &sync.WaitGroup{},

		l: l,
	}

	return w
}

func (wh *withdrawalHandler) handle(wg *sync.WaitGroup) {
	const op = errors.Op("chainTicker.tick")

	defer wg.Done()
	for t := range wh.ticker.C {
		ws, err := wh.cache.GetPendingWithdrawals(t.Add(-wh.windowsSize))
		if err != nil {
			wh.l.Error(string(op), errors.Wrap(err, op, "pending withdrawals").Error())
			continue
		}
		for _, w := range ws {
			wh.tracker.track(w)
		}

	}

}
