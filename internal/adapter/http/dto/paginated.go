package dto

import (
	"fmt"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type PaginatedUserOrdersRequest struct {
	CurrentPage int64     `json:"current_page"`
	PageSize    int64     `json:"page_size"`
	Fs          []*Filter `json:"filters"`
}

func (r *PaginatedUserOrdersRequest) Validate(userId int64) error {
	if r.CurrentPage < 1 {
		r.CurrentPage = 1
	}
	if r.PageSize < 1 {
		r.PageSize = 10
	}

	if r.PageSize > 100 {
		r.PageSize = 100
	}

	for _, f := range r.Fs {
		if err := r.ValidateFiltersForUser(f); err != nil {
			return err
		}

	}

	if userId != 0 {
		for _, f := range r.Fs {
			if f.Param == "user_id" {
				if f.Operator != "eq" || f.Values[0].(int64) != userId {
					return errors.Wrap(errors.ErrForbidden, errors.NewMesssage(fmt.Sprintf("user_id must be equal to %d", userId)))
				}
				return nil
			}
			// change param id with seq if exists
			if f.Param == "id" {
				f.Param = "seq"
			}

		}

		r.Fs = append(r.Fs, &Filter{
			Param:    "user_id",
			Operator: "eq",
			Values:   []interface{}{userId},
		})
	}

	return nil
}

func (r *PaginatedUserOrdersRequest) Map() *entity.PaginatedUserOrders {
	fs := []*entity.Filter{}

	for _, f := range r.Fs {
		fs = append(fs, f.ToEntity())
	}

	return &entity.PaginatedUserOrders{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Filters: fs,
		Orders:  []*entity.UserOrder{},
	}

}

type PaginatedUserOrdersResponse struct {
	CurrentPage int64         `json:"current_page"`
	PageSize    int64         `json:"page_size"`
	TotalNum    int64         `json:"total_num"`
	TotalPage   int64         `json:"total_page"`
	Orders      []interface{} `json:"orders"`
}

func (r *PaginatedUserOrdersResponse) Map(po *entity.PaginatedUserOrders, admin bool) {

	r.CurrentPage = po.Page
	r.PageSize = int64(len(po.Orders))
	r.TotalNum = po.Total
	// calc total_page
	if po.Total%po.PerPage == 0 {
		r.TotalPage = po.Total / po.PerPage
	} else {
		r.TotalPage = po.Total/po.PerPage + 1
	}
	r.Orders = []interface{}{}
	if admin {
		for _, o := range po.Orders {
			r.Orders = append(r.Orders, interface{}(AdminUOFromEntity(o)))
		}

	} else {
		for _, o := range po.Orders {
			r.Orders = append(r.Orders, interface{}(UOFromEntity(o)))
		}
	}
}
