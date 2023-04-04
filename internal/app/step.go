package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u OrderUseCase) GetMultiStep(o *entity.DexOrder, step uint) (entity.Tx, error) {
	st, ok := o.Steps[uint(step)]
	if !ok {
		return nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("step out of range"))
	}
	ex, err := u.exs.getByNID(st.Exchange)
	if err != nil {
		return nil, err
	}

	approvedBefore := st.Approved
	tx, err := ex.(entity.EVMDex).GetStep(o, uint(step))
	if err != nil {
		return nil, err
	}

	if !approvedBefore && st.Approved {
		go u.write(o)
	}
	return tx, nil
}
