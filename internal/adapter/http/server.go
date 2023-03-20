package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	app   *app.OrderUseCase
	repo  entity.OrderRepo
	pairs entity.PairsRepo
	fee   entity.FeeService
	l     logger.Logger
	pc    entity.PairConfigs
	v     *viper.Viper
	cf    *chainsFee
}

func NewServer(app *app.OrderUseCase, v *viper.Viper, pairs entity.PairsRepo,
	repo entity.OrderRepo, fee entity.FeeService, pc entity.PairConfigs, l logger.Logger) *Server {
	s := &Server{
		app:   app,
		repo:  repo,
		pairs: pairs,
		fee:   fee,
		pc:    pc,
		l:     l,
		v:     v,
		cf: &chainsFee{
			mux:   &sync.Mutex{},
			chain: make(map[string]float64),
			v:     v,
		},
	}

	s.cf.readConfig()

	return s

}

func (s *Server) NewOrder(ctx Context) {
	// userId, _ := ctx.GetKey("user_id")

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

	in := req.Input.ToEntity()
	out := req.Output.ToEntity()

	o, err := s.app.NewOrder(req.UserId, *req.Refund, *req.Receiver, in, out, req.AmountIn, req.LP)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	switch o.Type() {
	case entity.CEXOrder:
		ctx.JSON(dto.SingleStepResponse(o.(*entity.CexOrder)), nil)
		return
	default:
		o := o.(*entity.EvmOrder)
		tx, isApproveTx, err := s.app.GetMultiStep(o, 1)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
		ctx.JSON(dto.MultiStep(o.ObjectId.String(), o.Sender.Hex(), tx, 1, len(o.Steps), isApproveTx), nil)
		return
	}
}

func (s *Server) GetPaginatedForUser(ctx Context) {
	// userId, _ := ctx.GetKey("user_id")

	req := &dto.PaginatedOrdersRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pa := req.Map()
	if err := s.app.GetPaginated(pa); err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(dto.OrderResponse(pa, false), nil)
}

func (s *Server) GetPaginatedForAdmin(ctx Context) {

	pa := &dto.PaginatedOrdersRequest{}
	if err := ctx.Bind(pa); err != nil {
		ctx.JSON(nil, err)
		return
	}

	pao := pa.Map()
	if err := s.app.GetPaginated(pao); err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(dto.OrderResponse(pao, true), nil)
}

func (s *Server) SetTxId(ctx Context) {
	// userId, _ := ctx.GetKey("user_id")
	r := &dto.SetTxIdRequest{}
	if err := ctx.Bind(r); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := r.Validate(); err != nil {
		ctx.JSON(nil, err)
		return
	}

	oId, err := dto.ParseId(r.Id, entity.PrefOrder)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := s.app.SetTxId(oId, r.TxId); err != nil {
		ctx.JSON(nil, err)
		return
	}
	r.Msg = "done"
	ctx.JSON(r, nil)

}
