package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
	"time"
)

type PaginatedOrdersRequest struct {
	*PaginatedRequest
	Fs []*Filter `json:"filters"`
}

func (r *PaginatedOrdersRequest) Validate(userId int64) error {
	if r.PaginatedRequest == nil {
		r.PaginatedRequest = &PaginatedRequest{}
	}
	r.PaginatedRequest.Validate()
	for _, f := range r.Fs {
		if err := r.ValidateFilters(f); err != nil {
			return err
		}

	}

	if userId != 0 {
		for _, f := range r.Fs {
			if f.Param == "userId" {
				if f.Operator != "eq" || f.Values[0].(int64) != userId {
					return fmt.Errorf("userId must be equal to %d", userId)
				}
				return nil
			}
		}

		r.Fs = append(r.Fs, &Filter{
			Param:    "userId",
			Operator: "eq",
			Values:   []interface{}{userId},
		})
	}

	return nil
}

func (r *PaginatedOrdersRequest) Map() *entity.Paginated {
	fs := []*entity.Filter{}

	for _, f := range r.Fs {
		fs = append(fs, f.ToEntity())
	}

	return &entity.Paginated{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Filters: fs,
		Orders:  []entity.Order{},
	}

}

type PaginatedOrdersResp struct {
	*PaginatedResponse
	Orders []interface{} `json:"orders"`
}

func OrderResponse(p *entity.Paginated, admin bool) *PaginatedOrdersResp {
	r := &PaginatedOrdersResp{PaginatedResponse: PaginateResp(p, len(p.Orders))}

	r.Orders = []interface{}{}
	if admin {
		for _, o := range p.Orders {
			r.Orders = append(r.Orders, interface{}(adminOrderFromEntity(o)))
		}

	} else {
		for _, o := range p.Orders {
			r.Orders = append(r.Orders, interface{}(userOrderFromEntity(o)))
		}
	}
	return r
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
	case "userId":

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
	case "createdAt":
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
