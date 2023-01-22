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
	return &depositAggregator{
		k:     k,
		l:     k.l,
		t:     time.NewTicker(time.Second * 30),
		c:     c,
		wSize: time.Hour * 2,
	}
}

func (a *depositAggregator) run(stopCh chan struct{}) {
	agent := fmt.Sprintf("%s.depositAggregator.run", a.k.Id())
	for {
		select {
		case <-a.t.C:

			s := time.Now().Add(-a.wSize)
			e := time.Now()
			ds, err := a.aggregate("SUCCESS", s, e)
			if err != nil {
				a.l.Error(agent, err.Error())
				continue
			}

			dsf, err := a.aggregate("FAILURE", s, e)
			if err != nil {
				a.l.Error(agent, err.Error())

			}
			ds = append(ds, dsf...)

			for _, d := range ds {
				exist, err := a.c.ExistD(d.TxId)
				if err != nil {
					a.l.Error(agent, err.Error())
					continue
				}
				if !exist {
					a.c.SaveD(d)
					continue
				}
			}

		case <-stopCh:
			a.l.Debug(agent, "stopped")
			return
		}
	}

}

func (a *depositAggregator) aggregate(status string, start, end time.Time) ([]*depositeRecord, error) {
	op := errors.Op(fmt.Sprintf("%s.depositAggregator.aggregate", a.k.Id()))

	ps := make(map[string]string)
	ps["startAt"] = strconv.FormatInt(start.UnixMilli(), 10)
	ps["endAt"] = strconv.FormatInt(end.UnixMilli(), 10)
	ps["status"] = status

	paginate := &kucoin.PaginationParam{
		CurrentPage: 1,
		PageSize:    100,
	}
	for {

		res, err := a.k.api.Deposits(ps, paginate)
		if err = handleSDKErr(err, res); err != nil {
			return nil, errors.Wrap(err, op)
		}

		ds := []*kucoin.DepositModel{}
		pa, err := res.ReadPaginationData(&ds)
		if err != nil {
			return nil, err
		}

		rds := []*depositeRecord{}
		for _, d := range ds {
			if !d.IsInner {
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

func mapDeposit(d *kucoin.DepositModel) *depositeRecord {
	return &depositeRecord{
		TxId:     strings.Split(d.WalletTxId, "@")[0],
		Currency: d.Currency,
		Volume:   d.Amount,
		Status:   d.Status,
	}
}
