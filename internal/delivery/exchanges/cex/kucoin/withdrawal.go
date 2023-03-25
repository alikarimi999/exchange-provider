package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) Withdrawal(o *entity.CexOrder) (string, error) {
	agent := k.agent("Withdrawal")

	c := o.Withdrawal.Token
	opts, err := k.withdrawalOpts(c, o.Withdrawal.Address.Tag)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrBadRequest)
	}

	wc, err := k.supportedCoins.get(c.Symbol, c.Standard)
	if err != nil {
		return "", errors.Wrap(err, errors.ErrBadRequest)
	}

	vol := trim(o.Withdrawal.Volume, wc.WithdrawalPrecision)
	o.Withdrawal.Volume = vol
	// first transfer from trade account to main account
	res, err := k.writeApi.InnerTransferV2(uuid.New().String(), c.Symbol, "trade", "main", vol)
	if err = handleSDKErr(err, res); err != nil {
		return "", err
	}

	k.l.Debug(agent, fmt.Sprintf("%s %s transferred from trade account to main account", vol, c.Symbol))

	// then withdraw from main account
	res, err = k.writeApi.ApplyWithdrawal(c.Symbol, o.Withdrawal.Addr, vol, opts)
	if err = handleSDKErr(err, res); err != nil {
		return "", err
	}

	k.l.Debug(agent, fmt.Sprintf("%s %s withdrawn from main account", vol, c.Symbol))

	w := &kucoin.ApplyWithdrawalResultModel{}
	if err = res.ReadData(w); err != nil {
		return "", errors.Wrap(err, errors.ErrInternal)
	}
	return w.WithdrawalId, nil
}
