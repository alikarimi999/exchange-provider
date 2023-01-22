package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"sort"
	"strings"
)

type PaginatedPairsRequest struct {
	*PaginatedRequest
	Price bool      `json:"price"`
	Fs    []*Filter `json:"filters"`
}

type ExchangePairs struct {
	Pairs []*AdminPair `json:"pairs"`
}

type PaginatedPairsResp struct {
	*PaginatedResponse
	Pairs interface{} `json:"pairs"`
}

func (r *PaginatedPairsRequest) Validate(admin bool) error {
	if r.PaginatedRequest == nil {
		r.PaginatedRequest = &PaginatedRequest{}
	}
	r.PaginatedRequest.Validate()
	return r.validateFilters(admin)
}

func (r *PaginatedPairsRequest) validateFilters(admin bool) error {
	if admin {
		for _, f := range r.Fs {
			if f.Param != "exchanges" && f.Operator != "in" {
				return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid filters"))
			}
		}
	}
	return nil
}

func (r *PaginatedPairsRequest) ToEntity() *entity.Paginated {
	fs := []*entity.Filter{}

	for _, f := range r.Fs {
		fs = append(fs, f.ToEntity())
	}

	return &entity.Paginated{
		Page:    r.CurrentPage,
		PerPage: r.PageSize,
		Filters: fs,
		Pairs:   []*entity.Pair{},
	}

}

func PairsResp(p *entity.Paginated, admin bool) *PaginatedPairsResp {
	r := &PaginatedPairsResp{PaginatedResponse: PaginateResp(p, len(p.Pairs))}

	sort.Slice(p.Pairs, func(i, j int) bool {
		pi := p.Pairs[i]
		pj := p.Pairs[j]
		return strings.Compare(pi.String(), pj.String()) == -1
	})

	if admin {
		pairs := make(map[string][]*AdminPair)
		for _, p := range p.Pairs {
			pairs[p.Exchange] = append(pairs[p.Exchange], entityToAdminPair(p))
		}
		r.Pairs = pairs
	} else {
		pairs := []*UserPair{}
		for _, p := range p.Pairs {
			pairs = append(pairs, EntityToPairUser(p))
		}
		r.Pairs = pairs
	}
	return r
}
