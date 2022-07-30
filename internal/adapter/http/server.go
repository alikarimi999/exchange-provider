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

	o, err := s.app.NewUserOrder(req.UserId, req.Address, &entity.Coin{CoinId: req.BC, ChainId: req.BChain},
		&entity.Coin{CoinId: req.QC, ChainId: req.QChain}, req.Side)

	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.CreateOrderResponse{
		OrderId:         o.Id,
		DepositeId:      o.Deposite.Id,
		DepositeAddress: o.Deposite.Address,
	})
	return
}

func (s *Server) AdminGetUserOrder(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid user id")))
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid order id")))
		return
	}
	o, err := s.app.GetUserOrder(int64(userId), int64(id))
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.AdminUOFromEntity(o))
	return
}

func (s *Server) GetUserOrder(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid user id")))
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid order id")))
		return
	}
	o, err := s.app.GetUserOrder(int64(userId), int64(id))
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.UOFromEntity(o))
	return
}

func (s *Server) GetAllUserOrders(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid user id")))
		return
	}

	os, err := s.app.GetAllUserOrders(int64(userId))
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	osDTO := []*dto.AdminUserOrder{}
	for _, o := range os {
		osDTO = append(osDTO, dto.AdminUOFromEntity(o))
	}

	ctx.JSON(http.StatusOK, osDTO)
	return
}
