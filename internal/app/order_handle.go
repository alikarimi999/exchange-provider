package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"exchange-provider/pkg/errors"
)

type orderHandler struct {
	repo entity.OrderRepo

	ouc *OrderUseCase
	pc  entity.PairConfigs
	oc  entity.OrderCache
	wc  entity.WithdrawalCache
	*exStore
	oCh chan *entity.Order
	fee entity.FeeService

	l logger.Logger
}

func newOrderHandler(ouc *OrderUseCase, repo entity.OrderRepo, oc entity.OrderCache, pc entity.PairConfigs, wc entity.WithdrawalCache, fee entity.FeeService, exs *exStore, l logger.Logger) *orderHandler {
	oh := &orderHandler{
		repo: repo,

		ouc:     ouc,
		pc:      pc,
		oc:      oc,
		wc:      wc,
		exStore: exs,

		oCh: make(chan *entity.Order, 1024),
		fee: fee,

		l: l,
	}
	return oh
}

func (o *orderHandler) run(wg *sync.WaitGroup) {
	defer wg.Done()
	const op = errors.Op("User-Order-Handler.run")

	for order := range o.oCh {

		go func(ord *entity.Order) {

			for i, route := range ord.Routes {
				exc, err := o.exStore.get(route.Exchange)
				if err != nil {
					o.l.Error(string(op), fmt.Sprintf("failed to get exchange: '%s' due to error: ( %s )",
						route.Exchange, err.Error()))
					return
				}
				ex := exc.Exchange

				if i == 0 {

					aVol, sVol, rate, err := o.pc.ApplySpread(route.Input, route.Input, ord.Deposit.Volume)
					if err != nil {
						ord.Status = entity.OSFailed
						ord.FailedCode = entity.FCInternalError
						ord.FailedDesc = err.Error()

						o.l.Error(string(op), err.Error())
						if err := o.ouc.write(ord); err != nil {
							o.l.Error(string(op), fmt.Sprintf("failed to write order: '%d' due to error: ( %s )",
								ord.Id, err.Error()))
						}
						return
					}
					ord.Swaps[i].InAmount = aVol
					ord.SpreadVol = sVol
					ord.SpreadRate = rate
				}

				id, err := ex.Exchange(ord, i)

				if err != nil {
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCExOrdFailed
					ord.FailedDesc = err.Error()

					o.l.Error(string(op), err.Error())

					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), fmt.Sprintf("failed to write order: '%d' due to error: ( %s )",
							ord.Id, err.Error()))
					}
					return
				}

				ord.Swaps[i].ExId = id
				ord.Status = entity.OSWaitForExchangeOrderConfirm
				if err = o.ouc.write(ord); err != nil {
					o.l.Error(string(op), err.Error())
					return
				}

				done := make(chan struct{})
				pCh := make(chan bool)

				go ex.TrackExchangeOrder(ord, i, done, pCh)
				<-done

				switch ord.Swaps[i].Status {
				case entity.ExOrderSucceed:

					ord.Status = entity.OSExchangeOrderConfirmed
					// o.l.Debug(string(op), fmt.Sprintf("exchange order: '%s' confirmed", ord.ExchangeOrder.ExId))
					if err = o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())
						pCh <- false
						return
					}
					pCh <- true

					if i == 0 && len(ord.Routes) > 1 {
						continue
					}

					ord.Withdrawal.Total = ord.Swaps[i].OutAmount
					r, f, err := o.fee.ApplyFee(ord.UserId, ord.Withdrawal.Total)
					if err != nil {
						ord.Status = entity.OSFailed
						ord.FailedCode = entity.FCInternalError
						ord.FailedDesc = err.Error()

						o.l.Error(string(op), err.Error())

						if err := o.ouc.write(ord); err != nil {
							o.l.Error(string(op), err.Error())
						}
						return

					}

					ord.Withdrawal.Coin = route.Output
					ord.Withdrawal.Executed = r
					ord.Withdrawal.Fee = f

					id, err = ex.Withdrawal(ord)
					if err != nil {
						ord.Status = entity.OSFailed
						ord.FailedCode = entity.FCInternalError
						ord.FailedDesc = err.Error()

						o.l.Error(string(op), err.Error())

						if err := o.ouc.write(ord); err != nil {
							o.l.Error(string(op), err.Error())
						}
						return

					}

					ord.Withdrawal.WId = id
					ord.Withdrawal.Status = entity.WithdrawalPending
					ord.Status = entity.OSWaitForWithdrawalConfirm

					if err = o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())
						return
					}

					// add to withdrawal cache
					// and wait for withdrawal confirm
					if err := o.wc.AddPendingWithdrawal(ord.Withdrawal); err != nil {
						o.l.Error(string(op), err.Error())
					}
					return

				case entity.ExOrderFailed:
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCExOrdFailed
					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())
						pCh <- false
						return
					}
					pCh <- true
					return
				}

			}

		}(order)

	}

}

func (h *orderHandler) handle(o *entity.Order) {
	h.oCh <- o
}
