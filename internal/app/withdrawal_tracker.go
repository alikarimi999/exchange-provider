package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"sync"

	"order_service/pkg/errors"
)

type withdrawalTracker struct {
	wCh  chan *entity.Withdrawal
	repo entity.OrderRepo
	oc   entity.OrderCache
	wc   entity.WithdrawalCache
	exs  *exStore

	l logger.Logger
}

func newWithdrawalTracker(repo entity.OrderRepo, oc entity.OrderCache, wc entity.WithdrawalCache, exs *exStore, l logger.Logger) *withdrawalTracker {
	w := &withdrawalTracker{
		wCh:  make(chan *entity.Withdrawal, 1024),
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

	for {
		select {
		case wd := <-t.wCh:
			go func(w *entity.Withdrawal) {
				t.l.Debug(string(op), fmt.Sprintf("track withdrawal: '%s' order: '%d' user: '%d'", w.Id, w.OrderId, w.UserId))

				ex, err := t.exs.get(w.Exchange)
				if err != nil {
					t.l.Error(string(op), errors.Wrap(err, op, "exchange not found").Error())
					return
				}

				done := make(chan struct{})
				errCh := make(chan error)
				proccessedCh := make(chan bool)
				if err := ex.TrackWithdrawal(w, done, errCh, proccessedCh); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
					return
				}
				select {
				case <-done:

					switch w.Status {
					case entity.WithdrawalPending:
						t.l.Debug(string(op), fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d' is pending yet", w.Id, w.OrderId, w.UserId))
						return
					case entity.WithdrawalSucceed:

						o, err := t.oc.Get(w.UserId, w.OrderId)
						if err != nil {
							t.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
							proccessedCh <- false
							return
						}

						t.l.Debug(string(op), fmt.Sprintf("withdrawal: '%s' status changed to: '%s' , order: %d user: %d",
							w.Id, w.Status, w.OrderId, w.UserId))

						o.Status = entity.OrderStatusSecceed
						o.Withdrawal.Status = entity.WithdrawalSucceed
						o.Withdrawal.ExchangeFee = w.ExchangeFee
						o.Withdrawal.Executed = w.Executed
						o.Withdrawal.TxId = w.TxId

						if o.Side == "buy" {
							o.Size = w.Executed
						} else {
							o.Funds = w.Executed
						}

						if err := t.repo.Add(o); err != nil {
							t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())

							if err := t.oc.Add(o); err != nil {
								t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())
								proccessedCh <- false
								return
							}

							proccessedCh <- true
							if err := t.wc.DelPendingWithdrawal(w); err != nil {
								t.l.Error(string(op), errors.Wrap(err, op,
									fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
							}

							return

						}

						proccessedCh <- true

						if err := t.oc.Delete(w.UserId, w.OrderId); err != nil {
							t.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
						}

						if err := t.wc.DelPendingWithdrawal(w); err != nil {
							t.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
						}

						return

					case entity.WithdrawalFailed:

						o, err := t.oc.Get(w.UserId, w.OrderId)
						if err != nil {
							t.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
							proccessedCh <- false
							return
						}

						t.l.Debug(string(op), fmt.Sprintf("withdrawal: '%s' status changed to: '%s' , order: %d user: %d",
							w.Id, w.Status, w.OrderId, w.UserId))

						o.Broken = true
						o.BrokeReason = "withdrawal failed"
						o.Withdrawal.Status = entity.WithdrawalFailed

						if err := t.oc.Update(o); err != nil {
							t.l.Error(string(op), errors.Wrap(err, op, o.String()).Error())

							proccessedCh <- false
							return
						}
						proccessedCh <- true

						if err := t.wc.DelPendingWithdrawal(w); err != nil {
							t.l.Error(string(op), errors.Wrap(err, op,
								fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
						}
						return

					}

				case err := <-errCh:
					proccessedCh <- false
					t.l.Error(string(op), errors.Wrap(err, op,
						fmt.Sprintf("withdrawalId: '%s' orderId: '%d', userId: '%d'", w.Id, w.OrderId, w.UserId)).Error())
					return

				}
			}(wd)
		}
	}
}

func (t *withdrawalTracker) track(wi *entity.Withdrawal) {
	t.wCh <- wi
}
