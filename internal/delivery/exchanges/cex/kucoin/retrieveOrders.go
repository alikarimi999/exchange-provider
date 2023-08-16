package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	maxGoroutines = 10
)

func (ex *exchange) retreiveOrders(lastUpdate time.Time) error {
	agent := ex.agent("retreiveOrders")
	ex.l.Debug(agent, "processing unfinished orders...")

	duration := time.Until(lastUpdate.Add(-time.Hour))
	if err := ex.da.aggregateAll(duration); err != nil {
		return err
	}

	ws, err := ex.wa.aggregateAll(duration, true)
	if err != nil {
		return err
	}

	pairs := make(map[string]*entity.Pair)

	os0, err := ex.getOrders(types.ODepositTxIdSetted)
	if err != nil {
		return err
	}

	tokensEfa := make(map[string]float64)
	for _, o := range os0 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	os1, err := ex.getOrders(types.ODepositeConfimred)
	if err != nil {
		return err
	}

	for _, o := range os1 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	os2, err := ex.getOrders(types.OFirstSwapTracking)
	if err != nil {
		return err
	}

	for _, o := range os2 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	os3, err := ex.getOrders(types.OSecondSwapTracking)
	if err != nil {
		return err
	}
	for _, o := range os3 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	os4, err := ex.getOrders(types.OFirstSwapCompleted)
	if err != nil {
		return err
	}

	for _, o := range os4 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	os5, err := ex.getOrders(types.OSecondSwapCompleted)
	if err != nil {
		return err
	}
	for _, o := range os5 {
		to := o.(*types.Order)
		p, err := ex.pairs.Get(ex.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		efa, err := ex.tokensEfa(out, p, tokensEfa)
		if err != nil {
			return err
		}
		to.EstimateFeeAmount = efa
		pairs[o.ID().String()] = p
	}

	for i := 0; i < len(os4); i++ {
		if len(os4[i].(*types.Order).Swaps) == 1 {
			os5 = append(os5, os4[i])
			os4 = append(os4[:i], os4[i+1:]...)
			i--
		}
	}
	ex.l.Debug(agent, fmt.Sprintf("there are %d unfinished orders to process",
		len(os0)+len(os1)+len(os2)+len(os3)+len(os4)+len(os5)))

	ex.orderCheckAndFixWithdraw(os5, ws, pairs)

	os := []entity.Order{}
	os = append(os, os1...)
	os = append(os, os2...)
	os = append(os, os3...)
	os = append(os, os4...)

	for _, o := range os {
		ex.orderCheckAndFixSwaps(o.(*types.Order), pairs[o.ID().String()])
	}

	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, maxGoroutines)
	for _, o := range os0 {
		ch <- struct{}{}
		wg.Add(1)
		go func(to *types.Order, p *entity.Pair) {
			defer func() {
				wg.Done()
				<-ch
			}()
			ex.handleOrder(to, p)
		}(o.(*types.Order), pairs[o.ID().String()])
	}
	wg.Wait()
	return nil
}

// for orders that sends withdrawal requests to kucoin but
// the server restarts before update on the database
// Status: "OFirstSwapCompleted" and "OSecondSwapCompleted"
func (k *exchange) orderCheckAndFixWithdraw(os []entity.Order, ws []*dto.Withdrawal,
	ps map[string]*entity.Pair) {
	for _, o := range os {
		co := o.(*types.Order)
		id := co.ID().String()
		for _, w := range ws {
			if id == w.Remark {
				switch w.Status {
				case "SUCCESS":
					co.Withdrawal.KucoinFee, _ = strconv.ParseFloat(w.Fee, 64)
					co.Withdrawal.TxId = w.FixTxId()
					co.Status = types.OWithdrawalConfirmed
				case "FAILURE":
					co.Status = types.OWithdrawalFailed
					co.FailedDesc = "failed by kucoin"
				default:
					co.Withdrawal.Id = w.Id
					co.Status = types.OWithdrawalTracking
				}
				if err := k.repo.Update(o); err != nil {
					k.l.Debug(k.agent("orderCheckAndFixWithdraw"), err.Error())
					continue
				}
				if w.Status == "SUCCESS" || w.Status == "FAILURE" {
					k.wa.addToProccessedList(w.Id)
				}
			}
		}
	}
	for _, o := range os {
		co := o.(*types.Order)
		if (len(co.Swaps) == 1 && co.Status == types.OFirstSwapCompleted) ||
			co.Status == types.OSecondSwapCompleted {

			_, _, wc := getBcQcWcFeeRate(co, ps[o.ID().String()],
				len(co.Swaps)-1)
			k.withdrawal(co, wc, ps[o.ID().String()], true)
			k.repo.Update(o)
		}
	}
}

func (k *exchange) orderCheckAndFixSwaps(o *types.Order, p *entity.Pair) {
	var (
		bc, qc *Token
	)
	for i := range o.Swaps {
		bc, qc, _ = getBcQcWcFeeRate(o, p, i)
		if (i == 0 && o.Status == types.ODepositeConfimred) ||
			(i == 0 && o.Status == types.OFirstSwapTracking) ||
			(i == 1 && o.Status == types.OFirstSwapCompleted) ||
			(i == 1 && o.Status == types.OSecondSwapTracking) {
			if err := k.trackSwap(o, bc, qc, i); err != nil {
				if errors.ErrorCode(err) == errors.ErrNotFound {
					// createOrder request didn't send before sever down
					k.handleOrder(o, p)
					return
				}
				if i == 0 {
					o.Status = types.OFirstSwapFailed
				} else {
					o.Status = types.OSecondSwapFailed
				}
				o.FailedDesc = err.Error()
				k.repo.Update(o)
				return
			}
			if i == 0 {
				o.Status = types.OFirstSwapCompleted
			} else {
				o.Status = types.OSecondSwapCompleted
			}
			k.repo.Update(o)
			k.handleOrder(o, p)
			return
		}
	}
}

func (k *exchange) getOrders(status string) ([]entity.Order, error) {
	f0 := &entity.Filter{
		Param:    "order.Status",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{status},
	}

	f1 := &entity.Filter{
		Param:    "order.ExNid",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{k.NID()},
	}
	fs := []*entity.Filter{f0, f1}
	pa := &entity.Paginated{
		Filters: fs,
	}
	err := k.repo.GetPaginated(pa, false)
	if err != nil {
		return nil, err
	}
	return pa.Orders, nil
}
