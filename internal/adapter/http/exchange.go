package http

import (
	"fmt"
	"net/http"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
	"order_service/internal/delivery/exchanges/kucoin"
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
		ex, err := kucoin.NewKucoinExchange(cfg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := s.app.AddExchange(ex); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, fmt.Sprintf("exchange %s with accountId %s added!", id, ex.NID()))
		return
	default:
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("exchange %s not supported", id))
		return
	}

}

func (s *Server) GetExchangeList(ctx Context) {
	req := struct {
		Es []string `json:"exchange_names"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	exs := []*app.Exchange{}
	if len(req.Es) == 1 && req.Es[0] == "*" {
		exs = s.app.AllExchanges()
	} else {
		exs = s.app.AllExchanges(req.Es...)
	}

	res := &dto.GetAllExchangesResponse{
		Exchanges: make(map[string]*dto.Account),
	}

	for _, ex := range exs {
		res.Exchanges[ex.NID()] = &dto.Account{
			Status: string(ex.CurrentStatus),
			Conf:   ex.Configs(),
		}
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
	case app.ExchangeStatusActive, app.ExchangeStatusDeactive, app.ExchangeStatusDisabled:
		res, err := s.app.ChangeExchangeStatus(req.Id, req.Status, req.Force)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		r := &dto.ChangeExchangeStatusResponse{}
		r.FromEntity(res)
		ctx.JSON(http.StatusOK, r)

	default:
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("status %s not supported", req.Status))
	}

}
