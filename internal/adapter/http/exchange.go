package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	"exchange-provider/internal/delivery/exchanges/kucoin"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
)

func (s *Server) AddExchange(ctx Context) {
	id := ctx.Param("id")

	switch id {
	case "kucoin":
		cfg := &kucoin.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}
		ex, err := kucoin.NewKucoinExchange(cfg, s.pairs, s.rc, s.v, s.l, false)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

		if err := s.app.AddExchange(ex); err != nil {
			ctx.JSON(nil, err)
			return
		}
		cfg.Message = "done"
		ctx.JSON(cfg, nil)
		return

	case "dex":
		cfg := &evm.Config{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}

		ex, err := evm.NewEvmDex(cfg, s.pairs, s.v, s.l, false)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

		if err := s.app.AddExchange(ex); err != nil {
			ctx.JSON(nil, err)
			return
		}

		cfg.Message = "done"
		ctx.JSON(cfg, nil)
		return

	case "multichain":

		cfg := &multichain.Config{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}

		if err := cfg.Validate(); err != nil {
			ctx.JSON(nil, err)
			return
		}

		ex, err := multichain.NewMultichain(cfg, s.app.WalletStore, s.v, s.l, false)
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

		cfg.Name = "multichain"
		if err := s.app.AddExchange(ex); err != nil {
			ctx.JSON(nil, err)
			return
		}

		cfg.Msg = "done"
		ctx.JSON(cfg, nil)
		return
	default:
		err := errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("exchange %s not exists", id)))
		ctx.JSON(nil, err)
		return
	}

}

func (s *Server) GetExchangeList(ctx Context) {
	req := struct {
		Es []string `json:"exchanges"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
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
		err := errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage("set at least one exchange"))
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(res, nil)
}

func (s *Server) RemoveExchange(ctx Context) {
	id := ctx.Param("id")

	if err := s.app.RemoveExchange(id, true); err != nil {
		ctx.JSON(nil, err)
		return
	}

	s.v.Set(id, struct{}{})
	if err := s.v.WriteConfig(); err != nil {
		ctx.JSON(nil, err)
		return
	}

	ctx.JSON(
		struct {
			M string `json:"message"`
		}{
			M: "done",
		}, nil)
}
