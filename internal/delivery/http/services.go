package http

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) GetServicesConfig(ctx *gin.Context) {
	res := struct {
		Auth interface{} `json:"auth_service"`
	}{
		Auth: r.auth.Cofigs(),
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
	}
}
