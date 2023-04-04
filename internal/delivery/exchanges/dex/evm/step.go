package evm

import (
	"exchange-provider/internal/entity"
)

func (d *evmDex) SetStpes(o *entity.DexOrder, r *entity.Route) error {

	var lastStep uint
	need, err := d.needApproval(r, o.Sender, o.AmountIn)
	if err != nil {
		return err
	}

	lastStep++
	o.Steps[lastStep] = &entity.Step{Route: r, NeedApprove: need}

	return nil
}

func (d *evmDex) GetStep(o *entity.DexOrder, step uint) (entity.Tx, error) {
	etx := &entity.EvmTx{}
	s := o.Steps[step]
	if s.NeedApprove && !s.Approved {
		need, err := d.needApproval(s.Route, o.Sender, o.AmountIn)
		if err != nil {
			return nil, err
		}
		if need {
			tx, err := d.approveTx(s.Route, o.Receiver)
			etx.IsApproveTx = true
			etx.Tx = tx
			return etx, err
		} else {
			s.Approved = true
		}
	}

	tx, err := d.createTx(s.Route, o.Sender, o.Sender, o.Receiver, o.AmountIn, o.FeeRate)
	if err != nil {
		return nil, err
	}

	etx.IsApproveTx = false
	etx.Tx = tx
	return etx, nil
}
