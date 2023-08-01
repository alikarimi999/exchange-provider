package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"math"
	"math/big"
)

func (ex *exchange) UpdateStatus(eo entity.Order) error {
	o := eo.(*types.Order)
	if o.TxId == "" {
		return nil
	}

	if eo.STATUS() == entity.OCreated || eo.STATUS() == entity.OPending {
		p, err := ex.pairs.Get(ex.Id(), o.In.String(), o.Out.String())
		if err != nil {
			return err
		}
		var out *entity.Token
		if p.T1.Id.String() == o.Out.String() {
			out = p.T1
		} else {
			out = p.T2
		}

		var (
			v       *big.Int
			pending bool
		)

		if out.Native {
			v, pending, err = ex.getNativeTransferAmount(o.TxId, out)
		} else {
			v, pending, err = ex.getTokenTransferAmount(o.TxId, out, o.Receiver)
		}

		if err != nil {
			if err == errInvalidTx {
				o.Status = entity.OFailed
				o.FailedDesc = err.Error()
				ex.repo.Update(o)
				return nil
			}
			return err
		}

		if pending {
			o.Status = entity.OPending
			ex.repo.Update(o)
			return nil
		}

		if v != nil {
			o.Status = entity.OCompleted
			o.AmountOut, _ = new(big.Float).Quo(new(big.Float).SetInt(v),
				big.NewFloat(math.Pow10(int(out.Decimals)))).Float64()
		}
		ex.repo.Update(o)
	}
	return nil
}
