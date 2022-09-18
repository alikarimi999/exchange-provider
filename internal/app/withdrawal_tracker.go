package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"
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
	defer wg.Done()

	for wd := range t.wCh {
		go func(w *entity.Withdrawal) {
			w0 := *w
			// t.l.Debug(string(op), fmt.Sprintf("track withdrawal: '%s' order: '%d' user: '%d'", w.WId, w.OrderId, w.UserId))

			ex, err := t.exs.get(w.Exchange)
			if err != nil {
				t.l.Error(string(op), errors.Wrap(err, op, "exchange not found").Error())
				return
			}

			done := make(chan struct{})
			pCh := make(chan bool)
			go ex.TrackWithdrawal(w, done, pCh)

			<-done
			switch w.Status {
			case entity.WithdrawalPending:
				// t.l.Debug(string(op), fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d' is pending yet", w.WId, w.OrderId, w.UserId))
				pCh <- true
				return
			case entity.WithdrawalSucceed:

				o := &entity.UserOrder{Id: w.OrderId, UserId: w.UserId}
				if err := t.ouc.read(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.WId, w.OrderId, w.UserId)).Error())
					pCh <- false
					return
				}

				// t.l.Debug(string(op), fmt.Sprintf("withdrawal: '%s' status changed to: '%s' , order: %d user: %d",
				// w.WId, w.Status, w.OrderId, w.UserId))

				o.Status = entity.OSSucceed
				o.Withdrawal.Status = w.Status
				o.Withdrawal.ExchangeFee = w.ExchangeFee
				o.Withdrawal.Executed = w.Executed
				o.Withdrawal.TxId = w.TxId

				if err := t.ouc.write(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.oc.Delete(w.UserId, w.OrderId); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.WId, w.OrderId, w.UserId)).Error())
				}

				if err := t.wc.DelPendingWithdrawal(w0); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.WId, w.OrderId, w.UserId)).Error())
				}

				return

			case entity.WithdrawalFailed:

				o := &entity.UserOrder{Id: w.OrderId, UserId: w.UserId}
				if err := t.ouc.read(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.WId, w.OrderId, w.UserId)).Error())
					pCh <- false
					return
				}

				t.l.Debug(string(op), fmt.Sprintf("withdrawal: '%s' status changed to: '%s' , order: %d user: %d",
					w.WId, w.Status, w.OrderId, w.UserId))

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
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.WId, w.OrderId, w.UserId)).Error())
				}
				return

			}

		}(wd)
	}
}

func (t *withdrawalTracker) track(wi *entity.Withdrawal) {
	t.wCh <- wi
}
