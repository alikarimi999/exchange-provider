package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"net/http"
	"sync"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Server struct {
	app *app.OrderUseCase

	l  logger.Logger
	v  *viper.Viper
	rc *redis.Client
	cf *chainsFee
}

func NewServer(app *app.OrderUseCase, v *viper.Viper, rc *redis.Client, l logger.Logger) *Server {
	s := &Server{
		app: app,

		l:  l,
		v:  v,
		rc: rc,
		cf: &chainsFee{
			mux:   &sync.Mutex{},
			chain: make(map[string]float64),
		},
	}

	fmt.Println("Delete: http-server.go-NewServer")
	s.cf.chain["97"] = 1
	s.cf.chain["80001"] = 2

	return s

}

func (s *Server) NewUserOrder(ctx Context) {

	userId, _ := ctx.GetKey("user_id")
	req := dto.CreateOrderRequest{}
	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := req.Validate(); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	in, err := dto.ParseCoin(req.In)
	if err != nil {
		handlerErr(ctx, err)
		return
	}
	out, err := dto.ParseCoin(req.Out)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	routes, err := s.routing(in, out)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	o, err := s.app.NewOrder(userId.(int64), &entity.Address{Addr: req.Address, Tag: req.Tag}, routes)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.CreateOrderResponse{
		OrderId: o.Id,
		// DC:              dc,
		// MinDeposit:      minD,
		DepositeAddress: o.Deposit.Addr,
		AddressTag:      o.Deposit.Tag,
	})
}

func (s *Server) GetPaginatedForUser(ctx Context) {
	userId, _ := ctx.GetKey("user_id")

	pa := &dto.PaginatedUserOrdersRequest{}
	if err := ctx.Bind(pa); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, err.Error()))
		return
	}

	if err := pa.Validate(userId.(int64)); err != nil {
		handlerErr(ctx, err)
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
}

func (s *Server) SetTxId(ctx Context) {
	// userId, _ := ctx.GetKey("user_id")
	r := &dto.SetTxIdRequest{}
	if err := ctx.Bind(r); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := r.Validate(); err != nil {
		r.Msg = errors.ErrorMsg(err)
		ctx.JSON(http.StatusBadRequest, r)
		return
	}

	if err := s.app.SetTxId(r.Id, r.TxId); err != nil {
		r.Msg = errors.ErrorMsg(err)
		ctx.JSON(http.StatusBadRequest, r)
		return
	}

	r.Msg = "transaction id setted successfully"
	ctx.JSON(http.StatusOK, r)
}
