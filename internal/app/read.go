package app

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

// first try to read data from the cache and if not found try the persitence database
func (o *OrderUseCase) read(v interface{}) error {

	switch d := v.(type) {
	case *entity.UserOrder:
		var dd *entity.UserOrder
		var err error
		if d.Id > 0 && d.UserId > 0 {
			dd, err = readOrder(o.repo, o.cache, d.UserId, d.Id)
		} else if d.Seq > 0 && d.UserId > 0 {
			dd, err = readOrderBySeq(o.repo, o.cache, d.UserId, d.Seq)
		} else {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("order id or seq or userId not found"))
		}
		if err != nil {
			return err
		}

		*d = *dd
		return nil
	case *[]*entity.UserOrder:
		i := *d
		i, err := readAllUserOrders(o.repo, i[0].UserId)
		if err != nil {
			return err
		}
		*d = i
		return nil

	case *entity.PaginatedUserOrders:
		return readPagenateUserOrders(o.repo, d)

	default:
		errors.Wrap(errors.New(fmt.Sprintf("unsupported type %T", d)))
	}

	return errors.Wrap(errors.New("unsupported type"))
}
func readOrder(r entity.OrderRepo, c entity.OrderCache, userId, orderId int64) (*entity.UserOrder, error) {
	ord, er1 := c.Get(userId, orderId)
	if er1 != nil {
		var er2 error
		ord, er2 = r.Get(userId, orderId)
		if er2 != nil {
			if errors.ErrorCode(er2) == errors.ErrNotFound {
				return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("order %d for user %d not found", orderId, userId)))
			}
			return nil, errors.Wrap(errors.ErrInternal, errors.New(fmt.Sprintf("error ( %s ),\n error ( %s )", er1, er2)), fmt.Sprintf("order %d for user %d", orderId, userId))
		}
	}
	return ord, nil
}

func readOrderBySeq(r entity.OrderRepo, c entity.OrderCache, userId int64, seq int64) (*entity.UserOrder, error) {
	ord, er1 := c.GetBySeq(userId, seq)
	if er1 != nil {
		var er2 error
		ord, er2 = r.GetBySeq(userId, seq)
		if er2 != nil {
			if errors.ErrorCode(er2) == errors.ErrNotFound {
				return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage(fmt.Sprintf("order %d for user %d not found", seq, userId)))
			}
			return nil, errors.Wrap(errors.ErrInternal, errors.New(fmt.Sprintf("error ( %s ),\n error ( %s )", er1, er2)), fmt.Sprintf("order %d for user %d", seq, userId))
		}
	}
	return ord, nil
}

func readAllUserOrders(r entity.OrderRepo, userId int64) ([]*entity.UserOrder, error) {
	// read all orders from permenant db
	return r.GetAll(userId)
}

func readPagenateUserOrders(r entity.OrderRepo, pa *entity.PaginatedUserOrders) error {
	// read all orders from permenant db
	return r.GetPaginated(pa)

}
