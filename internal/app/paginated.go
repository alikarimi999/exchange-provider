package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

func (u *OrderUseCase) GetPaginated(pa *entity.PaginatedOrders) error {
	const op = errors.Op("Order-UseCase.GetAllUserOrders")

	if err := u.read(pa); err != nil {
		switch errors.ErrorCode(err) {
		case errors.ErrNotFound:
			err = errors.Wrap(err, op, &ErrMsg{msg: "orders not found"})
			u.l.Debug(string(op), err.Error())
		default:
			err = errors.Wrap(err, op, &ErrMsg{msg: "get orders failed, internal error"})
			u.l.Error(string(op), err.Error())
		}

		return err
	}
	return nil

}
