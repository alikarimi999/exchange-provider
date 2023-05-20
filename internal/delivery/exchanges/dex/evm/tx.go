package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
)

func (d *evmDex) CreateTx(ord entity.Order) (entity.Tx, error) {
	o := ord.(*types.Order)
	p, err := d.pairs.Get(d.Id(), o.In.String(), o.Out.String())
	if err != nil {
		return nil, err
	}

	etx := &entity.EvmTx{}
	var in, out *entity.Token
	if p.T1.Id.String() == o.In.String() {
		in = p.T1
		out = p.T2
	} else {
		in = p.T2
		out = p.T1
	}

	if o.NeedApprove && !o.Approved {
		need, err := d.needApproval(in, o.Sender, o.AmountIn)
		if err != nil {
			return nil, err
		}
		if need {
			tx, err := d.approveTx(in)
			etx.IsApproveTx = true
			etx.Tx = tx
			return etx, err
		} else {
			o.Approved = true
			d.repo.Update(o)
		}
	}

	tx, err := d.createTx(in, out, o.Sender, o.Sender, o.Receiver, o.AmountIn, o.FeeAmount+o.ExchangeFeeAmount)
	if err != nil {
		return nil, err
	}

	etx.IsApproveTx = false
	etx.Tx = tx
	return etx, nil
}
