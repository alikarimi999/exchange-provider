package kucoin

import (
	"fmt"
	"order_service/internal/delivery/exchanges/kucoin/dto"
	"order_service/pkg/logger"
	"strconv"
	"sync"
	"time"

	"order_service/pkg/errors"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/go-redis/redis/v9"
)

type withdrawalAggregator struct {
	api        *kucoin.ApiService
	l          logger.Logger
	c          *withdrawalCache
	ticker     *time.Ticker
	params     map[string]string
	windowSize time.Duration
}

func newWithdrawalAggregator(api *kucoin.ApiService, l logger.Logger, r *redis.Client) *withdrawalAggregator {
	return &withdrawalAggregator{
		api:        api,
		l:          l,
		c:          newWithdrawalCache(r, l),
		params:     make(map[string]string),
		ticker:     time.NewTicker(time.Minute * 10),
		windowSize: time.Hour * 1,
	}
}

func (wa *withdrawalAggregator) run(wg *sync.WaitGroup) {
	const op = errors.Op("Kucoin.WithdrawalAggregator.run")
	defer wg.Done()
start:
	for {
		select {
		case t := <-wa.ticker.C:
			ws := []*dto.Withdrawal{}
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

			ws = append(wss, wsf...)

			if l := len(ws); l > 0 {
				wa.l.Debug(string(op), fmt.Sprintf("aggregated '%d' withdrawals which occured from ( %s ) to ( %s ) ",
					l, t.Add(-wa.windowSize).String(), t.String()))
			}

			for _, w := range ws {
				p, err := wa.c.isAddable(w.Id)
				if err != nil {
					wa.l.Error(string(op), errors.Wrap(err, op).Error())
				}
				if p {
					wa.l.Debug(string(op), fmt.Sprintf("withdrawal '%s' cached before", w.Id))
					continue
				}
				if err := wa.c.recordWithdrawal(w); err != nil {
					wa.l.Error(string(op), errors.Wrap(err, op, w.String()).Error())
					continue start
				}

			}

		}

	}
}

func (wa *withdrawalAggregator) aggregate(status string, start, end time.Time) ([]*dto.Withdrawal, error) {

	wa.params["startAt"] = strconv.FormatInt(start.UnixMilli(), 10)
	wa.params["endAt"] = strconv.FormatInt(end.UnixMilli(), 10)
	wa.params["status"] = status

	paginate := &kucoin.PaginationParam{
		CurrentPage: 1,
		PageSize:    100,
	}
	for {

		resp, err := wa.api.Withdrawals(wa.params, paginate)
		if err != nil || resp.Code != "200000" {
			return nil, err
		}

		withdrawals := []*dto.Withdrawal{}
		pa, err := resp.ReadPaginationData(&withdrawals)
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
