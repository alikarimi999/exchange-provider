package dto

import (
	"exchange-provider/internal/entity"
)

type PaginatedOrdersRequest struct {
	*PaginatedRequest
	Fs []*Filter `json:"filters"`
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
