package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/delivery/exchanges/cex/binance"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
)

func (s *Server) AddExchange(ctx Context) {
	name := ctx.Param("name")

	switch name {
	case "kucoin":
		cfg := &kucoin.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}
		if s.app.ExchangeExists(cfg.Id) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchange already exists")))
			return
		}

		ex, err := kucoin.NewExchange(cfg, s.pairs, s.l, false, s.repo, s.fee, s.spread)
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

	case "binance":
		cfg := &binance.Configs{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}
		if s.app.ExchangeExists(cfg.Id) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchange already exists")))
			return
		}

		ex, err := binance.NewExchange(cfg, s.repo, s.pairs, s.spread, s.l, false)
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

	// case "swapspace":
	// 	cfg := &swapspace.Config{}
	// 	if err := ctx.Bind(cfg); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	if s.app.ExchangeExists(cfg.Id) {
	// 		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchange already exists")))
	// 		return
	// 	}

	// 	ex, err := swapspace.SwapSpace(cfg, s.repo, s.pairs, s.l)
	// 	if err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	if err := s.app.AddExchange(ex); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}
	// 	cfg.Message = "done"
	// 	ctx.JSON(cfg, nil)
	// 	return

	case "evmdex":
		cfg := &evm.Config{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}

		if s.app.ExchangeExists(cfg.Id) {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("exchange already exists")))
			return
		}

		ex, err := evm.NewEvmDex(cfg, s.repo, s.pairs, s.l)
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
	default:
		err := errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("exchange %s not exists", name)))
		ctx.JSON(nil, err)
		return
	}

}

func (s *Server) GetExchangeList(ctx Context) {
	sid := ctx.Param("id")
	var (
		id  int
		err error
	)
	if sid != "" {
		id, err = strconv.Atoi(sid)
		if err != nil || id == 0 {
			ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
				errors.NewMesssage("id must be a number greater than zero")))
			return
		}
	}

	res := &dto.GetAllExchangesResponse{}
	if id > 0 {
		ex, err := s.app.GetExchange(uint(id))
		if err != nil {
			ctx.JSON(nil, err)
			return
		}

		res.Exchanges = append(res.Exchanges, dto.Exchange{
			Type: string(ex.Type()),
			Conf: ex.Configs(),
		})
	} else {
		for _, ex := range s.app.AllExchanges() {
			res.Exchanges = append(res.Exchanges, dto.Exchange{
				Type: string(ex.Type()),
				Conf: ex.Configs(),
			})
		}
	}
	ctx.JSON(res, nil)
}

func (s *Server) CommandExchanges(ctx Context) {
	req := dto.LpsRequest{}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(nil, err)
		return
	}
	var enable bool
	res := dto.CmdResp{}
	switch req.Cmd {
	case "remove":
		if req.All {
			if err := s.app.RemoveExchange(0, true); err != nil {
				ctx.JSON(nil, err)
				return
			}
			ctx.JSON(struct {
				Msg string "json:\"msg\""
			}{Msg: "done"}, nil)
			return
		}

		for _, lp := range req.Lps {
			resp := struct {
				Lp  uint   "json:\"lp\""
				Msg string "json:\"msg\""
			}{Lp: lp}
			err := s.app.RemoveExchange(lp, false)
			if err != nil {
				resp.Msg = err.Error()
				res.LpsRes = append(res.LpsRes, resp)
				continue
			}
			resp.Msg = "done"
			res.LpsRes = append(res.LpsRes, resp)
		}
		ctx.JSON(res, nil)
		return
	case "enable":
		enable = true
	case "disable":
		enable = false
	default:
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest,
			fmt.Errorf("cmd '%s' is not supported", req.Cmd)))
		return
	}

	if req.All {
		if err := s.app.EnableDisable(0, enable, true); err != nil {
			ctx.JSON(nil, err)
			return
		}
		ctx.JSON(struct {
			Msg string "json:\"msg\""
		}{Msg: "done"}, nil)
		return
	}

	for _, lp := range req.Lps {
		resp := struct {
			Lp  uint   "json:\"lp\""
			Msg string "json:\"msg\""
		}{Lp: lp}
		err := s.app.EnableDisable(lp, enable, false)
		if err != nil {
			resp.Msg = err.Error()
			res.LpsRes = append(res.LpsRes, resp)
			continue
		}
		resp.Msg = "done"
		res.LpsRes = append(res.LpsRes, resp)
	}
	ctx.JSON(res, nil)
}
