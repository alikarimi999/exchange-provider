package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type wtFeed struct {
	w            *entity.Withdrawal
	done         chan<- struct{}
	proccessedCh <-chan bool
}

type withdrawalTracker struct {
	k      *kucoinExchange
	feedCh chan *wtFeed
	l      logger.Logger
	c      *cache
}

func newWithdrawalTracker(k *kucoinExchange, c *cache) *withdrawalTracker {
	return &withdrawalTracker{
		k:      k,
		feedCh: make(chan *wtFeed, 1024),
		l:      k.l,
		c:      c,
	}
}

func (t *withdrawalTracker) run(wg *sync.WaitGroup, stopCh chan struct{}) {
	op := errors.Op(fmt.Sprintf("%s.withdrawalTracker.run", t.k.Id()))
	t.l.Debug(string(op), "started")

	defer wg.Done()
	for {
		select {
		case feed := <-t.feedCh:
			go func(f *wtFeed) {
				txId := f.w.TxId
				wd, err := t.c.getWithdrawal(txId)
				if err != nil {
					if err != redis.Nil {
						t.l.Error(string(op), err.Error())
					}
					f.w.Status = entity.WithdrawalPending
					f.done <- struct{}{}
					<-f.proccessedCh
					return
				}
				switch wd.Status {
				case "SUCCESS":
					f.w.Status = entity.WithdrawalSucceed
					f.w.Fee = wd.Fee
					f.w.TxId = wd.FixTxId() + "-" + txId
					f.w.FeeCurrency = f.w.Token.String()
				case "FAILURE":
					f.w.Status = entity.WithdrawalFailed
					f.w.FailedDesc = "failed by exchange"
				}
				f.done <- struct{}{}

				if <-f.proccessedCh {
					if err := t.c.delWithdrawal(txId); err != nil {
						t.l.Error(string(op), errors.Wrap(err, op).Error())
					}
					if err := t.c.proccessedWithdrawal(txId); err != nil {
						t.l.Error(string(op), errors.Wrap(err, op).Error())
					}
				}
			}(feed)

		case <-stopCh:
			t.l.Debug(string(op), "stopped")
			return

		}
	}
}

func (t *withdrawalTracker) track(f *wtFeed) {
	t.feedCh <- f
}
