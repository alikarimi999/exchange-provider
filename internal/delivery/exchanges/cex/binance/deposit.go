package binance

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/try"
	"fmt"
	"strconv"
	"time"
)

const (
	depositPending  = 0
	depositSuccess  = 1
	depositCredited = 6
)

func (k *exchange) trackDeposit(o *types.Order, dc *Token) {
	t := dc.BlockTime * time.Duration(dc.UnLockConfirm)
	if t < time.Minute {
		time.Sleep(time.Minute)
		t *= 2
	}
	err := try.Do(100, func(attempt uint64) (bool, error) {
		ds, err := k.c.NewListDepositsService().TxID(o.Deposit.TxId).Do(context.Background())
		if err == nil {
			if len(ds) == 0 {
				return true, fmt.Errorf("deposit txId not found in binance deposit list")
			}

			if ds[0].Status == depositPending {
				time.Sleep(t / 2)
				return true, fmt.Errorf("binance deposit status is '%d'", ds[0].Status)
			}

			if ds[0].Coin != dc.Coin {
				o.Status = types.ODepositFailed
				o.FailedDesc = fmt.Sprintf("currency mismatch,'%s':'%s'",
					dc.Coin, ds[0].Coin)
				return false, nil
			}
			if ds[0].Status == depositSuccess || ds[0].Status == depositCredited {
				o.Status = types.ODepositeConfimred
			} else {
				return false, fmt.Errorf("binance deposit status is '%d'", ds[0].Status)
			}

			vol, _ := strconv.ParseFloat(ds[0].Amount, 64)
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
