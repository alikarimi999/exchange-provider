package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (u *OrderUseCase) RemovePair(ex entity.Exchange, t1, t2 entity.TokenId) error {
	f0 := &entity.Filter{
		Param:    "exchangeNid",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{ex.NID()},
	}

	f1 := &entity.Filter{
		Param:    "pairId",
		Operator: entity.FilterOperatorIN,
		Values: []interface{}{fmt.Sprintf("%s/%s", t1.String(), t2.String()),
			fmt.Sprintf("%s/%s", t2.String(), t1.String())},
	}

	f2 := &entity.Filter{
		Param:    "status",
		Operator: entity.FilterOperatorIN,
		Values:   []interface{}{entity.OCreated.String(), entity.OPending.String()},
	}

	pa := &entity.Paginated{
		Filters: []*entity.Filter{f0, f1, f2},
	}
	if err := u.repo.GetPaginated(pa, true); err != nil {
		return err
	}
	if pa.Total > 0 {
		return errors.Wrap(errors.ErrForbidden,
			fmt.Errorf("there are %d orders with 'created' and 'pending' status", pa.Total))
	}
	return ex.RemovePair(t1, t2)
}
