package evm

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/core/types"
)

func (d *EvmDex) SetStpes(o *entity.EvmOrder, r *entity.Route) error {

	var lastStep uint
	need, err := d.needApproval(r, o.Sender, o.AmountIn)
	if err != nil {
		return err
	}

	lastStep++
	o.Steps[lastStep] = &entity.EvmStep{Route: r, NeedApprove: need}

	return nil
}

func (d *EvmDex) GetStep(o *entity.EvmOrder, step uint) (*types.Transaction, bool, error) {
	s := o.Steps[step]
	if s.NeedApprove && !s.Approved {
		need, err := d.needApproval(s.Route, o.Sender, o.AmountIn)
		if err != nil {
			return nil, false, err
		}
		if need {
			tx, err := d.approveTx(s.Route, o.Receiver)
			return tx, true, err
		} else {
			s.Approved = true
		}
	}

	tx, err := d.createTx(s.Route, o.Sender, o.Receiver, o.AmountIn, o.FeeRate)
	return tx, false, err
}
