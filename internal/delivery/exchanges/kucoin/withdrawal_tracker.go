package kucoin

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
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
	txId := f.w.TxId
	wd, ok := t.c.getWithdrawal(txId)
	if !ok {
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
		t.c.delWithdrawal(txId)
		t.c.proccessedW(txId)
	}
}
