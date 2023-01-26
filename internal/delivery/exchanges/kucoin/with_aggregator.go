package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/pkg/logger"
	"fmt"
	"strconv"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
)

type withdrawalAggregator struct {
	k          *kucoinExchange
	l          logger.Logger
	c          *cache
	ticker     *time.Ticker
	params     map[string]string
	windowSize time.Duration
}

func newWithdrawalAggregator(k *kucoinExchange, c *cache) *withdrawalAggregator {
	wa := &withdrawalAggregator{
		k:          k,
		l:          k.l,
		c:          c,
		params:     make(map[string]string),
		ticker:     time.NewTicker(time.Minute * 2),
		windowSize: time.Hour * 1,
	}
	go wa.run(k.stopCh)
	return wa
}

func (wa *withdrawalAggregator) run(stopCh chan struct{}) {
	op := errors.Op(fmt.Sprintf("%s.withdrawalAggregator.run", wa.k.Id()))
	wa.l.Debug(string(op), "started")

	for {
		select {
		case t := <-wa.ticker.C:
			wss, err := wa.aggregate("SUCCESS", t.Add(-wa.windowSize), t)
			if err != nil {
				wa.l.Error(string(op), errors.Wrap(err, op).Error())
				continue
			}
			wsf, err := wa.aggregate("FAILURE", t.Add(-wa.windowSize), t)
			if err != nil {
				wa.l.Error(string(op), errors.Wrap(err, op).Error())
				continue

			}

			wss = append(wss, wsf...)

			for _, w := range wss {
				exists := wa.c.existsOrProccessedW(w.Id)
				if !exists {
					wa.c.recordWithdrawal(w)
				}
			}

		case <-stopCh:
			wa.l.Debug(string(op), "stopped")
			return
		}

	}
}

func (wa *withdrawalAggregator) aggregate(status string, start, end time.Time) ([]*dto.Withdrawal, error) {
	op := errors.Op(fmt.Sprintf("%s.withdrawalAggregator.aggregate", wa.k.Id()))
	wa.params["startAt"] = strconv.FormatInt(start.UnixMilli(), 10)
	wa.params["endAt"] = strconv.FormatInt(end.UnixMilli(), 10)
	wa.params["status"] = status

	paginate := &kucoin.PaginationParam{
		CurrentPage: 1,
		PageSize:    100,
	}
	for {

		res, err := wa.k.api.Withdrawals(wa.params, paginate)
		if err = handleSDKErr(err, res); err != nil {
			return nil, errors.Wrap(err, op)
		}

		withdrawals := []*dto.Withdrawal{}
		pa, err := res.ReadPaginationData(&withdrawals)
		if err != nil {
			return nil, err
		}

		ws := []*dto.Withdrawal{}
		for _, wd := range withdrawals {
			if !wd.IsInner {
				ws = append(ws, wd)

			}
		}

		if pa.CurrentPage < pa.TotalPage {
			paginate.CurrentPage = pa.CurrentPage + 1
			continue
		}

		return ws, nil

	}

}
