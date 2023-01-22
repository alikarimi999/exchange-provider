package dto

import "exchange-provider/internal/entity"

type PaginatedRequest struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
}

func (r *PaginatedRequest) Validate() error {
	if r.CurrentPage < 1 {
		r.CurrentPage = 1
	}
	if r.PageSize < 1 {
		r.PageSize = 10
	}

	if r.PageSize > 100 {
		r.PageSize = 100
	}
	return nil
}

type PaginatedResponse struct {
	CurrentPage int64 `json:"currentPage"`
	PageSize    int64 `json:"pageSize"`
	TotalNum    int64 `json:"totalNumbers"`
	TotalPage   int64 `json:"totalPage"`
}

func PaginateResp(p *entity.Paginated, size int) *PaginatedResponse {
	r := &PaginatedResponse{}
	r.CurrentPage = p.Page
	r.PageSize = int64(size)
	r.TotalNum = p.Total
	// calc total_page
	if p.Total%p.PerPage == 0 {
		r.TotalPage = p.Total / p.PerPage
	} else {
		r.TotalPage = p.Total/p.PerPage + 1
	}
	return r
}
