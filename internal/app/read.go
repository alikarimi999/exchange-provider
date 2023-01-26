package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

// first try to read data from the cache and if not found try the persitence database
func (o *OrderUseCase) read(v interface{}) error {

	switch d := v.(type) {
	case *entity.CexOrder:
		if d.Id != "" {
			dd, err := o.repo.Get(d.Id)
			if err != nil {
				return err
			}
			if dd.Type() != entity.CEXOrder {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invaled operation"))
			}
			*d = *dd.(*entity.CexOrder)
			return nil
		}
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("orderId not found"))

	case *entity.EvmOrder:
		dd, err := o.repo.Get(d.Id)
		if err != nil {
			return err
		}

		*d = *dd.(*entity.EvmOrder)
		return nil

	case *entity.Paginated:
		return readPaginateUserOrders(o.repo, d)

	default:
		return fmt.Errorf("unsupported type %T", d)
	}
}

func readPaginateUserOrders(r entity.OrderRepo, pa *entity.Paginated) error {
	// read all orders from permenant db
	return r.GetPaginated(pa)

}
