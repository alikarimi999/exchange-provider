package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"
)

func (k *kucoinExchange) TxIdSetted(o *entity.CexOrder) {
	agent := k.agent("TxIdSetted")
	done := make(chan struct{})
	pCh := make(chan bool)
	go k.trackDeposit(o, done, pCh)

	<-done
	if o.Deposit.Status == entity.DepositFailed {
		o.FailedCode = entity.FCDepositFailed
		o.Status = entity.OFailed
	}

	if err := k.repo.Update(o); err != nil {
		k.l.Error(agent, err.Error())
		pCh <- false
		return
	}
	pCh <- true

	o.Swaps[0].InAmount = fmt.Sprintf("%v", o.Deposit.Volume)

	id, err := k.Swap(o, 0)
	if err != nil {
		o.Status = entity.OFailed
		o.FailedCode = entity.FCExOrdFailed
		o.FailedDesc = err.Error()
		k.l.Error(agent, err.Error())
		if err := k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())

		}
		return
	}

	o.Swaps[0].TxId = id
	o.Status = entity.OWaitForSwapConfirm
	if err = k.repo.Update(o); err != nil {
		k.l.Error(agent, err.Error())
		return
	}
	go k.TrackSwap(o, 0, done, pCh)
	<-done

	switch o.Swaps[0].Status {
	case entity.SwapSucceed:
		if err = k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
			pCh <- false
			return
		}
		pCh <- true

		k.applySpreadAndFee(o, o.Routes[0])
		id, err = k.Withdrawal(o)
		if err != nil {
			o.Status = entity.OFailed
			o.FailedCode = entity.FCInternalError
			o.FailedDesc = err.Error()

			k.l.Error(agent, err.Error())

			if err := k.repo.Update(o); err != nil {
				k.l.Error(agent, err.Error())
			}
			return

		}

		o.Withdrawal.TxId = id
		o.Withdrawal.Status = entity.WithdrawalPending
		o.Status = entity.OWaitForWithdrawalConfirm

		if err := k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
			return
		}
		return

	case entity.SwapFailed:
		o.Status = entity.OFailed
		o.FailedCode = entity.FCExOrdFailed
		if err := k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
			pCh <- false
			return
		}
		pCh <- true
		return
	}
}
