package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

// first try to read data from the cache and if not found try the persitence database
func (o *OrderUseCase) read(v interface{}) error {

	switch d := v.(type) {
	case *entity.Order:
		var dd *entity.Order
		var err error
		if d.Id > 0 {
			dd, err = readOrder(o.repo, o.cache, d.Id)
		} else {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("orderId not found"))
		}
		if err != nil {
			return err
		}

		*d = *dd
		return nil
	case *[]*entity.Order:
		i := *d
		i, err := readAllUserOrders(o.repo, i[0].UserId)
		if err != nil {
			return err
		}
		*d = i
		return nil

	case *entity.PaginatedOrders:
		return readPagenateUserOrders(o.repo, d)

	default:
		errors.Wrap(errors.New(fmt.Sprintf("unsupported type %T", d)))
	}

	return errors.Wrap(errors.New("unsupported type"))
}
func readOrder(r entity.OrderRepo, c entity.OrderCache, orderId int64) (*entity.Order, error) {
	ord, er1 := c.Get(orderId)
	if er1 != nil {
		var er2 error
		ord, er2 = r.Get(orderId)
		if er2 != nil {
			if errors.ErrorCode(er2) == errors.ErrNotFound {
				return nil, errors.Wrap(errors.ErrNotFound,
					errors.NewMesssage(fmt.Sprintf("order %d not found", orderId)))
			}
			return nil, errors.Wrap(errors.ErrInternal, errors.New(fmt.Sprintf("error ( %s ),\n error ( %s )",
				er1, er2)), fmt.Sprintf("order %d ", orderId))
		}
	}
	return ord, nil
}

func readAllUserOrders(r entity.OrderRepo, userId int64) ([]*entity.Order, error) {
	// read all orders from permenant db
	return r.GetAll(userId)
}

func readPagenateUserOrders(r entity.OrderRepo, pa *entity.PaginatedOrders) error {
	// read all orders from permenant db
	return r.GetPaginated(pa)

}
