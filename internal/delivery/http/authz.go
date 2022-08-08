package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func CheckAccess(resource, action string, l logger.Logger) gin.HandlerFunc {
	const agent = "CheckAccess"
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("X-API-Key")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("X-API-Key header is required"))
			ctx.Abort()
			return
		}

		if token[:8] != "api_key_" {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("X-API-Key header is invalid"))
			ctx.Abort()
			return
		}

		ca := &CheckAccessRequest{
			ID:       token,
			Resource: resource,
			Action:   action,
			CheckIp:  true,
			Ip:       ctx.ClientIP(),
		}

		respBody, _ := json.Marshal(ca)
		bytes.NewReader(respBody)

		req, _ := http.NewRequest("POST", "http://localhost:9091/actores/access", bytes.NewReader(respBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}
		if resp.StatusCode != http.StatusOK {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "")
			ctx.Abort()
			return
		}

		cr := &CheckAccessResp{}
		err = json.Unmarshal(b, cr)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "")
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
		return
	}
}
