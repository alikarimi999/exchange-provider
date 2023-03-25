package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"strconv"
	"sync"
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

	pMux    *sync.RWMutex
	pTicker *time.Ticker
	list    map[string]struct{ time.Time }
}

func newWithdrawalAggregator(k *kucoinExchange, c *cache) *withdrawalAggregator {
	wa := &withdrawalAggregator{
		k:          k,
		l:          k.l,
		c:          c,
		params:     make(map[string]string),
		ticker:     time.NewTicker(time.Minute * 2),
		windowSize: time.Hour * 1,

		pMux:    &sync.RWMutex{},
		pTicker: time.NewTicker(2 * time.Hour),
		list:    make(map[string]struct{ time.Time }),
	}
	go wa.run(k.stopCh)
	return wa
}

func (wa *withdrawalAggregator) run(stopCh chan struct{}) {
	agent := wa.k.agent("withdrawalAggregator.run")

	for {
		select {
		case t := <-wa.ticker.C:
			wss, err := wa.aggregate("SUCCESS", t.Add(-wa.windowSize), t)
			if err != nil {
				wa.l.Error(agent, err.Error())
				continue
			}
			wsf, err := wa.aggregate("FAILURE", t.Add(-wa.windowSize), t)
			if err != nil {
				wa.l.Error(agent, err.Error())
				continue

			}
			wss = append(wss, wsf...)

			for _, wd := range wss {
				if !wa.isProccessed(wd.Id) {
					o, err := wa.k.repo.GetWithFilter("orders.withdrawal.id", wd.Id)
					if err != nil {
						if errors.ErrorCode(err) == errors.ErrNotFound {
							wa.addToList(wd.Id)
						}
						continue
					}

					co := o.(*entity.CexOrder)
					switch wd.Status {
					case "SUCCESS":
						co.Withdrawal.Status = entity.WithdrawalSucceed
						co.Withdrawal.Fee = wd.Fee
						co.Withdrawal.TxId = wd.FixTxId()
						co.Withdrawal.FeeCurrency = co.Withdrawal.Token.String()
						co.Status = entity.OSucceeded
						co.UpdatedAt = time.Now().Unix()
					case "FAILURE":
						co.Withdrawal.Status = entity.WithdrawalFailed
						co.Withdrawal.FailedDesc = "failed by exchange"
						co.Status = entity.OFailed
						co.UpdatedAt = time.Now().Unix()
					}
					if err := wa.k.repo.Update(o); err != nil {
						wa.l.Error(agent, err.Error())
						continue
					}
					wa.addToList(wd.Id)
				}
			}

		case <-wa.pTicker.C:
			wa.pMux.Lock()
			for id, s := range wa.list {
				if time.Since(s.Time) >= time.Duration(2*time.Hour) {
					delete(wa.list, id)
				}
			}
			wa.pMux.Unlock()
		case <-stopCh:
			wa.l.Debug(agent, "stopped")
			return
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

		res, err := wa.k.readApi.Withdrawals(wa.params, paginate)
		if err = handleSDKErr(err, res); err != nil {
			return nil, err
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

func (w *withdrawalAggregator) isProccessed(wId string) bool {
	w.pMux.RLock()
	defer w.pMux.RUnlock()
	_, ok := w.list[wId]
	return ok
}

func (w *withdrawalAggregator) addToList(wId string) {
	w.pMux.Lock()
	defer w.pMux.Unlock()
	w.list[wId] = struct{ time.Time }{Time: time.Now()}
}
