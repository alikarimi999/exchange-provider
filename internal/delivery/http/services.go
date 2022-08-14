package http

import (
	"order_service/internal/delivery/services/deposite"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetServicesConfig(ctx *gin.Context) {
	res := struct {
		Auth interface{} `json:"auth_service"`
		Dep  interface{} `json:"deposit_service"`
	}{
		Auth: r.auth.Cofigs(),
		Dep:  r.srv.GetDepositServiceConfig(),
	}
	ctx.JSON(200, res)
}

func (r *Router) ChangeSerivcesConfig(ctx *gin.Context) {
	s := ctx.Param("service")

	switch s {
	case "auth":
		cfg := &authConf{}
		if err := ctx.BindJSON(cfg); err != nil {
			ctx.JSON(400, "invalid config")
			return
		}

		resp := struct {
			*authConf
			Message string `json:"message"`
		}{
			authConf: cfg,
		}

		if err := r.auth.changeConfigs(cfg); err != nil {
			resp.Message = err.Error()

		} else {
			resp.Message = "configs changed"
		}
		ctx.JSON(200, resp)

	case "deposit":
		cfg := &deposite.DepositServiceConfigs{}
		if err := ctx.BindJSON(cfg); err != nil {
			ctx.JSON(400, "invalid config")
			return
		}

		resp := struct {
			*deposite.DepositServiceConfigs
			Message string `json:"message"`
		}{
			DepositServiceConfigs: cfg,
		}

		if err := r.srv.ChangeDepositServiceConfig(cfg); err != nil {
			resp.Message = err.Error()
		} else {
			resp.Message = "configs changed"
		}
		ctx.JSON(200, resp)
	}

}
