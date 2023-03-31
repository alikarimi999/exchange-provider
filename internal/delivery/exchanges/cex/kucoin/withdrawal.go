package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) withdrawal(o *entity.CexOrder) error {
	k.applySpreadAndFee(o, o.Routes[0])
	c := o.Withdrawal.Token
	wc, err := k.supportedCoins.get(c.String())
	if err != nil {
		return errors.Wrap(err, errors.ErrBadRequest)
	}
	opts := make(map[string]string)
	opts["chain"] = wc.Chain
	opts["memo"] = o.Withdrawal.Tag

	vol := trim(o.Withdrawal.Volume, wc.WithdrawalPrecision)
	o.Withdrawal.Volume = vol
	// first transfer from trade account to main account
	res, err := k.writeApi.InnerTransferV2(uuid.New().String(), wc.Currency, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return err
	}

	// then withdraw from main account
	res, err = k.writeApi.ApplyWithdrawal(wc.Currency, o.Withdrawal.Addr, vol, opts)
	if err = handleSDKErr(err, res); err != nil {
		return err
	}

	w := &kucoin.ApplyWithdrawalResultModel{}
	if err = res.ReadData(w); err != nil {
		return errors.Wrap(err, errors.ErrInternal)
	}

	o.Withdrawal.TxId = w.WithdrawalId
	o.Withdrawal.Status = entity.WithdrawalPending
	o.Status = entity.OWaitForWithdrawalConfirm
	return nil
}
