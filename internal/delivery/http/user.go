package http

import "github.com/gin-gonic/gin"

func (r *Router) userRoutes() {

	u := r.gin.Group("/orders")
	{

		u.GET("/:orderId/:step", r.checkAccess(false),
			func(ctx *gin.Context) {
				r.srv.GetStep(newContext(ctx, false))
			})

		u.GET("/:orderId", r.checkAccess(false),
			func(ctx *gin.Context) { r.srv.GetOrder(newContext(ctx, false)) })
		u.POST("/get", r.checkAccess(false),
			func(ctx *gin.Context) {
				r.srv.GetPaginatedForUser(newContext(ctx, false))
			})

		u.POST("/create", r.checkAccess(true),
			func(ctx *gin.Context) {
				r.srv.NewOrder(newContext(ctx, false))
			})

		u.POST("/set_tx_id", r.checkAccess(true),
			func(ctx *gin.Context) {
				r.srv.SetTxId(newContext(ctx, false))
			})

	}

	r.gin.POST("/tokens/allowance", r.checkAccess(false),
		func(ctx *gin.Context) { r.srv.Allowance(newContext(ctx, false)) })

	r.gin.POST("/estimate", r.checkAccess(false),
		func(ctx *gin.Context) {
			r.srv.EstimateAmountOut(newContext(ctx, false))
		})

	r.gin.POST("/pairs", r.checkAccess(true),
		func(ctx *gin.Context) {
			r.srv.GetPairs(newContext(ctx, false))
		})

}
