package http

import (
	"exchange-provider/pkg/logger"

	"exchange-provider/internal/adapter/http"
	"exchange-provider/internal/app"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

type Router struct {
	gin *gin.Engine
	srv *http.Server
	l   logger.Logger
	v   *viper.Viper

	auth *authService

	col *rateLimiter
	gls *limiters
}

func (r *Router) Run(addr ...string) {
	r.gin.Run(addr...)
}

func NewRouter(app *app.OrderUseCase, v *viper.Viper, rc *redis.Client, l logger.Logger) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	router := &Router{
		gin: engine,
		srv: http.NewServer(app, v, rc, l),
		l:   l,
		v:   v,
	}
	router.setup()
	return router
}

func (o *Router) orderSrvGrpV0() {
	o.userRoutes()
	o.adminRoutes()
}
