package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/core/types"
)

func (u OrderUseCase) GetMultiStep(o *entity.EvmOrder, step uint) (*types.Transaction, bool, error) {
	st, ok := o.Steps[uint(step)]
	if !ok {
		return nil, false, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("step out of range"))
	}
	ex, err := u.GetExchange(st.Exchange)
	if err != nil {
		return nil, false, err
	}

	approvedBefore := st.Approved
	tx, isApproveTx, err := ex.(entity.EVMDex).GetStep(o, uint(step))
	if err != nil {
		return nil, false, err
	}

	if !approvedBefore && st.Approved {
		go u.write(o)
	}
	return tx, isApproveTx, nil
}
