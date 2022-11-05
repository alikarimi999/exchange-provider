package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"exchange-provider/pkg/errors"
)

type withdrawalTracker struct {
	wCh chan *entity.Withdrawal

	ouc *OrderUseCase

	repo entity.OrderRepo
	oc   entity.OrderCache
	wc   entity.WithdrawalCache
	exs  *exStore

	l logger.Logger
}

func newWithdrawalTracker(ouc *OrderUseCase, repo entity.OrderRepo, oc entity.OrderCache, wc entity.WithdrawalCache, exs *exStore, l logger.Logger) *withdrawalTracker {
	w := &withdrawalTracker{
		wCh:  make(chan *entity.Withdrawal, 1024),
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
		go func(w *entity.Withdrawal) {
			w0 := *w
			ex, err := t.exs.get(w.Exchange)
			if err != nil {
				t.l.Error(agent, errors.Wrap(err, op, "exchange not found").Error())
				return
			}

			done := make(chan struct{})
			pCh := make(chan bool)

			t.l.Debug(agent, fmt.Sprintf("order %d", w.OrderId))
			go ex.TrackWithdrawal(w, done, pCh)

			<-done
			t.l.Debug(agent, fmt.Sprintf("orderId: '%d', staus: '%s'", w.OrderId, w.Status))

			switch w.Status {
			case entity.WithdrawalPending:
				pCh <- true
				return
			case entity.WithdrawalSucceed:
				o := &entity.UserOrder{Id: w.OrderId, UserId: w.UserId}
				if err := t.ouc.read(o); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", w.OrderId)).Error())
					pCh <- false
					return
				}

				o.Status = entity.OSSucceed
				o.Withdrawal.Status = w.Status
				o.Withdrawal.ExchangeFee = w.ExchangeFee
				o.Withdrawal.Executed = w.Executed
				o.Withdrawal.TxId = w.TxId

				if err := t.ouc.write(o); err != nil {
					t.l.Error(agent, errors.Wrap(err, op, o.String()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.oc.Delete(w.UserId, w.OrderId); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", w.OrderId)).Error())
				}

				if err := t.wc.DelPendingWithdrawal(w0); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", w.OrderId)).Error())
				}

				return

			case entity.WithdrawalFailed:

				o := &entity.UserOrder{Id: w.OrderId, UserId: w.UserId}
				if err := t.ouc.read(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", w.OrderId)).Error())
					pCh <- false
					return
				}

				o.Status = entity.OSFailed
				o.FailedCode = entity.FCWithdFailed
				o.Withdrawal.Status = w.Status
				o.Withdrawal.FailedDesc = w.FailedDesc

				if err := t.ouc.write(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.wc.DelPendingWithdrawal(w0); err != nil {
					t.l.Error(agent, errors.Wrap(err, op,
						fmt.Sprintf("order: '%d'", w.OrderId)).Error())
				}
				return

			}

		}(wd)
	}
}

func (t *withdrawalTracker) track(wi *entity.Withdrawal) {
	t.wCh <- wi
}
