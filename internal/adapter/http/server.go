package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Server struct {
	app   *app.OrderUseCase
	pairs entity.PairRepo
	l     logger.Logger
	v     *viper.Viper
	cf    *chainsFee
}

func NewServer(pairs entity.PairRepo, app *app.OrderUseCase, v *viper.Viper,
	l logger.Logger) *Server {
	s := &Server{
		app:   app,
		pairs: pairs,
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
		ctx.JSON(nil, err)
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(nil, err)
		return
	}

	in, err := dto.ParseToken(req.In)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}
	out, err := dto.ParseToken(req.Out)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	routes, err := s.routing(in, out)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	ex, _ := s.app.GetExchange(routes[0].Exchange)
	var o entity.Order
	switch ex.Type() {
	case entity.CEX:
		o, err = s.app.NewCexOrder(req.UserId, &entity.Address{Addr: req.Receiver, Tag: req.Tag}, routes)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
	case entity.EvmDEX:
		o, err = s.app.NewEvmOrder(req.UserId, common.HexToAddress(req.Sender),
			common.HexToAddress(req.Receiver), req.AmountIn, routes[0])
		if err != nil {
			ctx.JSON(nil, err)
			return
		}
	}

	ctx.JSON(dto.CreateOrderResponse(o), nil)
}

func (s *Server) GetPaginatedForUser(ctx Context) {
	// userId, _ := ctx.GetKey("user_id")

	req := &dto.PaginatedOrdersRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := req.Validate(""); err != nil {
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

	if err := pa.Validate(""); err != nil {
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
