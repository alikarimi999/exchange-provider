package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"order_service/pkg/logger"

	"github.com/gin-gonic/gin"
)

type CheckAccessRequest struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
	CheckIp  bool   `json:"check_ip"`
	Ip       string `json:"ip"`
}

type CheckAccessResp struct {
	UserId    int64  `json:"user_id"`
	HasAccess bool   `json:"has_access"`
	Msg       string `json:"msg"`
}

func (a *authService) CheckAccess(resource, action string, l logger.Logger) gin.HandlerFunc {
	const agent = "CheckAccess"
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("X-API-Key")
		if len(token) == 0 {
			ctx.JSON(http.StatusUnauthorized, "X-API-Key header is missing")
			ctx.Abort()
			return
		}

		if len(token) != 40 || token[:8] != "api_key_" {
			ctx.JSON(http.StatusUnauthorized, "X-API-Key header is invalid")
			ctx.Abort()
			return
		}

		addr, _ := netip.ParseAddr(ctx.ClientIP())

		if !addr.Is4() && !addr.IsLoopback() {
			ctx.JSON(http.StatusUnauthorized, "only ipV4 supports")
			ctx.Abort()
			return
		}

		var ip string
		if addr.Is6() && addr.IsLoopback() {
			ip = "127.0.0.1"
		} else {
			ip = addr.String()
		}

		ca := &CheckAccessRequest{
			ID:       token,
			Resource: resource,
			Action:   action,
			CheckIp:  a.checkIP(),
			Ip:       ip,
		}

		respBody, _ := json.Marshal(ca)
		req, err := a.request(bytes.NewReader(respBody))
		if err != nil {
			l.Error(agent, err.Error())
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			ctx.Abort()
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			l.Error(agent, err.Error())
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			ctx.Abort()
			return
		}
		if resp.StatusCode != http.StatusOK {
			b, err := io.ReadAll(resp.Body)
			l.Error(agent, fmt.Sprintf("status code: %d, %s %s", resp.StatusCode, string(b), err))
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			ctx.Abort()
			return
		}

		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			l.Error(agent, err.Error())
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			ctx.Abort()
			return
		}

		cr := &CheckAccessResp{}
		err = json.Unmarshal(b, cr)
		if err != nil {
			l.Error(agent, err.Error())
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			ctx.Abort()
			return
		}

		l.Debug(agent, fmt.Sprintf("CheckAccessResp: %+v", cr))

		if !cr.HasAccess {
			ctx.JSON(http.StatusUnauthorized, cr.Msg)
			ctx.Abort()
			return
		}
		if cr.UserId == 0 {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		ctx.Set("user_id", cr.UserId)
		ctx.Next()
	}
}
