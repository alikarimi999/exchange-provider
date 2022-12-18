package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"exchange-provider/pkg/errors"
)

type withdrawalTracker struct {
	wCh chan int64

	ouc *OrderUseCase

	repo entity.OrderRepo
	oc   entity.OrderCache
	wc   entity.WithdrawalCache
	exs  *exStore

	list []int64

	l logger.Logger
}

func newWithdrawalTracker(ouc *OrderUseCase, repo entity.OrderRepo, oc entity.OrderCache, wc entity.WithdrawalCache, exs *exStore, l logger.Logger) *withdrawalTracker {
	w := &withdrawalTracker{
		wCh:  make(chan int64, 1024),
		ouc:  ouc,
		repo: repo,
		oc:   oc,
		wc:   wc,
		exs:  exs,
		l:    l,
	}
	return w
}

func (t *withdrawalTracker) run(wg *sync.WaitGroup) {
	const op = errors.Op("Withdrawal-Tracker.run")
	const agent = "WithdrawalTracker"
	defer wg.Done()

	for wd := range t.wCh {
		go func(oId int64) {
			defer t.delete(oId)
			o := &entity.Order{Id: oId}
			if err := t.ouc.read(o); err != nil {
				t.l.Error(agent, errors.Wrap(err, op,
					fmt.Sprintf("order: '%d'", oId)).Error())
				return
			}

			ex, err := t.exs.get(o.Routes[len(o.Routes)-1].Exchange)
			if err != nil {
				t.l.Error(agent, errors.Wrap(err, op, "exchange not found").Error())
				return
			}

			done := make(chan struct{})
			pCh := make(chan bool)

			t.l.Debug(agent, fmt.Sprintf("order %d", oId))
			go ex.TrackWithdrawal(o, done, pCh)

			<-done
			t.l.Debug(agent, fmt.Sprintf("orderId: '%d', staus: '%s'", oId, o.Withdrawal.Status))

			switch o.Withdrawal.Status {
			case entity.WithdrawalPending:
				pCh <- true
				return
			case entity.WithdrawalSucceed:

				o.Status = entity.OSSucceed

				if err := t.ouc.write(o); err != nil {
					t.l.Error(agent, errors.Wrap(err, op, o.String()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.oc.Delete(oId); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", oId)).Error())
				}

				if err := t.wc.DelPendingWithdrawal(oId); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", oId)).Error())
				}

				return

			case entity.WithdrawalFailed:

				o.Status = entity.OSFailed
				o.FailedCode = entity.FCWithdFailed

				if err := t.ouc.write(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.wc.DelPendingWithdrawal(oId); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", oId)).Error())
				}
				return

			}

		}(wd)
	}
}

func (t *withdrawalTracker) track(id int64) {
	var exists bool
	for _, v := range t.list {
		if v == id {
			exists = true
		}
	}
	if !exists {
		t.list = append(t.list, id)
		t.wCh <- id
	}
}

func (t *withdrawalTracker) delete(id int64) {
	for i, v := range t.list {
		if v == id {
			t.list = append(t.list[:i], t.list[i+1:]...)
		}
	}
}
