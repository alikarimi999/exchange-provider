package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (k *exchange) SetTxId(ord entity.Order, txId string) error {
	o := ord.(*types.Order)
	if ord.STATUS() != entity.OCreated {
		return errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage(fmt.Sprintf("unable to set txId for order in '%s' status", ord.STATUS())))
	}

	if o.Deposit.TxId != "" {
		return errors.Wrap(errors.ErrForbidden,
			errors.NewMesssage("txId for this order has setted before"))
	}

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

	f, _, err := k.exchangeFeeAmount(out, p)
	if err != nil {
		return err
	}
	if err := k.setOrderFeeRate(p); err != nil {
		return err
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

func (k *exchange) handleOrder(o *types.Order, p *entity.Pair) {
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

		if p.T1.Id.Symbol != p.T2.Id.Symbol {
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
	}

	k.withdrawal(o, wc, p, false)
	k.repo.Update(o)
}
