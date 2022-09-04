package app

import (
	"order_service/internal/entity"
	"order_service/pkg/logger"
)

type exTrackerFeed struct {
	eo   *entity.ExchangeOrder
	ex   entity.Exchange
	done chan struct{}
	pCh  chan bool
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

	for feed := range h.feedCh {
		go func(f *exTrackerFeed) {
			done := make(chan struct{})
			pCh := make(chan bool)
			go f.ex.TrackOrder(f.eo, done, pCh)
			<-done
			f.done <- struct{}{}
			pCh <- (<-f.pCh)
		}(feed)
	}
}

func (h *exOrderTracker) track(feed *exTrackerFeed) {
	h.feedCh <- feed
}
