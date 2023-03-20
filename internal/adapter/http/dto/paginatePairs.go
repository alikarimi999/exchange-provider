package dto

import (
	"exchange-provider/internal/entity"
)

type PaginatedPairsRequest struct {
	PaginatedRequest
	Fs []*Filter `json:"filters"`
}

type PaginatedPairsResp struct {
	PaginatedResponse
	Pairs []Pair `json:"pairs"`
}

func (r *PaginatedPairsRequest) ToEntity() *entity.Paginated {
	fs := []*entity.Filter{}

	for _, f := range r.Fs {
		fs = append(fs, f.ToEntity())
	}

	r.PaginatedRequest.Validate()
	return &entity.Paginated{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Filters: fs,
		Pairs:   []*entity.Pair{},
	}

}
