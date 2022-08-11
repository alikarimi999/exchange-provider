package http

import "github.com/gin-gonic/gin"

func (o *Router) adminRoutes() {
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

			ps.POST("", func(ctx *gin.Context) {
				o.srv.GetExchangesPairs(newContext(ctx))
			})

			ps.POST("/get_min", func(ctx *gin.Context) {
				o.srv.GetMinPairDeposit(newContext(ctx))
			})

			ps.POST("/get_all_min", func(ctx *gin.Context) {
				o.srv.GetAllMinDeposit(newContext(ctx))
			})

			ps.DELETE("", func(ctx *gin.Context) {
				o.srv.RemovePair(newContext(ctx))
			})
		}

		fee := a.Group("/fee")
		{
			fee.POST("/default", func(ctx *gin.Context) {
				o.srv.ChangeDefaultFee(newContext(ctx))
			})

			fee.GET("/default", func(ctx *gin.Context) {
				o.srv.GetDefaultFee(newContext(ctx))
			})

			fee.POST("/get_by_users", func(ctx *gin.Context) {
				o.srv.GetUsersFee(newContext(ctx))
			})

			fee.POST("/change_by_user", func(ctx *gin.Context) {
				o.srv.ChangeUserFee(newContext(ctx))
			})

		}

		spread := a.Group("/spread")
		{

			spread.GET("/get_all", func(ctx *gin.Context) {
				o.srv.GetAllPairsSpread(newContext(ctx))
			})

			spread.POST("/change", func(ctx *gin.Context) {
				o.srv.ChangePairSpread(newContext(ctx))
			})

			spread.GET("/default", func(ctx *gin.Context) {
				o.srv.GetDefaultSpread(newContext(ctx))
			})

			spread.POST("/default", func(ctx *gin.Context) {
				o.srv.ChangeDefaultSpread(newContext(ctx))
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
