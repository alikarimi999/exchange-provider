package http

import (
	"fmt"
	"net/http"
	"order_service/internal/delivery/exchanges/kucoin"
)

func (s *Server) AddExchange(ctx Context) {
	id := ctx.Param("id")

	if s.app.ExchangeExists(id) {
		ex, err := s.app.GetExchange(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("exchange %s with configs ( %s ) already exists", id, ex.Configs()))
		return
	}

	switch id {
	case "kucoin":
		cfg := &kucoin.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ex := kucoin.NewKucoinExchange()

		if err := s.app.AddExchange(ex, cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s with configs ( %s ) added", ex.ID(), cfg))
		return
	default:
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("exchange %s not supported", id))
		return
	}

}

func (s *Server) ChangeExchangeAccount(ctx Context) {
	const agent = "http.Server.ChangeExchangeAccount"
	id := ctx.Param("id")

	if !s.app.ExchangeExists(id) {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("exchange %s not exists", id))
		return
	}

	switch id {
	case "kucoin":
		cfg := &kucoin.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		s.l.Info(agent, fmt.Sprintf("change exchange %s account with configs %s", id, cfg))

		if err := s.app.ChangeExchangeAccount(id, cfg); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	ex, err := s.app.GetExchange(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s with configs ( %s ) changed", ex.ID(), ex.Configs()))
}
