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
	v0 := o.gin.Group("/orders")
	{

		v0.POST("/get", CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPaginatedForUser(newContext(ctx))
			})

		v0.POST("/create", CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.NewUserOrder(newContext(ctx))
			})

		v0.POST("/set_tx_id", CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.SetTxId(newContext(ctx))
			})

	}

	p := o.gin.Group("/pairs")
	{
		p.POST("/get", CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPairsToUser(newContext(ctx))
			})
	}

	a := o.gin.Group("/admin")

	{

		os := a.Group("/orders")
		{
			os.POST("/", func(ctx *gin.Context) {
				o.srv.GetPaginatedForAdmin(newContext(ctx))
			})
		}

		ps := a.Group("/pairs")
		{
			ps.POST("/add", func(ctx *gin.Context) {
				o.srv.AddPairs(newContext(ctx))
			})

			ps.POST("/get_all", func(ctx *gin.Context) {
				o.srv.GetExchangesPairs(newContext(ctx))
			})

			ps.DELETE("", func(ctx *gin.Context) {
				o.srv.RemovePair(newContext(ctx))
			})
		}

		fee := a.Group("/fee")
		{
			fee.POST("/", func(ctx *gin.Context) {
				o.srv.ChangeFee(newContext(ctx))
			})

			fee.GET("/", func(ctx *gin.Context) {
				o.srv.GetFee(newContext(ctx))
			})
		}

		es := a.Group("/exchanges")
		{
			es.POST("/list", func(ctx *gin.Context) {
				o.srv.GetExchangeList(newContext(ctx))
			})
			es.POST("/change_status", func(ctx *gin.Context) {
				o.srv.ChangeStatus(newContext(ctx))
			})
			es.POST("/add_account/:id", func(ctx *gin.Context) { o.srv.AddExchange(newContext(ctx)) })
		}
	}

}
