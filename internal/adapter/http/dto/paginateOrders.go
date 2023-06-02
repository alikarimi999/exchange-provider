package dto

import (
	"exchange-provider/internal/entity"
)

type PaginatedReq struct {
	PaginatedRequest
	Desc bool      `json:"desc"`
	Fs   []*Filter `json:"filters"`
}

func (r *PaginatedReq) Map() *entity.Paginated {
	fs := []*entity.Filter{}
	for _, f := range r.Fs {
		if len(f.Values) > 0 {
			fs = append(fs, f.ToEntity())
		}
	}

	return &entity.Paginated{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Desc:    r.Desc,
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
			r.Orders = append(r.Orders, interface{}(o))
		}

	} else {
		for _, o := range p.Orders {
			r.Orders = append(r.Orders, interface{}(OrderFromEntityForUser(o)))
		}
	}
	return r
}
