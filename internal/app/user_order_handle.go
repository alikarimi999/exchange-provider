package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"exchange-provider/pkg/errors"
)

// handle user orders in multiple steps
// steps:
// 1. open a new order in exchange to exchange user provided coin to requested coin
// 2. track the exchange order
// 3. if exchange order succeed, withdrawal the requested coin to user provided address
// 4. if the exchange return the withdrawal id, add the withdrawal to withdrawal cache and withdrawal handler proccess will track it's status

type orderHandler struct {
	repo entity.OrderRepo

	ouc *OrderUseCase
	pc  entity.PairConfigs
	oc  entity.OrderCache
	wc  entity.WithdrawalCache
	*exStore
	eTracker *exOrderTracker
	oCh      chan *entity.UserOrder
	fee      entity.FeeService

	l logger.Logger
}

func newOrderHandler(ouc *OrderUseCase, repo entity.OrderRepo, oc entity.OrderCache, pc entity.PairConfigs, wc entity.WithdrawalCache, fee entity.FeeService, exs *exStore, l logger.Logger) *orderHandler {
	oh := &orderHandler{
		repo: repo,

		ouc:      ouc,
		pc:       pc,
		oc:       oc,
		wc:       wc,
		exStore:  exs,
		eTracker: newExOrderTracker(oc, l),
		oCh:      make(chan *entity.UserOrder, 1024),
		fee:      fee,

		l: l,
	}
	return oh
}

func (o *orderHandler) run(wg *sync.WaitGroup) {
	defer wg.Done()
	const op = errors.Op("User-Order-Handler.run")

	go o.eTracker.run()

	for order := range o.oCh {

		go func(ord *entity.UserOrder) {
			// o.l.Debug(string(op), fmt.Sprintf("handle order: '%d' for user: '%d'", ord.Id, ord.UserId))
			exc, err := o.exStore.get(ord.Exchange)
			if err != nil {
				o.l.Error(string(op), fmt.Sprintf("failed to get exchange: '%s' due to error: ( %s )", ord.Exchange, err.Error()))
				return
			}
			ex := exc.Exchange

			var size string
			var funds string
			if ord.Side == "buy" {
				aVol, sVol, rate, err := o.pc.ApplySpread(ord.BC, ord.QC, ord.Deposit.Volume)
				if err != nil {
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCInternalError
					ord.FailedDesc = err.Error()

					o.l.Error(string(op), err.Error())
					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), fmt.Sprintf("failed to write order: '%s' due to error: ( %s )", ord.String(), err.Error()))
					}
					return
				}
				funds = aVol
				ord.SpreadVol = sVol
				ord.SpreadRate = rate
			} else {
				size = ord.Deposit.Volume
			}
			// 1. open a new order in exchange to exchange user provided coin to requested coin
			id, err := ex.Exchange(ord, size, funds)
			if err != nil {
				ord.Status = entity.OSFailed
				ord.FailedCode = entity.FCExOrdFailed
				ord.FailedDesc = err.Error()

				o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())

				if err := o.ouc.write(ord); err != nil {
					o.l.Error(string(op), fmt.Sprintf("failed to write order: '%s' due to error: ( %s )", ord.String(), err.Error()))
				}
				return

			}

			ord.ExchangeOrder.ExId = id
			ord.Status = entity.OSWaitForExchangeOrderConfirm
			if err = o.ouc.write(ord); err != nil {
				o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
				return
			}

			ef := &exTrackerFeed{
				o:    ord,
				ex:   ex,
				done: make(chan struct{}),
				pCh:  make(chan bool),
			}

			go o.eTracker.track(ef)
			// o.l.Debug(string(op), fmt.Sprintf(" waiting for exchange order: '%s' confirmation", ord.ExchangeOrder.ExId))

			<-ef.done
			switch ord.ExchangeOrder.Status {
			case entity.ExOrderSucceed:

				ord.Status = entity.OSExchangeOrderConfirmed
				// o.l.Debug(string(op), fmt.Sprintf("exchange order: '%s' confirmed", ord.ExchangeOrder.ExId))
				if err = o.ouc.write(ord); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					ef.pCh <- false
					return
				}
				ef.pCh <- true

				var wc *entity.Coin
				switch ord.Side {
				case "buy":
					ord.Withdrawal.Total = ord.ExchangeOrder.Size
					wc = ord.BC

				case "sell":
					wc = ord.QC

					aVol, sVol, rate, err := o.pc.ApplySpread(ord.BC, ord.QC, ord.ExchangeOrder.Funds)
					if err != nil {
						ord.Status = entity.OSFailed
						ord.FailedCode = entity.FCInternalError
						ord.FailedDesc = err.Error()

						o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
						if err := o.ouc.write(ord); err != nil {
							o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
						}
						return
					}
					ord.SpreadRate = rate
					ord.Withdrawal.Total = aVol
					ord.SpreadVol = sVol
				}

				r, f, err := o.fee.ApplyFee(ord.UserId, ord.Withdrawal.Total)
				if err != nil {
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCInternalError
					ord.FailedDesc = err.Error()

					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())

					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					}
					return

				}

				// o.l.Debug(string(op), fmt.Sprintf("order: %d  transferring '%s' %v to '%s'", ord.Id, r, wc, ord.Withdrawal.Address))
				ord.Withdrawal.Coin = wc
				ord.Withdrawal.Fee = f
				id, err = ex.Withdrawal(ord, wc, ord.Withdrawal.Address, r)
				if err != nil {
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCInternalError
					ord.FailedDesc = err.Error()

					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())

					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					}
					return

				}

				ord.Withdrawal.WId = id
				ord.Withdrawal.Status = entity.WithdrawalPending
				ord.Status = entity.OSWaitForWithdrawalConfirm

				// o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' withdrawal order: '%s' created", ord.Id, ord.UserId, ord.Withdrawal.WId))
				if err = o.ouc.write(ord); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					return
				}

				// add to withdrawal cache
				// and wait for withdrawal confirm
				if err := o.wc.AddPendingWithdrawal(ord.Withdrawal); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.Withdrawal.String()).Error())
				}
				// o.l.Debug(string(op), fmt.Sprintf("order: '%d' for user: '%d' is waiting for withdrawal: '%s' to confirm", ord.Id, ord.UserId, ord.Withdrawal.WId))
				return

			case entity.ExOrderPending:
				ef.pCh <- true
				o.handle(ord)
				return

			case entity.ExOrderFailed:
				ord.Status = entity.OSFailed
				ord.FailedCode = entity.FCExOrdFailed
				if err := o.ouc.write(ord); err != nil {
					o.l.Error(string(op), errors.Wrap(err, op, ord.String()).Error())
					ef.pCh <- false
					return
				}
				ef.pCh <- true
				return
			}

		}(order)

	}

}

func (h *orderHandler) handle(o *entity.UserOrder) {
	h.oCh <- o
}
