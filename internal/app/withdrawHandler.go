package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"sync"
	"time"
)

type withdrawalHandler struct {
	tracker     *withdrawalTracker
	r           entity.OrderRepo
	wg          *sync.WaitGroup
	ticker      *time.Ticker
	windowsSize time.Duration
	l           logger.Logger
}

func newWithdrawalHandler(ouc *OrderUseCase, repo entity.OrderRepo,
	exs *exStore, l logger.Logger) *withdrawalHandler {

	w := &withdrawalHandler{
		r:           repo,
		tracker:     newWithdrawalTracker(ouc, repo, exs, l),
		ticker:      time.NewTicker(time.Second * 30),
		windowsSize: time.Second * 30,
		wg:          &sync.WaitGroup{},

		l: l,
	}

	return w
}

func (wh *withdrawalHandler) handle() {
	const op = errors.Op("chainTicker.tick")

	for t := range wh.ticker.C {
		oIds, err := wh.r.GetPendingWithdrawals(t.Add(-wh.windowsSize))
		if err != nil {
			wh.l.Error(string(op), errors.Wrap(err, op, "pending withdrawals").Error())
			continue
		}
		for _, oId := range oIds {
			wh.tracker.track(oId)
		}

	}

}
