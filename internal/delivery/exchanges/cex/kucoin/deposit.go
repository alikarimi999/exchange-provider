package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/try"
	"strconv"
	"time"
)

func (k *exchange) trackDeposit(o *types.Order, dc *Token) {
	t := dc.BlockTime * time.Duration(dc.ConfirmBlocks)
	if t < time.Minute {
		time.Sleep(time.Minute)
		t *= 2
	}
	err := try.Do(100, func(attempt uint64) (bool, error) {
		d, ok := k.cache.getD(o.Deposit.TxId)
		if ok {
			if err := d.MatchCurrency(dc); err != nil {
				o.Status = types.ODepositFailed
				o.FailedDesc = err.Error()
				return false, nil
			}
			if d.Status == "SUCCESS" {
				o.Status = types.ODepositeConfimred
			} else {
				o.Status = types.ODepositFailed
				o.FailedDesc = "order failed in kucoin"
			}
			vol, _ := strconv.ParseFloat(d.Volume, 64)
			o.Deposit.Amount = vol
			return false, nil
		}

		time.Sleep(t / 2)
		return true, errors.Wrap(errors.ErrNotFound)
	})

	if err != nil {
		o.Status = types.ODepositFailed
		o.FailedDesc = err.Error()
	}
}
