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
	}
	router.orderSrvGrpV0()
	return router
}

func (o *Router) orderSrvGrpV0() {
	v0 := o.gin.Group("/orders")
	{
		v0.POST("", func(ctx *gin.Context) {
			o.srv.NewUserOrder(newContext(ctx))
		})

		v0.GET("/:userId/:id", func(ctx *gin.Context) {
			o.srv.GetUserOrder(newContext(ctx))
		})

		v0.GET("/:userId", func(ctx *gin.Context) {
			o.srv.GetAllUserOrders(newContext(ctx))
		})

	}

	admin := v0.Group("/admin")

	{
		admin.POST("/add_pairs", func(ctx *gin.Context) {
			o.srv.AddPairs(newContext(ctx))
		})

		admin.POST("/get_pairs", func(ctx *gin.Context) {
			o.srv.GetExchangesPairs(newContext(ctx))
		})
	}

}
