package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	app    *app.OrderUseCase
	repo   entity.OrderRepo
	pairs  entity.PairsRepo
	fee    entity.FeeTable
	spread entity.SpreadTable
	api    entity.ApiService
	l      logger.Logger
	v      *viper.Viper
	cf     *chainsFee
}

func NewServer(app *app.OrderUseCase, v *viper.Viper, pairs entity.PairsRepo, api entity.ApiService,
	repo entity.OrderRepo, fee entity.FeeTable, spread entity.SpreadTable, l logger.Logger) *Server {
	s := &Server{
		app:    app,
		repo:   repo,
		pairs:  pairs,
		fee:    fee,
		spread: spread,
		l:      l,
		api:    api,
		v:      v,
		cf: &chainsFee{
			mux:   &sync.Mutex{},
			chain: make(map[string]float64),
			v:     v,
		},
	}

	s.cf.readConfig()
	return s
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
