package http

import "github.com/gin-gonic/gin"

func (o *Router) userRoutes() {

	u := o.gin.Group("/orders")
	{

		u.POST("/get", CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPaginatedForUser(newContext(ctx))
			})

		u.POST("/create", CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.NewUserOrder(newContext(ctx))
			})

		u.POST("/set_tx_id", CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.SetTxId(newContext(ctx))
			})

	}

	p := o.gin.Group("/pairs")
	{
		p.GET("", CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPairsToUser(newContext(ctx))
			})

	}

	f := o.gin.Group("/fee")

	{
		f.GET("", CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetFeeToUser(newContext(ctx))
			})
	}
}
