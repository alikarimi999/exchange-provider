package binance

import (
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (ex *exchange) TxIdSetted(ord entity.Order, txId string) error {
	agent := ex.agent("TxIdSetted")
	o := ord.(*types.Order)
	p, err := ex.pairs.Get(ex.Id(), o.In.String(), o.Out.String())
	if err != nil {
		return err
	}

	var out *entity.Token
	if p.T1.String() == o.Out.String() {
		out = p.T1
	} else {
		out = p.T2
	}

	f, err := ex.exchangeFeeAmount(out, p)
	if err != nil {
		ex.l.Debug(agent, err.Error())
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("try again"))
	}
	if err := ex.setOrderFeeRate(p); err != nil {
		ex.l.Debug(agent, err.Error())
		return errors.Wrap(errors.ErrInternal, errors.NewMesssage("try again"))
	}

	o.Deposit.TxId = txId
	o.Status = types.ODepositTxIdSetted
	if err := ex.repo.Update(o); err != nil {
		return errors.Wrap(errors.ErrInternal)
	}
	o.ExchangeFeeAmount = f
	go ex.handleOrder(o, p)
	return nil
}

func (ex *exchange) handleOrder(o *types.Order, p *entity.Pair) {
	if o.Status == types.ODepositTxIdSetted {
		var dc *Token
		if p.T1.String() == o.In.String() {
			dc = p.T1.ET.(*Token)
		} else {
			dc = p.T2.ET.(*Token)
		}
		ex.trackDeposit(o, dc)
		ex.repo.Update(o)

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

		if err := ex.swap(o, bc, qc, i); err != nil {
			if i == 0 {
				o.Status = types.OFirstSwapFailed
			} else {
				o.Status = types.OSecondSwapFailed
			}
			o.FailedDesc = err.Error()
			ex.repo.Update(o)
			return
		}

		if i == 0 {
			o.Status = types.OFirstSwapCompleted
		} else {
			o.Status = types.OSecondSwapCompleted
		}
		ex.repo.Update(o)
	}

	ex.withdrawal(o, wc, p)
	ex.repo.Update(o)
}
