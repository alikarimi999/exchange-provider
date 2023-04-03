package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/try"
	"fmt"
	"strconv"
	"time"
)

func (k *kucoinExchange) SetDepositddress(o *entity.CexOrder) error {
	kc, err := k.supportedCoins.get(o.Deposit.String())
	if err != nil {
		return err
	}

	o.Deposit.Address.Addr = kc.DepositAddress
	o.Deposit.Address.Tag = kc.DepositTag
	return nil
}

func (k *kucoinExchange) trackDeposit(o *entity.CexOrder, dc *Token) {
	t := dc.BlockTime * time.Duration(dc.ConfirmBlocks)
	if t < time.Minute {
		time.Sleep(time.Minute)
	}
	err := try.Do(50, func(attempt uint64) (bool, error) {
		d, ok := k.cache.getD(o.Deposit.TxId)
		if ok {
			if !d.MatchCurrency(dc) {
				o.Deposit.Status = entity.DepositFailed
				o.Deposit.FailedDesc = fmt.Sprintf("currency mismatch, user: `%s`, exchange: `%s` ",
					o.Deposit.Symbol, d.Currency)
				return false, nil
			}
			o.Deposit.Status = entity.DepositConfirmed
			vol, _ := strconv.ParseFloat(d.Volume, 64)
			o.Deposit.Volume = vol
			return false, nil

		}

		if t < time.Minute {
			time.Sleep(time.Minute)
		} else {
			time.Sleep(t / 8)
		}
		return true, errors.Wrap(errors.ErrNotFound)
	})

	if err != nil {
		o.Deposit.Status = entity.DepositFailed
		o.Deposit.FailedDesc = err.Error()
	}
}
