package http

import "github.com/gin-gonic/gin"

func (r *Router) userRoutes() {

	u := r.gin.Group("/orders")
	{

		u.GET("/:orderId/:step", r.CheckAccess(false),
			func(ctx *gin.Context) {
				r.srv.GetStep(newContext(ctx, false))
			})

		u.POST("/get", r.CheckAccess(false),
			func(ctx *gin.Context) {
				r.srv.GetPaginatedForUser(newContext(ctx, false))
			})

		u.POST("/create", r.CheckAccess(true),
			func(ctx *gin.Context) {
				r.srv.NewOrder(newContext(ctx, false))
			})

		u.POST("/set_tx_id", r.CheckAccess(true),
			func(ctx *gin.Context) {
				r.srv.SetTxId(newContext(ctx, false))
			})

	}

	r.gin.POST("/estimate", r.CheckAccess(false),
		func(ctx *gin.Context) {
			r.srv.EstimateAmountOut(newContext(ctx, false))
		})

	r.gin.POST("/pairs", r.CheckAccess(true),
		func(ctx *gin.Context) {
			r.srv.GetPairs(newContext(ctx, false))
		})

}
