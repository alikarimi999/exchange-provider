package kucoin

import (
	"fmt"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"sync"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

type trackerFedd struct {
	eo   *entity.ExchangeOrder
	done chan<- struct{}
	pCh  <-chan bool
}

type orderTracker struct {
	k      *kucoinExchange
	feedCh chan *trackerFedd
	api    *kucoin.ApiService
	l      logger.Logger
}

func newOrderTracker(k *kucoinExchange, api *kucoin.ApiService, l logger.Logger) *orderTracker {
	return &orderTracker{
		k:      k,
		feedCh: make(chan *trackerFedd, 1024),
		api:    api,
		l:      l,
	}
}

func (t *orderTracker) run(wg *sync.WaitGroup, stopCh chan struct{}) {
	op := errors.Op(fmt.Sprintf("%s.orderTracker.run", t.k.NID()))

	t.l.Debug(string(op), "started")

	defer wg.Done()
	for {
		select {
		case feed := <-t.feedCh:
			func(f *trackerFedd) {

				resp, err := t.api.Order(f.eo.ExId)
				if err = handleSDKErr(err, resp); err != nil {
					t.l.Error(string(op), err.Error())
					f.eo.Status = entity.ExOrderFailed
					f.eo.FailedDesc = err.Error()
					f.done <- struct{}{}
					<-f.pCh
					return
				}

				order := &kucoin.OrderModel{}
				if err = resp.ReadData(order); err != nil {
					t.l.Error(string(op), err.Error())
					f.eo.Status = entity.ExOrderFailed
					f.eo.FailedDesc = err.Error()
					f.done <- struct{}{}
					<-f.pCh
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
				<-f.pCh

			}(feed)

		case <-stopCh:
			t.l.Debug(string(op), "stopped")
			return
		}
	}
}

func (o *orderTracker) track(feed *trackerFedd) {
	o.feedCh <- feed
}
