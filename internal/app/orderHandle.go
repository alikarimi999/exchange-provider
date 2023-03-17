package app

// type orderHandler struct {
// 	repo entity.OrderRepo

// 	ouc *OrderUseCase
// 	pc  entity.PairConfigs
// 	*exStore

// 	fee entity.FeeService

// 	l logger.Logger
// }

// func newOrderHandler(ouc *OrderUseCase, repo entity.OrderRepo, pc entity.PairConfigs, fee entity.FeeService, exs *exStore, l logger.Logger) *orderHandler {
// 	oh := &orderHandler{
// 		repo: repo,

// 		ouc:     ouc,
// 		pc:      pc,
// 		exStore: exs,

// 		fee: fee,

// 		l: l,
// 	}
// 	return oh
// }

// func (h *orderHandler) handle(o *entity.CexOrder) {
// 	const op = errors.Op("orderHandler.handle")

// 	ex, err := h.exStore.get(o.Routes[0].Exchange)
// 	if err != nil {
// 		o.Deposit.FailedDesc = err.Error()
// 		h.ouc.write(o.Deposit)
// 		return
// 	}
// 	cex := ex.(entity.Cex)
// 	done := make(chan struct{})
// 	pCh := make(chan bool)
// 	go cex.TrackDeposit(o, done, pCh)

// 	<-done
// 	if o.Deposit.Status == entity.DepositFailed {
// 		o.FailedCode = entity.FCDepositFailed
// 		o.Status = entity.OFailed
// 	}
// 	if err := h.ouc.write(o); err != nil {
// 		h.l.Error(string(op), err.Error())
// 		pCh <- false
// 		return
// 	}
// 	pCh <- true
// 	if o.Deposit.Status != entity.DepositConfirmed {
// 		return
// 	}

// 	for i, route := range o.SortedRoutes() {
// 		ex, err := h.exStore.get(route.Exchange)
// 		if err != nil {
// 			h.l.Error(string(op), err.Error())
// 			return
// 		}
// 		cex = ex.(entity.Cex)
// 		if i == 0 {
// 			o.Swaps[i].InAmount = fmt.Sprintf("%v", o.Deposit.Volume)
// 		} else {
// 			o.Swaps[i].InAmount = o.Swaps[i-1].OutAmount
// 		}

// 		id, err := cex.Swap(o, i)
// 		if err != nil {
// 			o.Status = entity.OFailed
// 			o.FailedCode = entity.FCExOrdFailed
// 			o.FailedDesc = err.Error()
// 			h.l.Error(string(op), err.Error())
// 			if err := h.ouc.write(o); err != nil {
// 				h.l.Error(string(op), err.Error())

// 			}
// 			return
// 		}

// 		o.Swaps[i].TxId = id
// 		o.Status = entity.OWaitForSwapConfirm
// 		if err = h.ouc.write(o); err != nil {
// 			h.l.Error(string(op), err.Error())
// 			return
// 		}

// 		done := make(chan struct{})
// 		pCh := make(chan bool)

// 		go cex.TrackSwap(o, i, done, pCh)
// 		<-done

// 		switch o.Swaps[i].Status {
// 		case entity.SwapSucceed:
// 			if i == (len(o.Routes) - 1) {
// 				o.Status = entity.OSwapConfirmed
// 			}

// 			if err = h.ouc.write(o); err != nil {
// 				h.l.Error(string(op), err.Error())
// 				pCh <- false
// 				return
// 			}
// 			pCh <- true

// 			if i < (len(o.Routes) - 1) {
// 				continue
// 			}

// 			if err := h.applySpreadAndFee(o, route); err != nil {
// 				h.l.Error(string(op), err.Error())
// 				if err := h.ouc.write(o); err != nil {
// 					h.l.Error(string(op), err.Error())
// 				}
// 				return
// 			}

// 			id, err = cex.Withdrawal(o)
// 			if err != nil {
// 				o.Status = entity.OFailed
// 				o.FailedCode = entity.FCInternalError
// 				o.FailedDesc = err.Error()

// 				h.l.Error(string(op), err.Error())

// 				if err := h.ouc.write(o); err != nil {
// 					h.l.Error(string(op), err.Error())
// 				}
// 				return

// 			}

// 			o.Withdrawal.TxId = id
// 			o.Withdrawal.Status = entity.WithdrawalPending
// 			o.Status = entity.OWaitForWithdrawalConfirm

// 			if err = h.ouc.write(o); err != nil {
// 				h.l.Error(string(op), err.Error())
// 				return
// 			}

// 			if err := h.repo.AddPendingWithdrawal(o.ID()); err != nil {
// 				h.l.Error(string(op), err.Error())
// 			}
// 			return

// 		case entity.SwapFailed:
// 			o.Status = entity.OFailed
// 			o.FailedCode = entity.FCExOrdFailed
// 			if err := h.ouc.write(o); err != nil {
// 				h.l.Error(string(op), err.Error())
// 				pCh <- false
// 				return
// 			}
// 			pCh <- true
// 			return
// 		}

// 	}
// }
