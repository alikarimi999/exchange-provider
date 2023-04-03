package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"
	"time"
)

func (k *kucoinExchange) TxIdSetted(o *entity.CexOrder) {
	agent := k.agent("TxIdSetted")

	dc, err := k.supportedCoins.get(o.Deposit.String())
	if err != nil {
		o.Deposit.Status = entity.DepositFailed
		o.Deposit.FailedDesc = err.Error()
		return
	}

	k.trackDeposit(o, dc)
	o.UpdatedAt = time.Now().Unix()
	if o.Deposit.Status == entity.DepositFailed {
		o.FailedCode = entity.FCDepositFailed
		o.Status = entity.OFailed
	}

	if err := k.repo.Update(o); err != nil {
		k.l.Error(agent, fmt.Sprintf("( %s ) ( %s )", o, err.Error()))
	}
	k.cache.removeD(o.Deposit.TxId)
	k.cache.proccessedD(o.Deposit.TxId)

	if o.Deposit.Status != entity.DepositConfirmed {
		return
	}

	o.Swaps[0].InAmount = fmt.Sprintf("%v", o.Deposit.Volume)
	id, err := k.swap(o, 0)
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
	}

	k.trackSwap(o, 0)
	switch o.Swaps[0].Status {
	case entity.SwapSucceed:
		if err = k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
		}
		if err := k.withdrawal(o); err != nil {
			k.l.Error(agent, err.Error())

			o.Status = entity.OFailed
			o.FailedCode = entity.FCInternalError
			o.FailedDesc = err.Error()
			if err := k.repo.Update(o); err != nil {
				k.l.Error(agent, err.Error())
			}
			return
		}

		if err := k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
		}
		return

	case entity.SwapFailed:
		o.Status = entity.OFailed
		o.FailedCode = entity.FCExOrdFailed
		if err := k.repo.Update(o); err != nil {
			k.l.Error(agent, err.Error())
		}
		return
	}
}
