package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	uniswapv3 "exchange-provider/internal/delivery/exchanges/uniswap/v3"
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
		ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s added", ex.NID()))
		return

	case "uniswapv3":

		cfg := &dto.Config{}

		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if len(cfg.Providers) == 0 {
			ctx.JSON(http.StatusBadRequest, "at least one provider must be specified")
			return
		}

		conf := cfg.Map()
		conf.Name = id
		ex, err := uniswapv3.NewExchange(conf, s.rc, s.v, s.l, false)
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

		cfg.Id = ex.NID()
		cfg.Accounts = conf.Accounts
		cfg.Msg = "exchange added"
		ctx.JSON(http.StatusOK, cfg)
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
	exs := []*app.Exchange{}
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
		res.Exchanges[ex.NID()] = &dto.Account{
			Status: string(ex.CurrentStatus),
			Conf:   ex.Configs(),
		}
	}

	if len(res.Exchanges) == 0 {
		ctx.JSON(http.StatusNotFound, "no exchange found")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (s *Server) ChangeStatus(ctx Context) {
	req := struct {
		Id     string `json:"id"`
		Status string `json:"status"`
		Force  bool   `json:"force"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if req.Id == "" || req.Status == "" {
		ctx.JSON(http.StatusBadRequest, "id and status are required")
		return
	}

	switch req.Status {
	case app.ExchangeStatusActive, app.ExchangeStatusDeactive, app.ExchangeStatusDisable:
		res, err := s.app.ChangeExchangeStatus(req.Id, req.Status, req.Force)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		r := &dto.ChangeExchangeStatusResponse{}
		r.FromEntity(res)
		ctx.JSON(http.StatusOK, r)

	case "remove":
		if err := s.app.RemoveExchange(req.Id, req.Force); err != nil {
			handlerErr(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s removed!", req.Id))
	default:
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("status %s not supported", req.Status))
	}

}
