package http

import (
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateChains(ctx *gin.Context) {
	req := &multichain.UpdateChainsReq{Chains: make(map[multichain.ChainId][]string)}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ex, err := s.app.GetExchange("multichain")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	cmd := make(map[string]interface{})
	cmd[multichain.UpdateChains] = req
	res, err := ex.Command(cmd)
	if err != nil {
		ctx.JSON(200, err.Error())
		return
	}

	ctx.JSON(200, res)
}
