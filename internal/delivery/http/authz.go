package http

import (
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils"
	"fmt"
	"net/netip"

	"github.com/gin-gonic/gin"
)

func (r *Router) CheckAccess(write bool) gin.HandlerFunc {
	const agent = "CheckAccess"
	return func(context *gin.Context) {
		ctx := newContext(context, false)

		token := ctx.GetHeader("X-API-Key")
		if len(token) == 0 {
			ctx.JSON(nil, errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage("X-API-Key header is missing")))
			context.Abort()
			return
		}

		addr, _ := netip.ParseAddr(context.ClientIP())

		if !addr.Is4() && !addr.IsLoopback() {
			ctx.JSON(nil, errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage("only ipV4 supports")))
			context.Abort()
			return
		}

		var ip string
		if addr.Is6() && addr.IsLoopback() {
			ip = "127.0.0.1"
		} else {
			ip = addr.String()
		}

		at, err := r.api.Get(utils.Hash(token))
		if err != nil {
			if errors.ErrorCode(err) == errors.ErrNotFound {
				ctx.JSON(nil, errors.Wrap(errors.ErrForbidden,
					errors.NewMesssage(fmt.Sprintf("api key '%s' not found", token))))
				context.Abort()
				return
			}
			ctx.JSON(nil, errors.Wrap(errors.ErrForbidden))
			context.Abort()
			return
		}

		if write && !at.Write {
			ctx.JSON(nil, errors.Wrap(errors.ErrForbidden,
				errors.NewMesssage(fmt.Sprintf("X-API-Key doesn't have write access", token))))
			context.Abort()
			return
		}

		if at.CheckIp {
			var ipValid bool
			for _, i := range at.Ips {
				if ip == i {
					ipValid = true
					break
				}
			}

			if !ipValid {
				ctx.JSON(nil, errors.Wrap(errors.ErrForbidden,
					errors.NewMesssage(fmt.Sprintf("invalid ip %s", ip))))
				context.Abort()
				return
			}
		}

		ctx.SetApi(at)
		context.Next()
	}
}
