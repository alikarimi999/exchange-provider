package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (k *kucoinExchange) TxIdSetted(ord entity.Order, txId string) error {
	agent := k.agent("TxIdSetted")
	o := ord.(*types.Order)
	p, err := k.pairs.Get(k.Id(), o.In.String(), o.Out.String())
	if err != nil {
		return err
	}

	var out *entity.Token
	if p.T1.String() == o.Out.String() {
		out = p.T1
	} else {
		out = p.T2
	}

	f, err := k.exchangeFeeAmount(out, p)
	if err != nil {
		k.l.Debug(agent, err.Error())
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("try again"))
	}
	if err := k.orderFeeRat(p); err != nil {
		k.l.Debug(agent, err.Error())
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("try again"))
	}

	o.Deposit.TxId = txId
	o.Status = types.ODepositTxIdSetted
	if err := k.repo.Update(o); err != nil {
		return errors.Wrap(errors.ErrInternal)
	}
	o.ExchangeFeeAmount = f
	go k.handleOrder(o, p)
	return nil
}

func (k *kucoinExchange) handleOrder(o *types.Order, p *entity.Pair) {
	if o.Status == types.ODepositTxIdSetted {
		var dc *Token
		if p.T1.String() == o.In.String() {
			dc = p.T1.ET.(*Token)
		} else {
			dc = p.T2.ET.(*Token)
		}

		k.trackDeposit(o, dc)
		k.repo.Update(o)
		k.cache.removeD(o.Deposit.TxId)
		k.cache.proccessedD(o.Deposit.TxId)

		if o.Status != types.ODepositeConfimred {
			return
		}
	}

	var (
		bc, qc, wc *Token
		swaps      []uint
	)
	if len(o.Swaps) == 2 {
		swaps = []uint{0, 1}
	} else {
		swaps = []uint{0}
	}

	for i := range swaps {
		bc, qc, wc = getBcQcWcFeeRate(o, p, i)
		if (i == 0 && o.Status == types.OFirstSwapCompleted) ||
			(o.Status == types.OSecondSwapCompleted) {
			continue
		}

		if err := k.swap(o, bc, qc, i); err != nil {
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
			o.Status = types.OFirstSwapTracking
		} else {
			o.Status = types.OSecondSwapTracking
		}
		k.repo.Update(o)

		if err := k.trackSwap(o, bc, qc, i); err != nil {
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
	}

	k.withdrawal(o, wc, p, false)
	k.repo.Update(o)
}
