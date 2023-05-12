package kucoin

import (
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
)

type depositAggregator struct {
	k     *kucoinExchange
	l     logger.Logger
	t     *time.Ticker
	c     *cache
	wSize time.Duration
}

func newDepositAggregator(k *kucoinExchange, c *cache) *depositAggregator {
	da := &depositAggregator{
		k:     k,
		l:     k.l,
		t:     time.NewTicker(15 * time.Second),
		c:     c,
		wSize: time.Hour * 2,
	}
	return da
}

func (a *depositAggregator) run(stopCh chan struct{}) {
	agent := fmt.Sprintf("%s.depositAggregator.run", a.k.NID())

	for {
		select {
		case <-a.t.C:
			a.aggregateAll(-a.wSize)
		case <-stopCh:
			a.l.Debug(agent, "stopped")
			return
		}
	}

}

func (a *depositAggregator) aggregateAll(windSize time.Duration) error {
	agent := a.k.agent("aggregateAll")
	s := time.Now().Add(windSize)
	e := time.Now()
	ds, err := a.aggregate("SUCCESS", s, e)
	if err != nil {
		a.l.Debug(agent, err.Error())
		return err
	}

	dsf, err := a.aggregate("FAILURE", s, e)
	if err != nil {
		a.l.Debug(agent, err.Error())
		return err
	}

	ds = append(ds, dsf...)
	for _, d := range ds {
		if !a.c.existsOrProccessedD(d.TxId) {
			a.c.saveD(d)
			continue
		}
	}
	return nil
}

func (a *depositAggregator) aggregate(status string, start, end time.Time) ([]*depositRecord, error) {
	op := errors.Op(fmt.Sprintf("%s.depositAggregator.aggregate", a.k.NID()))

	ps := make(map[string]string)
	ps["startAt"] = strconv.FormatInt(start.UnixMilli(), 10)
	ps["endAt"] = strconv.FormatInt(end.UnixMilli(), 10)
	ps["status"] = status

	paginate := &kucoin.PaginationParam{
		CurrentPage: 1,
		PageSize:    100,
	}
	for {

		res, err := a.k.readApi.Deposits(ps, paginate)
		if err = handleSDKErr(err, res); err != nil {
			return nil, errors.Wrap(err, op)
		}

		ds := []*kucoin.DepositModel{}
		pa, err := res.ReadPaginationData(&ds)
		if err != nil {
			return nil, err
		}

		rds := []*depositRecord{}
		for _, d := range ds {
			if !d.IsInner && d.WalletTxId != "" {
				rds = append(rds, mapDeposit(d))
			}
		}

		if pa.CurrentPage < pa.TotalPage {
			paginate.CurrentPage = pa.CurrentPage + 1
			continue
		}

		return rds, nil
	}

}

func mapDeposit(d *kucoin.DepositModel) *depositRecord {
	return &depositRecord{
		TxId:         strings.Split(d.WalletTxId, "@")[0],
		Currency:     d.Currency,
		Volume:       d.Amount,
		Status:       d.Status,
		DownloadedAt: time.Now(),
	}
}
