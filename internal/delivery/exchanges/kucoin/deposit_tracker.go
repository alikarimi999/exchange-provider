package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
	"time"

	"exchange-provider/pkg/try"
)

type dtFeed struct {
	d         *entity.Deposit
	blockTime time.Duration
	confirms  int64
	done      chan<- struct{}
	pCh       <-chan bool
}
type depositTracker struct {
	fCh chan *dtFeed
	k   *kucoinExchange
	c   *cache
	l   logger.Logger
}

func newDepositTracker(k *kucoinExchange, c *cache) *depositTracker {
	return &depositTracker{
		k:   k,
		c:   c,
		l:   k.l,
		fCh: make(chan *dtFeed),
	}
}

func (t *depositTracker) run(wg *sync.WaitGroup, stopCh chan struct{}) {
	agent := fmt.Sprintf("%s.depositTracker.run", t.k.NID())

	defer wg.Done()

	for {
		select {
		case feed := <-t.fCh:

			go func(f *dtFeed) {
				err := try.Do(10, func(attempt uint64) (bool, error) {
					d, err := t.c.GetD(f.d.TxId)
					if err == nil {
						if !d.MatchCurrency(f.d) {
							f.d.Status = entity.DepositFailed
							f.d.FailedDesc = fmt.Sprintf("currency mismatch, user: `%s`, exchange: `%s` ",
								f.d.CoinId, d.Currency)
							f.done <- struct{}{}
							return false, nil
						}
						f.d.Status = entity.DepositConfirmed
						f.d.Volume = d.Volume
						f.done <- struct{}{}
						return false, nil

					}

					t.l.Debug(agent, fmt.Sprintf("attempt: `%d`, txId: `%s`", attempt, f.d.TxId))
					t := f.blockTime * time.Duration(f.confirms)
					time.Sleep(t / 2)

					return true, err
				})

				if err != nil {
					t.l.Debug(agent, err.Error())
					f.d.Status = entity.DepositFailed
					f.d.FailedDesc = err.Error()
					f.done <- struct{}{}
				}
				// remove the deposit from the cache if tracker's signal successfuly proccessed by consumer.
				if <-f.pCh {
					if err := t.c.RemoveD(f.d.TxId); err != nil {
						t.l.Error(agent, errors.Wrap(agent, err, fmt.Sprintf("TxId: `%s`", f.d.TxId)).Error())
					}
				}
			}(feed)

		case <-stopCh:
			t.l.Debug(agent, "stopped")
			return
		}

	}
}
