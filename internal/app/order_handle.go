package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
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

			for i, route := range ord.SortedRoutes() {
				ex, err := o.exStore.get(route.Exchange)
				if err != nil {
					o.l.Error(string(op), err.Error())
					return
				}

				if i == 0 {
					ord.Swaps[i].InAmount = ord.Deposit.Volume
				} else {
					ord.Swaps[i].InAmount = ord.Swaps[i-1].OutAmount
				}

				id, err := ex.Exchange(ord, i)
				if err != nil {
					ord.Status = entity.OSFailed
					ord.FailedCode = entity.FCExOrdFailed
					ord.FailedDesc = err.Error()
					o.l.Error(string(op), err.Error())
					if err := o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())

					}
					return
				}

				ord.Swaps[i].TxId = id
				ord.Status = entity.OSWaitForSwapConfirm
				if err = o.ouc.write(ord); err != nil {
					o.l.Error(string(op), err.Error())
					return
				}

				done := make(chan struct{})
				pCh := make(chan bool)

				go ex.TrackExchangeOrder(ord, i, done, pCh)
				<-done

				switch ord.Swaps[i].Status {
				case entity.SwapSucceed:
					if i == (len(ord.Routes) - 1) {
						ord.Status = entity.OSSwapConfirmed
					}

					if err = o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())
						pCh <- false
						return
					}
					pCh <- true

					if i < (len(ord.Routes) - 1) {
						continue
					}

					if err := o.applySpreadAndFee(ord, route); err != nil {
						o.l.Error(string(op), err.Error())
						if err := o.ouc.write(ord); err != nil {
							o.l.Error(string(op), err.Error())
						}
						return
					}

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

					ord.Withdrawal.TxId = id
					ord.Withdrawal.Status = entity.WithdrawalPending
					ord.Status = entity.OSWaitForWithdrawalConfirm

					if err = o.ouc.write(ord); err != nil {
						o.l.Error(string(op), err.Error())
						return
					}

					if err := o.wc.AddPendingWithdrawal(ord.Id); err != nil {
						o.l.Error(string(op), err.Error())
					}
					return

				case entity.SwapFailed:
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
