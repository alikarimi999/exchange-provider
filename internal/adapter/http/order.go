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

		ctx.JSON(dto.MultiStep(o, tx), nil)
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

	pa.Exs = s.exs.GetAllMap()
	if err := s.repo.GetPaginated(pa, false); err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(dto.OrderResponse(pa, false), nil)
}

func (s *Server) GetPaginatedForAdmin(ctx Context) {

	req := &dto.PaginatedReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Validate()
	pa := req.Map()
	pa.Exs = s.exs.GetAllMap()
	if err := s.repo.GetPaginated(pa, false); err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(dto.OrderResponse(pa, true), nil)
}

func (s *Server) GetOrder(ctx Context) {
	api := ctx.GetApi()
	oId := ctx.Param("orderId")

	f0 := &entity.Filter{
		Param:    "busid",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{api.BusId},
	}

	f1 := &entity.Filter{
		Param:    "id",
		Operator: entity.FilterOperatorEqual,
		Values:   []interface{}{oId},
	}
	pa := &entity.Paginated{
		Filters: []*entity.Filter{f0, f1},
	}

	pa.Exs = s.exs.GetAllMap()
	if err := s.repo.GetPaginated(pa, false); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if len(pa.Orders) == 0 {
		ctx.JSON(nil, errors.Wrap(errors.ErrNotFound))
		return
	}
	o := pa.Orders[0]
	ctx.JSON(dto.OrderFromEntityForUser(o), nil)
}
