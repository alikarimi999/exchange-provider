package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/cex/swapspace"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/pkg/errors"
	"fmt"
	"strconv"
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
		ex, err := kucoin.NewKucoinExchange(cfg, s.v, s.l,
			false, s.repo, s.pc, s.fee)
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

	case "swapspace":
		cfg := &swapspace.Config{}
		if err := ctx.Bind(cfg); err != nil {
			ctx.JSON(nil, err)
			return
		}

		ex, err := swapspace.SwapSpace(cfg, s.repo, s.l)
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

		ex, err := evm.NewEvmDex(cfg, s.v, s.l, false)
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

	// case "multichain":

	// 	cfg := &multichain.Config{}
	// 	if err := ctx.Bind(cfg); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	if err := cfg.Validate(); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	ex, err := multichain.NewMultichain(cfg, s.app.WalletStore, s.v, s.l, false)
	// 	if err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	cfg.Name = "multichain"
	// 	if err := s.app.AddExchange(ex); err != nil {
	// 		ctx.JSON(nil, err)
	// 		return
	// 	}

	// 	cfg.Msg = "done"
	// 	ctx.JSON(cfg, nil)
	// 	return
	default:
		err := errors.Wrap(errors.ErrNotFound,
			errors.NewMesssage(fmt.Sprintf("exchange %s not exists", id)))
		ctx.JSON(nil, err)
		return
	}

}

func (s *Server) GetExchangeList(ctx Context) {

	res := &dto.GetAllExchangesResponse{
		Exchanges: make(map[uint]*dto.Account),
	}

	for _, ex := range s.app.AllExchanges() {
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
	sid := ctx.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	if err := s.app.RemoveExchange(uint(id), true); err != nil {
		ctx.JSON(nil, err)
		return
	}

	s.v.Set(sid, struct{}{})
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
