package dto

import (
	"exchange-provider/internal/entity"
	"fmt"
)

type PaginatedOrdersRequest struct {
	CurrentPage int64     `json:"current_page"`
	PageSize    int64     `json:"page_size"`
	Fs          []*Filter `json:"filters"`
}

func (r *PaginatedOrdersRequest) Validate(userId int64) error {
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
		if err := r.ValidateFilters(f); err != nil {
			return err
		}

	}

	if userId != 0 {
		for _, f := range r.Fs {
			if f.Param == "user_id" {
				if f.Operator != "eq" || f.Values[0].(int64) != userId {
					return fmt.Errorf("user_id must be equal to %d", userId)
				}
				return nil
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

func (r *PaginatedOrdersRequest) Map() *entity.PaginatedOrders {
	fs := []*entity.Filter{}

	for _, f := range r.Fs {
		fs = append(fs, f.ToEntity())
	}

	return &entity.PaginatedOrders{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Filters: fs,
		Orders:  []*entity.Order{},
	}

}

type PaginatedUserOrdersResponse struct {
	CurrentPage int64         `json:"current_page"`
	PageSize    int64         `json:"page_size"`
	TotalNum    int64         `json:"total_num"`
	TotalPage   int64         `json:"total_page"`
	Orders      []interface{} `json:"orders"`
}

func (r *PaginatedUserOrdersResponse) Map(po *entity.PaginatedOrders, admin bool) {

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
			r.Orders = append(r.Orders, interface{}(AdminOrderFromEntity(o)))
		}

	} else {
		for _, o := range po.Orders {
			r.Orders = append(r.Orders, interface{}(UOFromEntity(o)))
		}
	}
}
