package dto

import (
	"fmt"
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"time"
)

type Filter struct {
	Param  string        `json:"param"`
	Cond   string        `json:"cond"`
	Values []interface{} `json:"values"`
}

func (f *Filter) ToEntity() *entity.Filter {
	return &entity.Filter{
		Param:  f.Param,
		Cond:   entity.ParseFilterCond(f.Cond),
		Values: f.Values,
	}
}

func (r *PaginatedUserOrdersRequest) ValidateFiltersForUser(f *Filter) error {

	if f.Param == "" || f.Values == nil {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("invalid filter : %+v", f)))
	}

	switch f.Cond {
	case "eq", "neq", "gt", "gte", "lt", "lte":
		if len(f.Values) != 1 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
				fmt.Sprintf("for this conditions `eq`, `neq`, `gt`, `gte`, `lt`, `lte` only one value is allowed, but got %d", len(f.Values))))
		}
	case "in", "notin":
		if len(f.Values) == 0 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
				fmt.Sprintf("for this conditions `in`, `notin` at least one value is required, but got %d", len(f.Values))))

		}

	case "between":
		if len(f.Values) != 2 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
				fmt.Sprintf("for this conditions `between` only two values are allowed, but got %d", len(f.Values))))
		}
	default:
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
			fmt.Sprintf("invalid condition : %s", f.Cond)))

	}

	switch f.Param {
	case "user_id":

		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
					fmt.Sprintf("invalid value type for param : %s, expected number, but got %T", f.Param, v)))
			}
			f.Values[i] = int64(n)
		}

		return nil

	case "id":
		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
					fmt.Sprintf("invalid value type for param : %s, expected number, but got %T", f.Param, v)))
			}
			f.Values[i] = n
		}

	case "status":

		for i, v := range f.Values {
			s, ok := f.Values[i].(string)
			if !ok {
				return errors.Wrap(errors.NewMesssage(fmt.Sprintf("invalid value type for param : %s, expected string, but got %T", f.Param, v)))
			}

			f.Values[i] = s

		}
		return nil

	case "base_coin", "quote_coin":
		for i, v := range f.Values {
			s, ok := v.(string)
			if !ok {
				return errors.Wrap(errors.NewMesssage(fmt.Sprintf("invalid value type for param : %s, expected string, but got %T", f.Param, v)))
			}

			f.Values[i] = s

		}
		return nil

	case "exchange":
		for i, v := range f.Values {
			s, ok := v.(string)
			if !ok {
				return errors.Wrap(errors.NewMesssage(fmt.Sprintf("invalid value type for param : %s, expected string, but got %T", f.Param, v)))
			}

			f.Values[i] = s

		}
		return nil

	case "side":
		for i, v := range f.Values {
			s, ok := v.(string)
			if !ok {
				return errors.Wrap(errors.NewMesssage(fmt.Sprintf("invalid value type for param : %s, expected string, but got %T", f.Param, v)))
			}

			if s != "buy" && s != "sell" {
				return errors.Wrap(errors.NewMesssage(fmt.Sprintf("side filter value must be one of `buy`, `sell`, got %s", v)))
			}

			f.Values[i] = s

		}

		// query between two dates
		// recieve two dates in epoch format
		// only condition `in` is allowed
	case "created_at":
		for i, v := range f.Values {
			n, ok := v.(float64)
			if !ok {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(
					fmt.Sprintf("invalid value type for param : %s, expected number, but got %T", f.Param, v)))
			}
			f.Values[i] = time.Unix(int64(n), 0)
		}
		return nil

	}
	return nil
}
