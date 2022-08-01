package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"

	"order_service/pkg/errors"
)

type exTrackerFeed struct {
	eo      *entity.ExchangeOrder
	ex      entity.Exchange
	succeed chan bool
}

type exOrderTracker struct {
	feedCh chan *exTrackerFeed
	cache  entity.OrderCache

	l logger.Logger
}

func newExOrderTracker(cache entity.OrderCache, l logger.Logger) *exOrderTracker {
	eo := &exOrderTracker{
		feedCh: make(chan *exTrackerFeed, 1024),
		cache:  cache,
		l:      l,
	}
	return eo
}

func (h *exOrderTracker) run() {
	const op = errors.Op("Exchange-Order-Tracker.run")
	for {
		select {
		case feed := <-h.feedCh:
			go func(f *exTrackerFeed) {
				done := make(chan struct{})
				errCh := make(chan error)
				go f.ex.TrackOrder(f.eo, done, errCh)

				select {
				case <-done:
					f.succeed <- true
				case err := <-errCh:
					// TODO: handle the error
					f.succeed <- false
					h.l.Error(string(op), errors.Wrap(err, op, fmt.Sprintf("exchangeOrderId: '%s', orderId: '%d', userId: '%d'", f.eo.ExId, f.eo.OrderId, f.eo.UserId)).Error())
					return
				}

			}(feed)
		}
	}
}

func (h *exOrderTracker) track(feed *exTrackerFeed) {
	h.feedCh <- feed
}
