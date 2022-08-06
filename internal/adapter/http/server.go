package http

import (
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"strconv"

	"order_service/pkg/errors"
)

type Server struct {
	app *app.OrderUseCase
	l   logger.Logger
}

func NewServer(app *app.OrderUseCase, l logger.Logger) *Server {
	return &Server{
		app: app,
		l:   l,
	}
}

func (s *Server) NewUserOrder(ctx Context) {

	req := dto.CreateOrderRequest{}
	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := req.Validate(); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	bc := &entity.Coin{CoinId: req.BC, ChainId: req.BChain}
	qc := &entity.Coin{CoinId: req.QC, ChainId: req.QChain}

	ex, err := s.app.SelectExchangeByPair(bc, qc)
	if err != nil {
		handlerErr(ctx, err)
	}

	o, err := s.app.NewUserOrder(req.UserId, req.Address, bc, qc, req.Side, ex)

	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.CreateOrderResponse{
		OrderId:         o.Id,
		DepositeId:      o.Deposite.Id,
		DepositeAddress: o.Deposite.Address,
		AddressTag:      o.Deposite.Tag,
	})
	return
}

func (s *Server) GetPaginatedForUser(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid user id")))
		return
	}

	pa := &dto.PaginatedUserOrdersRequest{}
	if err := ctx.Bind(pa); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := pa.Validate(int64(userId)); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	pao := pa.Map()
	if err := s.app.GetPaginated(pao); err != nil {
		handlerErr(ctx, err)
		return
	}
	r := &dto.PaginatedUserOrdersResponse{}
	r.Map(pao, false)
	ctx.JSON(http.StatusOK, r)
	return
}

func (s *Server) GetPaginatedForAdmin(ctx Context) {

	pa := &dto.PaginatedUserOrdersRequest{}
	if err := ctx.Bind(pa); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := pa.Validate(0); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	pao := pa.Map()
	if err := s.app.GetPaginated(pao); err != nil {
		handlerErr(ctx, err)
		return
	}
	r := &dto.PaginatedUserOrdersResponse{}
	r.Map(pao, true)
	ctx.JSON(http.StatusOK, r)
	return
}
