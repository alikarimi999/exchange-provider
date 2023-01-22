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

	if need {
		lastStep++
		o.Steps[lastStep] = &entity.EvmStep{Route: r, IsApprove: true}
	}
	lastStep++
	o.Steps[lastStep] = &entity.EvmStep{Route: r, IsApprove: false}

	return nil
}

func (d *EvmDex) GetStep(o *entity.EvmOrder, step uint) (*types.Transaction, error) {
	s := o.Steps[step]
	if s.IsApprove {
		return d.approveTx(s.Route, o.Receiver)
	}
	return d.createTx(s.Route, o.Sender, o.Receiver, o.AmountIn, o.FeeRate)
}
