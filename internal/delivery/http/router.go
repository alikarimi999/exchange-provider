package http

import (
	"exchange-provider/pkg/logger"

	"exchange-provider/internal/adapter/http"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Router struct {
	gin  *gin.Engine
	srv  *http.Server
	l    logger.Logger
	v    *viper.Viper
	api  entity.ApiService
	user string
	pass string

	col *rateLimiter
	gls *limiters
}

func (r *Router) Run(addr ...string) error {
	return r.gin.Run(addr...)
}

func NewRouter(app *app.OrderUseCase, repo entity.OrderRepo, pairs entity.PairsRepo,
	fee entity.FeeTable, api entity.ApiService, v *viper.Viper, spread entity.SpreadTable,
	l logger.Logger, user, pass string) *Router {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	router := &Router{
		gin: engine,
		srv: http.NewServer(app, v, pairs, api, repo, fee, spread, l),
		l:   l,
		v:   v,
		api: api,

		user: user,
		pass: pass,
	}
	router.setup()
	return router
}

func (o *Router) orderSrvGrpV0() {
	o.userRoutes()
	o.adminRoutes()
}
