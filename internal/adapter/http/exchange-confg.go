package http

import (
	"exchange-provider/internal/delivery/exchanges/cex/binance"
	"exchange-provider/internal/delivery/exchanges/cex/kucoin"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge"
	"exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"
)

func (s *Server) UpdateConfig(ctx Context) {
	sid := ctx.Param("id")
	var (
		cfg interface{}
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

	ex, err := s.app.GetExchange(uint(id))
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	switch ex.Type() {
	case entity.CEX:
		switch ex.Name() {
		case "kucoin":
			cfg = &kucoin.Config{}
			err = ctx.Bind(cfg)
		case "binance":
			cfg = &binance.Config{}
			err = ctx.Bind(cfg)
		}
	case entity.CrossDex:
		switch ex.Name() {
		case "allbridge":
			cfg = &allbridge.Config{}
			err = ctx.Bind(cfg)
		}
	case entity.EvmDEX:
		cfg = &evm.Config{}
		err = ctx.Bind(cfg)
	}

	if err != nil {
		ctx.JSON(nil, errors.Wrap(errors.ErrBadRequest, err))
		return
	}
	if err := ex.UpdateConfigs(cfg, s.exs); err != nil {
		ctx.JSON(nil, err)
		return
	}
	ctx.JSON(cfg, nil)
}
