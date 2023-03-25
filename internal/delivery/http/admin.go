package http

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) adminRoutes() {
	a := r.gin.Group("/admin", gin.BasicAuth(gin.Accounts{
		r.user: r.pass,
	}))
	{

		ss := a.Group("/services")
		{
			ss.GET("", func(ctx *gin.Context) { r.GetServicesConfig(ctx) })
			ss.POST("/:service", func(ctx *gin.Context) { r.ChangeSerivcesConfig(ctx) })
			ss.POST("update_chains_fee", func(ctx *gin.Context) { r.srv.UpdateChainsFee(ctx) })

		}

		os := a.Group("/orders")
		{
			os.POST("/", func(ctx *gin.Context) {
				r.srv.GetPaginatedForAdmin(newContext(ctx, true))
			})
		}

		ps := a.Group("/pairs")
		{
			ps.POST("/add/:id", func(ctx *gin.Context) {
				r.srv.AddPairs(newContext(ctx, true))
			})

			// ps.POST("", func(ctx *gin.Context) {
			// 	r.srv.GetPairsToAdmin(newContext(ctx, true))
			// })

			ps.POST("/get_min_deposit", func(ctx *gin.Context) {
				r.srv.GetMinPairDeposit(newContext(ctx, true))
			})

			ps.POST("/change_min_deposit", func(ctx *gin.Context) {
				r.srv.ChangeMinDeposit(newContext(ctx, true))
			})

			ps.POST("/get_all_min_deposit", func(ctx *gin.Context) {
				r.srv.GetAllMinDeposit(newContext(ctx, true))
			})

			// ps.DELETE("", func(ctx *gin.Context) {
			// 	r.srv.RemovePair(newContext(ctx, true))
			// })
		}

		fee := a.Group("/fee")
		{
			fee.POST("/default", func(ctx *gin.Context) {
				r.srv.ChangeDefaultFee(newContext(ctx, true))
			})

			fee.GET("/default", func(ctx *gin.Context) {
				r.srv.GetDefaultFee(newContext(ctx, true))
			})

			fee.POST("/get_by_users", func(ctx *gin.Context) {
				r.srv.GetUsersFee(newContext(ctx, true))
			})

			fee.POST("/change_by_user", func(ctx *gin.Context) {
				r.srv.ChangeUserFee(newContext(ctx, true))
			})

		}

		spread := a.Group("/spread")
		{

			spread.GET("/get_all", func(ctx *gin.Context) {
				r.srv.GetAllPairsSpread(newContext(ctx, true))
			})

			spread.POST("/change", func(ctx *gin.Context) {
				r.srv.ChangePairSpread(newContext(ctx, true))
			})

			spread.GET("/default", func(ctx *gin.Context) {
				r.srv.GetDefaultSpread(newContext(ctx, true))
			})

			spread.POST("/default", func(ctx *gin.Context) {
				r.srv.ChangeDefaultSpread(newContext(ctx, true))
			})
		}

		es := a.Group("/exchanges")
		{
			es.GET("/list", func(ctx *gin.Context) {
				r.srv.GetExchangeList(newContext(ctx, true))
			})
			es.DELETE("/:id", func(ctx *gin.Context) {
				r.srv.RemoveExchange(newContext(ctx, true))
			})
			es.POST("/add/:name", func(ctx *gin.Context) { r.srv.AddExchange(newContext(ctx, true)) })
		}

		limiter := a.Group("/limiter")
		{
			limiter.GET("", func(ctx *gin.Context) {
				res := struct {
					GL struct {
						Max    uint64
						Period string
					} `json:"general_limiter"`
					Col struct {
						Max    uint64
						Period string
					} `json:"create_order_limiter"`
				}{
					GL: struct {
						Max    uint64
						Period string
					}{
						Max:    r.gls.conf.Max,
						Period: r.gls.conf.Period.String(),
					},
					Col: struct {
						Max    uint64
						Period string
					}{
						Max:    r.col.conf.Max,
						Period: r.col.conf.Period.String(),
					},
				}
				ctx.JSON(200, res)
			})

			limiter.POST("", func(ctx *gin.Context) {
				r.changeLimitersConf(ctx)
			})
		}
	}

}
