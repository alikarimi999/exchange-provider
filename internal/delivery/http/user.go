package http

import "github.com/gin-gonic/gin"

func (o *Router) userRoutes() {

	u := o.gin.Group("/orders")
	{

		u.POST("/get", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPaginatedForUser(newContext(ctx))
			})

		u.POST("/create", Limiter(o.col), o.auth.CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.NewUserOrder(newContext(ctx))
			})

		u.POST("/set_tx_id", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.SetTxId(newContext(ctx))
			})

	}

	p := o.gin.Group("/pairs")
	{
		p.GET("", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPairsToUser(newContext(ctx))
			})

	}

	f := o.gin.Group("/fee")

	{
		f.GET("", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetFeeToUser(newContext(ctx))
			})
	}
}
