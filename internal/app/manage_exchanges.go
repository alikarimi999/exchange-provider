package app

import (
	"exchange-provider/pkg/errors"
	"fmt"
)

func (o *OrderUseCase) RemoveExchange(nid string, force bool) error {
	op := "OrderUseCase.RemoveExchange"

	// if !force {
	// 	// First, check whether there is a request being processed for this exchange.
	// 	// if there is, we cannot disable it
	// 	t, err := o.totalPendingOrders(ex.Exchange)
	// 	if err != nil {
	// 		return errors.Wrap(op, err)
	// 	}

	// 	if t > 0 {
	// 		o.l.Info(string(op), fmt.Sprintf("unable to remove exchange %s because there are %d pending orders", nid, t))
	// 		return errors.Wrap(errors.ErrBadRequest,
	// 			errors.NewMesssage(fmt.Sprintf("exchange %s has %d pending orders, so you can't remove it, unless force it", nid, t)))
	// 	}
	// }

	if err := o.exs.remove(nid); err != nil {
		return errors.Wrap(op, err)
	}

	o.l.Info(string(op), fmt.Sprintf("exchange %s removed", nid))
	return nil
}
