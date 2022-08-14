package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Limiter(l *rateLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-API-Key")

		if len(token) > 9 && token[:8] != "api_key_" {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("X-API-Key header is invalid"))
			ctx.Abort()
			return
		}

		ip := ctx.ClientIP()

		if l.allow(ctx.Request.Context(), ip, token) {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusTooManyRequests, fmt.Sprintf("too many requests"))
			ctx.Abort()
		}
	}
}
