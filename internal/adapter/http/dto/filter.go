package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
	"time"
)

type Filter struct {
	Param    string        `json:"param"`
	Operator string        `json:"operator"`
	Values   []interface{} `json:"values"`
}

func (f *Filter) ToEntity() *entity.Filter {
	return &entity.Filter{
		Param:    f.Param,
		Operator: entity.ParseFilterOperator(f.Operator),
		Values:   f.Values,
	}
}

func (r *PaginatedOrdersRequest) ValidateFilters(f *Filter) error {

	if f.Param == "" || f.Values == nil {
		return fmt.Errorf("invalid filter : %+v", f)
	}

	switch f.Operator {
	case "eq", "neq", "gt", "gte", "lt", "lte":
		if len(f.Values) != 1 {
			return fmt.Errorf("for this operators `eq`, `neq`, `gt`, `gte`, `lt`, `lte` only one value is allowed, but got %d", len(f.Values))
		}
	case "in", "notin":
		if len(f.Values) == 0 {
			return fmt.Errorf("for this operators `in`, `notin` at least one value is required, but got %d", len(f.Values))

		}

	case "between":
		if len(f.Values) != 2 {
			return fmt.Errorf("for this operators `between` only two values are allowed, but got %d", len(f.Values))
		}
	default:
		return fmt.Errorf("invalid operators : '%s'", f.Operator)
	}

	switch f.Param {
	case "user_id":

		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return fmt.Errorf("invalid value type for param : %s, expected number, but got %T", f.Param, v)
			}
			f.Values[i] = int64(n)
		}

		return nil

	case "id":
		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return fmt.Errorf("invalid value type for param : %s, expected number, but got %T", f.Param, v)
			}
			f.Values[i] = n
		}

	case "status":

		for i, v := range f.Values {
			s, ok := f.Values[i].(string)
			if !ok {
				return fmt.Errorf("invalid value type for param : %s, expected string, but got %T", f.Param, v)
			}

			f.Values[i] = s

		}
		return nil

	case "exchange":
		for i, v := range f.Values {
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("invalid value type for param : %s, expected string, but got %T", f.Param, v)
			}

			f.Values[i] = s

		}
		return nil

		// query between two dates
		// recieve two dates in epoch format
		// only operators `in` is allowed
	case "created_at":
		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return fmt.Errorf("invalid value type for param : %s, expected number, but got %T", f.Param, v)
			}
			f.Values[i] = time.Unix(int64(n), 0)
		}
		return nil

	default:
		return fmt.Errorf("invalid param '%s'", f.Param)
	}
	return nil
}
