package http

import (
	"order_service/pkg/logger"

	"order_service/internal/adapter/http"
	"order_service/internal/app"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
	srv *http.Server
	l   logger.Logger
}

func (r *Router) Run(addr ...string) {
	r.gin.Run(addr...)
}

func NewRouter(app *app.OrderUseCase, l logger.Logger) *Router {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	router := &Router{
		gin: engine,
		srv: http.NewServer(app, l),
		l:   l,
	}
	router.orderSrvGrpV0()
	return router
}

func (o *Router) orderSrvGrpV0() {
	o.userRoutes()
	o.adminRoutes()
}
