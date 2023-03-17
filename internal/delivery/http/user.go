package http

import "github.com/gin-gonic/gin"

func (o *Router) userRoutes() {

	u := o.gin.Group("/orders")
	{

		u.GET("/:orderId/:step", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetStep(newContext(ctx, false))
			})

		u.POST("/get", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.GetPaginatedForUser(newContext(ctx, false))
			})

		u.POST("/create", Limiter(o.col), o.auth.CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.NewOrder(newContext(ctx, false))
			})

		u.POST("/set_tx_id", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "write", o.l),
			func(ctx *gin.Context) {
				o.srv.SetTxId(newContext(ctx, false))
			})

	}

	t := o.gin.Group("/tokens")
	{
		t.GET("", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
			func(ctx *gin.Context) {
				o.srv.Tokens(newContext(ctx, false))
			})
	}

	o.gin.POST("/estimate", Limiter(o.gls.addLimiter()), o.auth.CheckAccess("orders", "read", o.l),
		func(ctx *gin.Context) {
			o.srv.EstimateAmountOut(newContext(ctx, false))
		})

}
