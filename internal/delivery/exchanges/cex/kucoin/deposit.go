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
	kc, err := k.supportedCoins.get(o.Deposit.Symbol, o.Deposit.Standard)
	if err != nil {
		return err
	}

	o.Deposit.Address.Addr = kc.Address
	o.Deposit.Address.Tag = kc.Tag
	return nil
}

func (k *kucoinExchange) trackDeposit(o *entity.CexOrder, t *Token) {
	err := try.Do(15, func(attempt uint64) (bool, error) {
		d, ok := k.cache.getD(o.Deposit.TxId)
		if ok {
			if !d.MatchCurrency(o.Deposit) {
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

		t := (t.BlockTime + (5 * time.Second)) * time.Duration(t.ConfirmBlocks)
		time.Sleep(t / 2)
		return true, errors.Wrap(errors.ErrNotFound)
	})

	if err != nil {
		o.Deposit.Status = entity.DepositFailed
		o.Deposit.FailedDesc = err.Error()
	}
}
