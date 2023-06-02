package binance

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/pkg/logger"
	"strconv"
	"sync"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/adshao/go-binance/v2"
)

const (
	processing = iota + 4
	failure
	completed
)

type withdrawalAggregator struct {
	ex         *exchange
	l          logger.Logger
	ticker     *time.Ticker
	windowSize time.Duration

	pMux           *sync.RWMutex
	pTicker        *time.Ticker
	proccessedList map[string]struct{ time.Time }
}

func newWithdrawalAggregator(ex *exchange) *withdrawalAggregator {
	wa := &withdrawalAggregator{
		ex:         ex,
		l:          ex.l,
		ticker:     time.NewTicker(time.Second * 30),
		windowSize: time.Hour * 2,

		pMux:           &sync.RWMutex{},
		pTicker:        time.NewTicker(2 * time.Hour),
		proccessedList: make(map[string]struct{ time.Time }),
	}
	return wa
}

func (wa *withdrawalAggregator) run(stopCh <-chan struct{}) {
	agent := wa.ex.agent("withdrawalAggregator.run")
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
	withPending bool) ([]*binance.Withdraw, error) {
	agent := wa.ex.agent("aggregateAll")
	t := time.Now()

	var (
		ws  []*binance.Withdraw
		err error
	)
	if withPending {
		ws, err = wa.aggregate(processing, t.Add(windSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
			return nil, err
		}
	} else {
		ws, err = wa.aggregate(completed, t.Add(windSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
			return nil, err
		}

		ws1, err := wa.aggregate(failure, t.Add(-wa.windowSize), t)
		if err != nil {
			wa.l.Debug(agent, err.Error())
		}

		ws = append(ws, ws1...)
	}
	for _, wd := range ws {
		if !wa.isProccessed(wd.ID) && (wd.Status == completed || wd.Status == failure) {
			os, err := wa.ex.repo.GetWithFilter("order.withdrawal.id", wd.ID)
			if err != nil {
				if errors.ErrorCode(err) == errors.ErrNotFound {
					wa.addToProccessedList(wd.ID)
				}
				continue
			}
			co := os[0].(*types.Order)
			if co.Status == types.OWithdrawalTracking {
				switch wd.Status {
				case completed:
					co.Withdrawal.Amount, _ = strconv.ParseFloat(wd.Amount, 64)
					co.Withdrawal.BinanceFee, _ = strconv.ParseFloat(wd.TransactionFee, 64)
					co.Withdrawal.TxId = wd.TxID
					co.Status = types.OWithdrawalConfirmed
				case failure:
					co.Status = types.OWithdrawalFailed
					co.FailedDesc = "failed by binance"
				}
				if err := wa.ex.repo.Update(os[0]); err != nil {
					wa.l.Debug(agent, err.Error())
					continue
				}
			}
			wa.addToProccessedList(wd.ID)
		}
	}
	return ws, nil
}

func (wa *withdrawalAggregator) aggregate(status int,
	start, end time.Time) ([]*binance.Withdraw, error) {
	return wa.ex.c.NewListWithdrawsService().Status(status).StartTime(start.UnixMilli()).
		EndTime(end.UnixMilli()).Do(context.Background())

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
