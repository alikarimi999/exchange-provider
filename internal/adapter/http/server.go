package http

import (
	"fmt"
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
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
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, "invalid request"))
		return
	}

	if !req.Validate() {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, fmt.Sprintf("invalid request %+v", req)))
		return
	}

	r, err := dto.Coin(req.RCoin, req.RChain)
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	p, err := dto.Coin(req.PCoin, req.PChain)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	o, err := s.app.NewUserOrder(req.UserId, req.Address, r, p)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.CreateOrderResponse{
		Id:              o.Id,
		DepositeId:      o.Deposite.Id,
		DepositeAddress: o.Deposite.Address,
	})
	return
}

func (s *Server) GetUserOrder(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, "invalid user id"))
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, "invalid order id"))
		return
	}
	o, err := s.app.GetUserOrder(int64(userId), int64(id))
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.UoFromEntity(o))
	return
}

func (s *Server) GetAllUserOrders(ctx Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, "invalid user id"))
		return
	}

	os, err := s.app.GetAllUserOrders(int64(userId))
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	osDTO := []*dto.UserOrder{}
	for _, o := range os {
		osDTO = append(osDTO, dto.UoFromEntity(o))
	}

	ctx.JSON(http.StatusOK, osDTO)
	return
}
