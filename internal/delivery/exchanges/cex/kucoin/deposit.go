package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/try"
	"fmt"
	"strconv"
	"time"
)

func (k *kucoinExchange) trackDeposit(o *types.Order, dc *Token) {
	t := dc.BlockTime * time.Duration(dc.ConfirmBlocks)
	err := try.Do(20, func(attempt uint64) (bool, error) {
		fmt.Println(attempt)
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
