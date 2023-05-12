package app

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (o *OrderUseCase) RemoveExchange(id uint, all bool) error {

	f0 := &entity.Filter{
		Param:    "exchangeNid",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{},
	}

	f1 := &entity.Filter{
		Param:    "status",
		Operator: entity.FilterOperatorIN,
		Values:   []interface{}{entity.OCreated.String(), entity.OPending.String()},
	}

	pa := &entity.Paginated{
		Filters: []*entity.Filter{f0, f1},
	}

	if all {
		for _, ex := range o.exs.getAll() {
			f0.Values = append(f0.Values, ex.NID())
		}
		if err := o.repo.GetPaginated(pa, true); err != nil {
			return err
		}
		if pa.Total > 0 {
			return errors.Wrap(errors.ErrForbidden,
				fmt.Errorf("there are %d orders with 'created' and 'pending' status", pa.Total))
		}

		return o.exs.RemoveAll()
	}
	ex, err := o.exs.get(id)
	if err != nil {
		return err
	}
	f0.Values = append(f0.Values, ex.NID())
	if err := o.repo.GetPaginated(pa, true); err != nil {
		return err
	}
	if pa.Total > 0 {
		return errors.Wrap(errors.ErrForbidden,
			fmt.Errorf("there are %d orders with 'created' and 'pending' status", pa.Total))
	}
	return o.exs.Remove(id)
}

func (o *OrderUseCase) EnableDisable(id uint, enable, all bool) error {
	if all {
		return o.exs.enableDisableAll(enable)
	}
	return o.exs.enableDisable(id, enable)
}
