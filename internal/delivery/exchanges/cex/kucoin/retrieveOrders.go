package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/dto"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"
	"time"
)

func (k *exchange) retreiveOrders() error {
	if err := k.da.aggregateAll(-2 * time.Hour); err != nil {
		return err
	}

	ws, err := k.wa.aggregateAll(-2*time.Hour, true)
	if err != nil {
		return err
	}

	pairs := make(map[string]*entity.Pair)

	os0, err := k.getOrders(types.ODepositTxIdSetted)
	if err != nil {
		return err
	}

	for _, o := range os0 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	os1, err := k.getOrders(types.ODepositeConfimred)
	if err != nil {
		return err
	}

	for _, o := range os1 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	os2, err := k.getOrders(types.OFirstSwapTracking)
	if err != nil {
		return err
	}

	for _, o := range os2 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	os3, err := k.getOrders(types.OSecondSwapTracking)
	if err != nil {
		return err
	}
	for _, o := range os3 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	os4, err := k.getOrders(types.OFirstSwapCompleted)
	if err != nil {
		return err
	}

	for _, o := range os4 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	os5, err := k.getOrders(types.OSecondSwapCompleted)
	if err != nil {
		return err
	}
	for _, o := range os5 {
		to := o.(*types.Order)
		p, err := k.pairs.Get(k.Id(), to.In.String(), to.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.String() == to.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		f, err := k.exchangeFeeAmount(out, p)
		if err != nil {
			return err
		}
		to.ExchangeFeeAmount = f
		pairs[o.ID().String()] = p
	}

	for i := 0; i < len(os4); i++ {
		if len(os4[i].(*types.Order).Swaps) == 1 {
			os5 = append(os5, os4[i])
			os4 = append(os4[:i], os4[i+1:]...)
			i--
		}
	}

	k.orderCheckAndFixWithdraw(os5, ws, pairs)

	os := []entity.Order{}
	os = append(os, os1...)
	os = append(os, os2...)
	os = append(os, os3...)
	os = append(os, os4...)

	for _, o := range os {
		k.orderCheckAndFixSwaps(o.(*types.Order), pairs[o.ID().String()])
	}

	for _, o := range os0 {
		go func(to *types.Order, p *entity.Pair) {
			k.handleOrder(to, p)
		}(o.(*types.Order), pairs[o.ID().String()])
	}

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
