package http

import (
	"exchange-provider/internal/adapter/http/dto"
	udto "exchange-provider/internal/delivery/exchanges/dex/dto"
	"exchange-provider/internal/delivery/exchanges/dex/multichain"
	kdto "exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"
	"net/http"
	"sync"
)

func (s *Server) AddPairs(ctx Context) {
	nid := ctx.Param("id")
	ex, err := s.app.GetExchange(nid)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	res := &entity.AddPairsResult{}

	switch ex.Id() {
	case "kucoin":
		req := &dto.KucoinAddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
			return
		}

		if err := req.Validate(); err != nil {
			handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
			return
		}

		kps := &kdto.AddPairsRequest{}
		for _, p := range req.Pairs {
			kps.Pairs = append(kps.Pairs, p.Map())
		}

		res, err = s.app.AddPairs(ex, kps)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

	case "uniswapv3", "panckakeswapv2":
		req := &udto.AddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
			return
		}

		if err := req.Validate(); err != nil {
			handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
			return
		}

		res, err = s.app.AddPairs(ex, req)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

	case "multichain":
		req := &multichain.AddPairsRequest{}
		if err := ctx.Bind(req); err != nil {
			handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
			return
		}
		res, err = s.app.AddPairs(ex, req)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

	}

	ctx.JSON(200, dto.FromEntity(res))
}

func (s *Server) GetExchangesPairs(ctx Context) {
	req := &dto.GetAllPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp := &dto.GetAllPairsResponse{
		Exchanges: make(map[string]*dto.Exchange),
	}

	var exs []entity.Exchange
	if len(req.Es) == 0 || len(req.Es) == 1 && req.Es[0] == "*" {
		exs = s.app.AllExchanges()

	} else {
		for _, nid := range req.Es {
			ex, err := s.app.GetExchange(nid)
			if err != nil {
				resp.Messages = append(resp.Messages, err.Error())
				continue
			}

			exs = append(exs, ex)
		}
	}

	wg := &sync.WaitGroup{}
	for _, exc := range exs {
		wg.Add(1)
		go func(ex entity.Exchange) {
			defer wg.Done()
			ps, err := s.app.GetAllPairsByExchange(ex)
			if err != nil {
				resp.Messages = append(resp.Messages, err.Error())
				return
			}

			for _, p := range ps {
				dp := dto.PairDTO(p)
				resp.Exchanges[ex.Id()].Pairs = append(resp.Exchanges[ex.Id()].Pairs, dp)
			}
		}(exc)
	}
	wg.Wait()
	ctx.JSON(200, resp)
}

func (s *Server) RemovePair(ctx Context) {
	req := &dto.RemovePairRequest{}
	if err := ctx.Bind(req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	bc, qc, err := req.Parse()
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ex, err := s.app.GetExchange(req.Exchange)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	err = s.app.RemovePair(ex, bc, qc, req.Force)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(200, fmt.Sprintf("pair '%s/%s' removed from %s", bc, qc, ex.Id()))
}
