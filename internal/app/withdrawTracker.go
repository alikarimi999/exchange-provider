package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"

	"exchange-provider/pkg/errors"
)

type withdrawalTracker struct {
	wCh chan *entity.ObjectId

	ouc *OrderUseCase

	r   entity.OrderRepo
	exs *exStore

	list []string

	l logger.Logger
}

func newWithdrawalTracker(ouc *OrderUseCase, repo entity.OrderRepo, exs *exStore, l logger.Logger) *withdrawalTracker {
	w := &withdrawalTracker{
		wCh: make(chan *entity.ObjectId, 1024),
		ouc: ouc,
		r:   repo,
		exs: exs,
		l:   l,
	}
	return w
}

func (t *withdrawalTracker) run() {
	const op = errors.Op("Withdrawal-Tracker.run")
	const agent = "WithdrawalTracker"

	for wd := range t.wCh {
		go func(oId *entity.ObjectId) {
			defer t.delete(oId)
			o := &entity.CexOrder{ObjectId: oId}
			if err := t.ouc.read(o); err != nil {
				t.l.Error(agent, fmt.Sprintf("order: '%s'", oId))
				return
			}

			ex, err := t.exs.get(o.Routes[len(o.Routes)-1].Exchange)
			if err != nil {
				t.l.Error(agent, errors.Wrap(err, op, "exchange not found").Error())
				return
			}

			done := make(chan struct{})
			pCh := make(chan bool)

			go ex.(entity.Cex).TrackWithdrawal(o, done, pCh)

			<-done
			switch o.Withdrawal.Status {
			case entity.WithdrawalPending:
				pCh <- true
				return
			case entity.WithdrawalSucceed:

				t.l.Debug(agent, fmt.Sprintf("order: '%s', staus: '%s'", oId, o.Withdrawal.Status))
				o.Status = entity.OSucceeded

				if err := t.ouc.write(o); err != nil {
					t.l.Error(agent, errors.Wrap(err, op, o.ID()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.r.DelPendingWithdrawal(oId); err != nil {
					t.l.Error(agent, fmt.Sprintf("order: '%s'", oId))
				}

				return

			case entity.WithdrawalFailed:

				t.l.Debug(agent, fmt.Sprintf("order: '%s', staus: '%s'", oId, o.Withdrawal.Status))
				o.Status = entity.OFailed
				o.FailedCode = entity.FCWithdFailed

				if err := t.ouc.write(o); err != nil {
					t.l.Error(string(op), errors.Wrap(err, op, o.ID()).Error())
					pCh <- false
					return
				}
				pCh <- true

				if err := t.r.DelPendingWithdrawal(oId); err != nil {
					t.l.Error(agent, fmt.Sprintf("order: '%s'", oId))
				}
				return

			}

		}(wd)
	}
}

func (t *withdrawalTracker) track(id *entity.ObjectId) {
	var exists bool
	for _, v := range t.list {
		if v == id.String() {
			exists = true
		}
	}
	if !exists {
		t.list = append(t.list, id.String())
		t.wCh <- id
	}
}

func (t *withdrawalTracker) delete(id *entity.ObjectId) {
	for i, v := range t.list {
		if v == id.String() {
			t.list = append(t.list[:i], t.list[i+1:]...)
		}
	}
}
