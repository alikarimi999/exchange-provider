package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type wtFeed struct {
	w            *entity.Withdrawal
	done         chan<- struct{}
	proccessedCh <-chan bool
}

type withdrawalTracker struct {
	k *kucoinExchange
	l logger.Logger
	c *cache
}

func newWithdrawalTracker(k *kucoinExchange, c *cache) *withdrawalTracker {
	return &withdrawalTracker{
		k: k,
		l: k.l,
		c: c,
	}
}

func (t *withdrawalTracker) track(f *wtFeed) {
	op := errors.Op(fmt.Sprintf("%s.withdrawalTracker.track", t.k.Id()))
	txId := f.w.TxId
	wd, err := t.c.getWithdrawal(txId)
	if err != nil {
		if err != redis.Nil {
			t.l.Error(string(op), err.Error())
		}
		f.w.Status = entity.WithdrawalPending
		f.done <- struct{}{}
		<-f.proccessedCh
		return
	}
	switch wd.Status {
	case "SUCCESS":
		f.w.Status = entity.WithdrawalSucceed
		f.w.Fee = wd.Fee
		f.w.TxId = wd.FixTxId() + "-" + txId
		f.w.FeeCurrency = f.w.Token.String()
	case "FAILURE":
		f.w.Status = entity.WithdrawalFailed
		f.w.FailedDesc = "failed by exchange"
	}
	f.done <- struct{}{}

	if <-f.proccessedCh {
		if err := t.c.delWithdrawal(txId); err != nil {
			t.l.Error(string(op), errors.Wrap(err, op).Error())
		}
		if err := t.c.proccessedWithdrawal(txId); err != nil {
			t.l.Error(string(op), errors.Wrap(err, op).Error())
		}
	}
}
