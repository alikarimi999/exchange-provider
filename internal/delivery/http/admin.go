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

			ps.POST("", func(ctx *gin.Context) {
				r.srv.GetPairs(newContext(ctx, true))
			})

			ps.POST("/update", func(ctx *gin.Context) {
				r.srv.UpdatePairs(newContext(ctx, true))
			})

			ps.POST("cmd", func(ctx *gin.Context) {
				r.srv.CommandPairs(newContext(ctx, true))
			})
		}

		fee := a.Group("/fee")
		{
			fee.POST("/default", func(ctx *gin.Context) {
				r.srv.ChangeDefaultFee(newContext(ctx, true))
			})

			fee.GET("/default", func(ctx *gin.Context) {
				r.srv.GetDefaultFee(newContext(ctx, true))
			})

			fee.GET("", func(ctx *gin.Context) {
				r.srv.GetFees(newContext(ctx, true))
			})

			fee.POST("/change-by_user", func(ctx *gin.Context) {
				r.srv.ChangeUserFee(newContext(ctx, true))
			})

		}

		spread := a.Group("/spread")
		{
			spread.GET("", func(ctx *gin.Context) {
				r.srv.GetAll(newContext(ctx, true))
			})

			spread.POST("", func(ctx *gin.Context) {
				r.srv.AddSpread(newContext(ctx, true))
			})

			spread.POST("/remove", func(ctx *gin.Context) {
				r.srv.RemoveSpread(newContext(ctx, true))
			})
		}

		es := a.Group("/exchanges")
		{
			es.GET("", func(ctx *gin.Context) {
				r.srv.GetExchangeList(newContext(ctx, true))
			})

			es.GET("/:id", func(ctx *gin.Context) {
				r.srv.GetExchangeList(newContext(ctx, true))
			})

			es.POST("/cmd", func(ctx *gin.Context) {
				r.srv.CommandExchanges(newContext(ctx, true))
			})

			es.POST("/add/:name", func(ctx *gin.Context) { r.srv.AddExchange(newContext(ctx, true)) })
		}

		api := a.Group("/api")
		{
			api.GET("/:id", func(ctx *gin.Context) {
				r.srv.GetApi(newContext(ctx, true))
			})
			api.POST("/generate", func(ctx *gin.Context) {
				r.srv.GenerateAPIToken(newContext(ctx, true))
			})

			api.POST("/add-ip", func(ctx *gin.Context) {
				r.srv.AddIP(newContext(ctx, true))
			})

			api.POST("remove-ip", func(ctx *gin.Context) {
				r.srv.RemoveIp(newContext(ctx, true))
			})

			api.POST("/level", func(ctx *gin.Context) {
				r.srv.UpdateLevel(newContext(ctx, true))
			})

			api.DELETE("/:id", func(ctx *gin.Context) {
				r.srv.Remove(newContext(ctx, true))
			})
		}
		// limiter := a.Group("/limiter")
		// {
		// 	limiter.GET("", func(ctx *gin.Context) {
		// 		res := struct {
		// 			GL struct {
		// 				Max    uint64
		// 				Period string
		// 			} `json:"general_limiter"`
		// 			Col struct {
		// 				Max    uint64
		// 				Period string
		// 			} `json:"create_order_limiter"`
		// 		}{
		// 			GL: struct {
		// 				Max    uint64
		// 				Period string
		// 			}{
		// 				Max:    r.gls.conf.Max,
		// 				Period: r.gls.conf.Period.String(),
		// 			},
		// 			Col: struct {
		// 				Max    uint64
		// 				Period string
		// 			}{
		// 				Max:    r.col.conf.Max,
		// 				Period: r.col.conf.Period.String(),
		// 			},
		// 		}
		// 		ctx.JSON(200, res)
		// 	})

		// 	limiter.POST("", func(ctx *gin.Context) {
		// 		r.changeLimitersConf(ctx)
		// 	})
		// }
	}

}
