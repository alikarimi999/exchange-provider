package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strings"
)

func (s *Server) NewOrder(ctx Context) {
	api := ctx.GetApi()

	req := &dto.CreateOrderRequest{}
	if err := ctx.Bind(req); err != nil {
		err = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error()))
		ctx.JSON(nil, err)
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Input.ToUpper()
	req.Output.ToUpper()
	o, err := s.app.NewOrder(req.UserId, req.Sender, req.Refund, req.Receiver, req.Input,
		req.Output, req.AmountIn, req.LP, api)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	switch o.Type() {
	case entity.CEXOrder:
		ctx.JSON(dto.SingleStepResponse(o))
		return
	case entity.EVMOrder:
		tx, err := s.app.GetMultiStep(o, 1)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

		ctx.JSON(dto.MultiStep(o, tx, 1, 1), nil)
		return
	}
}

func (s *Server) GetPaginatedForUser(ctx Context) {
	api := ctx.GetApi()

	req := &dto.PaginatedReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Validate()
	pa := req.Map()

	fs := []*entity.Filter{}
	for _, f := range pa.Filters {
		if strings.ToLower(f.Param) == "busid" {
			continue
		}
		fs = append(fs, f)
	}

	f := &entity.Filter{
		Param:    "busid",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{api.BusId},
	}

	fs = append([]*entity.Filter{f}, fs...)
	pa.Filters = fs

	if err := s.repo.GetPaginated(pa, false); err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(dto.OrderResponse(pa, false), nil)
}

func (s *Server) GetPaginatedForAdmin(ctx Context) {

	pa := &dto.PaginatedReq{}
	if err := ctx.Bind(pa); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pao := pa.Map()
	if err := s.repo.GetPaginated(pao, false); err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(dto.OrderResponse(pao, true), nil)
}
