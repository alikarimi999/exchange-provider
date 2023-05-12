package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
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
	windowSize time.Duration

	pMux           *sync.RWMutex
	pTicker        *time.Ticker
	proccessedList map[string]struct{ time.Time }
}

func newWithdrawalAggregator(k *kucoinExchange, c *cache) *withdrawalAggregator {
	wa := &withdrawalAggregator{
		k:          k,
		l:          k.l,
		c:          c,
		ticker:     time.NewTicker(time.Second * 15),
		windowSize: time.Hour * 2,

		pMux:           &sync.RWMutex{},
		pTicker:        time.NewTicker(2 * time.Hour),
		proccessedList: make(map[string]struct{ time.Time }),
	}
	return wa
}

func (wa *withdrawalAggregator) run(stopCh chan struct{}) {
	agent := wa.k.agent("withdrawalAggregator.run")

	for {
		select {
		case <-wa.ticker.C:
			wa.aggregateAll(-wa.windowSize, false)
		case <-wa.pTicker.C:
			wa.pMux.Lock()
			for id, s := range wa.proccessedList {
				if time.Now().After(s.Time.Add(2 * time.Hour)) {
					delete(wa.proccessedList, id)
				}
			}
			wa.pMux.Unlock()
		case <-stopCh:
			wa.l.Debug(agent, "stopped")
			return
		}

	}
}

func (wa *withdrawalAggregator) aggregateAll(windSize time.Duration,
	withPending bool) ([]*dto.Withdrawal, error) {
	agent := wa.k.agent("aggregateAll")
	t := time.Now()

	var (
		ws  []*dto.Withdrawal
		err error
	)
	if withPending {
		ws, err = wa.aggregate("", t.Add(windSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
			return nil, err
		}
	} else {
		ws, err = wa.aggregate("SUCCESS", t.Add(windSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
			return nil, err
		}

		ws1, err := wa.aggregate("FAILURE", t.Add(-wa.windowSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
		}

		ws = append(ws, ws1...)
	}
	for _, wd := range ws {
		if !wa.isProccessed(wd.Id) && (wd.Status == "SUCCESS" || wd.Status == "FAILURE") {
			os, err := wa.k.repo.GetWithFilter("order.withdrawal.id", wd.Id)
			if err != nil {
				if errors.ErrorCode(err) == errors.ErrNotFound {
					wa.addToProccessedList(wd.Id)
				}
				continue
			}

			co := os[0].(*types.Order)
			if co.Status == types.OWithdrawalTracking {
				switch wd.Status {
				case "SUCCESS":
					co.Withdrawal.KucoinFee, _ = strconv.ParseFloat(wd.Fee, 64)
					co.Withdrawal.TxId = wd.FixTxId()
					co.Status = types.OWithdrawalConfirmed
				case "FAILURE":
					co.Status = types.OWithdrawalFailed
					co.FailedDesc = "failed by kucoin"
				}
				if err := wa.k.repo.Update(os[0]); err != nil {
					continue
				}
			}
			wa.addToProccessedList(wd.Id)
		}
	}
	return ws, nil
}

func (wa *withdrawalAggregator) aggregate(status string, start, end time.Time) ([]*dto.Withdrawal, error) {
	params := make(map[string]string)
	params["startAt"] = strconv.FormatInt(start.UnixMilli(), 10)
	params["endAt"] = strconv.FormatInt(end.UnixMilli(), 10)
	if status != "" {
		params["status"] = status
	}

	paginate := &kucoin.PaginationParam{
		CurrentPage: 1,
		PageSize:    100,
	}
	for {

		res, err := wa.k.readApi.Withdrawals(params, paginate)
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
	_, ok := w.proccessedList[wId]
	return ok
}

func (w *withdrawalAggregator) addToProccessedList(wId string) {
	w.pMux.Lock()
	defer w.pMux.Unlock()
	w.proccessedList[wId] = struct{ time.Time }{Time: time.Now()}
}
