package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/delivery/exchanges/dex"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	"exchange-provider/internal/entity"
	"fmt"
	"net/http"
)

func (s *Server) AddExchange(ctx Context) {
	id := ctx.Param("id")

	switch id {
	case "kucoin":
		cfg := &kucoin.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ex, err := kucoin.NewKucoinExchange(cfg, s.rc, s.v, s.l, false)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		if err := s.app.AddExchange(ex); err != nil {
			handlerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s added", ex.Id()))
		return

	case "dex":

		cfg := &dto.Config{}

		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if len(cfg.Providers) == 0 {
			ctx.JSON(http.StatusBadRequest, "at least one provider must be specified")
			return
		}

		conf, err := cfg.Map()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ex, err := dex.NewDEX(conf, s.app.WalletStore, s.rc, s.v, s.l, false)
		if err != nil {
			cfg.Msg = err.Error()
			ctx.JSON(http.StatusOK, cfg)
			return
		}

		if err := s.app.AddExchange(ex); err != nil {
			cfg.Msg = err.Error()
			ctx.JSON(http.StatusOK, cfg)
			return
		}

		cfg.Id = ex.Id()
		cfg.Accounts = conf.Accounts
		cfg.Msg = "exchange added"
		ctx.JSON(http.StatusOK, cfg)
		return

	case "multichain":

		cfg := &multichain.Config{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := cfg.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ex, err := multichain.NewMultichain(cfg, s.app.WalletStore, s.v, s.l, false)
		if err != nil {
			ctx.JSON(200, err.Error())
			return
		}

		cfg.Name = "multichain"
		if err := s.app.AddExchange(ex); err != nil {
			ctx.JSON(200, err.Error())
			return
		}

		cfg.Msg = "exchange added"
		ctx.JSON(200, cfg)
		return
	default:
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("exchange %s not supported", id))
		return
	}

}

func (s *Server) GetExchangeList(ctx Context) {
	req := struct {
		Es []string `json:"exchanges"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := &dto.GetAllExchangesResponse{
		Exchanges: make(map[string]*dto.Account),
	}
	exs := []entity.Exchange{}
	if len(req.Es) == 0 || len(req.Es) == 1 && req.Es[0] == "*" {
		exs = s.app.AllExchanges()
	} else {
		for _, e := range req.Es {
			ex, err := s.app.GetExchange(e)
			if err != nil {
				res.Msgs = append(res.Msgs, err.Error())
				continue
			}

			exs = append(exs, ex)
		}
	}

	for _, ex := range exs {
		res.Exchanges[ex.Id()] = &dto.Account{
			Conf: ex.Configs(),
		}
	}

	if len(res.Exchanges) == 0 {
		ctx.JSON(http.StatusNotFound, "no exchange found")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) RemoveExchange(ctx Context) {
	id := ctx.Param("id")

	if err := s.app.RemoveExchange(id, true); err != nil {
		handlerErr(ctx, err)
		return
	}

	s.v.Set(id, struct{}{})
	if err := s.v.WriteConfig(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s removed!", id))

}
