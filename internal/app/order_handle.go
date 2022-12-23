package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"exchange-provider/pkg/errors"
)

type orderHandler struct {
	repo entity.OrderRepo

	ouc *OrderUseCase
	pc  entity.PairConfigs
	oc  entity.OrderCache
	wc  entity.WithdrawalCache
	*exStore

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

		fee: fee,

		l: l,
	}
	return oh
}

func (h *orderHandler) handle(o *entity.Order) {
	const op = errors.Op("User-Order-Handler.handle")

	ex, err := h.exStore.get(o.Routes[0].Exchange)
	if err != nil {
		o.Deposit.FailedDesc = err.Error()
		h.ouc.write(o.Deposit)
		return
	}

	done := make(chan struct{})
	pCh := make(chan bool)
	go ex.TrackDeposit(o, done, pCh)

	<-done
	if err := h.ouc.write(o.Deposit); err != nil {
		h.l.Error(string(op), err.Error())
		pCh <- false
		return
	}
	pCh <- true
	if o.Deposit.Status != entity.DepositConfirmed {
		o.FailedCode = entity.FCDepositFailed
		if err := h.ouc.write(o); err != nil {
			h.l.Error(string(op), err.Error())
		}
		return
	}

	if err := h.ouc.write(o); err != nil {
		h.l.Error(string(op), err.Error())
		return
	}

	for i, route := range o.SortedRoutes() {
		ex, err := h.exStore.get(route.Exchange)
		if err != nil {
			h.l.Error(string(op), err.Error())
			return
		}

		if i == 0 {
			o.Swaps[i].InAmount = o.Deposit.Volume
		} else {
			o.Swaps[i].InAmount = o.Swaps[i-1].OutAmount
		}

		id, err := ex.Exchange(o, i)
		if err != nil {
			o.Status = entity.OSFailed
			o.FailedCode = entity.FCExOrdFailed
			o.FailedDesc = err.Error()
			h.l.Error(string(op), err.Error())
			if err := h.ouc.write(o); err != nil {
				h.l.Error(string(op), err.Error())

			}
			return
		}

		o.Swaps[i].TxId = id
		o.Status = entity.OSWaitForSwapConfirm
		if err = h.ouc.write(o); err != nil {
			h.l.Error(string(op), err.Error())
			return
		}

		done := make(chan struct{})
		pCh := make(chan bool)

		go ex.TrackExchangeOrder(o, i, done, pCh)
		<-done

		switch o.Swaps[i].Status {
		case entity.SwapSucceed:
			if i == (len(o.Routes) - 1) {
				o.Status = entity.OSSwapConfirmed
			}

			if err = h.ouc.write(o); err != nil {
				h.l.Error(string(op), err.Error())
				pCh <- false
				return
			}
			pCh <- true

			if i < (len(o.Routes) - 1) {
				continue
			}

			if err := h.applySpreadAndFee(o, route); err != nil {
				h.l.Error(string(op), err.Error())
				if err := h.ouc.write(o); err != nil {
					h.l.Error(string(op), err.Error())
				}
				return
			}

			id, err = ex.Withdrawal(o)
			if err != nil {
				o.Status = entity.OSFailed
				o.FailedCode = entity.FCInternalError
				o.FailedDesc = err.Error()

				h.l.Error(string(op), err.Error())

				if err := h.ouc.write(o); err != nil {
					h.l.Error(string(op), err.Error())
				}
				return

			}

			o.Withdrawal.TxId = id
			o.Withdrawal.Status = entity.WithdrawalPending
			o.Status = entity.OSWaitForWithdrawalConfirm

			if err = h.ouc.write(o); err != nil {
				h.l.Error(string(op), err.Error())
				return
			}

			if err := h.wc.AddPendingWithdrawal(o.Id); err != nil {
				h.l.Error(string(op), err.Error())
			}
			return

		case entity.SwapFailed:
			o.Status = entity.OSFailed
			o.FailedCode = entity.FCExOrdFailed
			if err := h.ouc.write(o); err != nil {
				h.l.Error(string(op), err.Error())
				pCh <- false
				return
			}
			pCh <- true
			return
		}

	}
}
