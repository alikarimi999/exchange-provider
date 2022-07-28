package kucoin

import (
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

type trackerFedd struct {
	eo   *entity.ExchangeOrder
	done chan<- struct{}
	err  chan<- error
}

type orderTracker struct {
	feedCh chan *trackerFedd
	api    *kucoin.ApiService
	l      logger.Logger
}

func newOrderTracker(api *kucoin.ApiService, l logger.Logger) *orderTracker {
	return &orderTracker{
		feedCh: make(chan *trackerFedd, 1024),
		api:    api,
		l:      l,
	}
}

func (t *orderTracker) run(wg *sync.WaitGroup) {
	const op = errors.Op("Kucoin.orderTracker.run")

	t.l.Debug(string(op), "started")

	defer wg.Done()
	for {
		select {
		case feed := <-t.feedCh:
			func(f *trackerFedd) {

				resp, err := t.api.Order(f.eo.Id)
				if err = handleSDKErr(err, resp); err != nil {
					f.err <- errors.Wrap(err, op, errors.ErrInternal)
					return
				}

				order := &kucoin.OrderModel{}
				if err = resp.ReadData(order); err != nil {
					f.err <- errors.Wrap(err, op, errors.ErrInternal)
					return
				}
				f.eo.Symbol = order.Symbol
				f.eo.Funds = order.DealFunds
				f.eo.Size = order.DealSize
				f.eo.Side = order.Side
				f.eo.Fee = order.Fee
				f.eo.FeeCurrency = order.FeeCurrency
				f.eo.Status = entity.ExOrderSucceed
				f.done <- struct{}{}
				return

			}(feed)
		}
	}
}

func (o *orderTracker) track(feed *trackerFedd) {
	o.feedCh <- feed
}
