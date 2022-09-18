package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
)

type exTrackerFeed struct {
	o    *entity.UserOrder
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
			go f.ex.TrackExchangeOrder(f.o, done, pCh)
			<-done
			f.done <- struct{}{}
			pCh <- (<-f.pCh)
		}(feed)
	}
}

func (h *exOrderTracker) track(feed *exTrackerFeed) {
	h.feedCh <- feed
}
