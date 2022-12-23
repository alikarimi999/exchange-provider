package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) Withdrawal(o *entity.Order) (string, error) {
	op := errors.Op(fmt.Sprintf("%s.Withdrawal", k.Id()))

	c := o.Withdrawal.Token

	opts, err := k.withdrawalOpts(c, o.Withdrawal.Tag)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	wc, err := k.supportedCoins.get(c.TokenId, c.ChainId)
	if err != nil {
		return "", errors.Wrap(err, op, errors.ErrBadRequest)
	}

	vol := trim(o.Withdrawal.Volume, wc.WithdrawalPrecision)
	o.Withdrawal.Volume = vol
	// first transfer from trade account to main account
	res, err := k.api.InnerTransferV2(uuid.New().String(), c.TokenId, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s transferred from trade account to main account", vol, c.TokenId))

	// then withdraw from main account
	res, err = k.api.ApplyWithdrawal(c.TokenId, o.Withdrawal.Addr, vol, opts)
	if err = handleSDKErr(err, res); err != nil {
		return "", errors.Wrap(err, op)
	}

	k.l.Debug(string(op), fmt.Sprintf("%s %s withdrawn from main account", vol, c.TokenId))

	w := &kucoin.ApplyWithdrawalResultModel{}
	if err = res.ReadData(w); err != nil {
		return "", errors.Wrap(err, op, errors.ErrInternal)
	}
	return w.WithdrawalId, nil
}
