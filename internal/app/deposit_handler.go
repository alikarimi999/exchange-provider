package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"sync"
)

type depositHandler struct {
	ch chan *entity.Order
	o  *OrderUseCase
	l  logger.Logger
}

func newDepositHandler(o *OrderUseCase) *depositHandler {
	return &depositHandler{
		ch: make(chan *entity.Order),
		o:  o,
		l:  o.l,
	}
}

func (h *depositHandler) handle(wg *sync.WaitGroup) {
	const agent = "depositHandler.handle"
	defer wg.Done()

	for de := range h.ch {
		go func(o *entity.Order) {
			d := o.Deposit
			ex, err := h.o.exs.get(d.Exchange)
			if err != nil {
				d.FailedDesc = err.Error()
				h.o.write(d)
				return
			}

			done := make(chan struct{})
			pCh := make(chan bool)
			go ex.TrackDeposit(o, done, pCh)

			<-done
			// h.l.Debug(agent, fmt.Sprintf("deposit `%d` for order `%d` status changed to  `%s`", d.Id, d.OrderId, d.Status))
			if err := h.o.write(d); err != nil {
				h.l.Error(agent, err.Error())
				pCh <- false
				return
			}
			pCh <- true
			switch d.Status {
			case entity.DepositConfirmed:
				h.confirmedDposit(d.OrderId)
				return
			case entity.DepositFailed:
				h.failedDeposit(d.OrderId)
				return
			}

		}(de)
	}

}

func (h *depositHandler) confirmedDposit(orderId int64) error {
	const op = errors.Op("depositHandler.confirmedDposit")

	o := &entity.Order{Id: orderId}

	err := h.o.read(o)
	if err != nil {
		h.l.Error(string(op), err.Error())
		return err
	}

	// minBc, minQc := h.o.pc.PairMinDeposit(o.BC, o.QC)
	// vf, _ := strconv.ParseFloat(o.Deposit.Volume, 64)

	// switch o.Side {
	// case "buy":
	// 	if vf < minQc {
	// 		o.Deposit.Status = entity.DepositFailed
	// 		o.Deposit.FailedDesc = BR_InsufficientDepositVolume
	// 		o.Status = entity.OSFailed
	// 		o.FailedCode = entity.FCDepositFailed

	// 		if err := h.o.write(o); err != nil {
	// 			h.l.Error(string(op), err.Error())
	// 			return err
	// 		}
	// 		return nil
	// 	}

	// case "sell":
	// 	if vf < minBc {
	// 		o.Deposit.Status = entity.DepositFailed
	// 		o.Deposit.FailedDesc = BR_InsufficientDepositVolume
	// 		o.Status = entity.OSFailed
	// 		o.FailedCode = entity.FCDepositFailed

	// 		if err := h.o.write(o); err != nil {
	// 			h.l.Error(string(op), err.Error())
	// 			return err
	// 		}
	// 		return nil
	// 	}
	// }

	o.Status = entity.OSDepositeConfimred

	if err := h.o.write(o); err != nil {
		h.l.Error(string(op), err.Error())
		return err
	}

	h.o.oh.handle(o)
	return nil

}

func (u *depositHandler) failedDeposit(orderId int64) {
	const agent = "depositHandler.failedDeposit"
	o := &entity.Order{Id: orderId}

	err := u.o.read(o)
	if err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			u.l.Debug(agent, err.Error())
			return
		default:
			u.l.Error(agent, err.Error())
			return
		}
	}

	if o.Deposit.Status == entity.DepositFailed {
		o.Status = entity.OSFailed
		o.FailedCode = entity.FCDepositFailed

		if err := u.o.write(o); err != nil {
			u.l.Error(agent, err.Error())
		}
	}
}
