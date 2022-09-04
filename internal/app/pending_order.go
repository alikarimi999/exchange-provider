package app

import (
	"order_service/internal/entity"
)

func (o *OrderUseCase) totalPendingOrders(ex entity.Exchange, fs ...*entity.Filter) (total int64, err error) {
	f1 := &entity.Filter{
		Param:    "exchange",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{ex.NID()},
	}

	f2 := &entity.Filter{
		Param:    "status",
		Operator: entity.FilterOperatorNotIn,
		Values:   []interface{}{"succeed", "failed"},
	}

	f3 := &entity.Filter{
		Param:    "broken",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{false},
	}

	pa := &entity.PaginatedUserOrders{
		Page:    1,
		PerPage: 1,
		Total:   0,
		Filters: []*entity.Filter{f1, f2, f3},
		Orders:  []*entity.UserOrder{},
	}

	pa.Filters = append(pa.Filters, fs...)

	if err = o.GetPaginated(pa); err != nil {
		return 0, err
	}

	return pa.Total, nil
}
