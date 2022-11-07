package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/internal/app"
	udto "exchange-provider/internal/delivery/exchanges/dex/dto"
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
	if ex.CurrentStatus == app.ExchangeStatusDisable {
		ctx.JSON(404, "exchange is disable")
		return
	}

	res := &entity.AddPairsResult{}

	switch ex.Name() {
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

	var exs []*app.Exchange
	if len(req.Es) == 0 || len(req.Es) == 1 && req.Es[0] == "*" {
		for _, ex := range s.app.AllExchanges() {
			if ex.CurrentStatus == app.ExchangeStatusDisable {
				continue
			}
			exs = append(exs, ex)
		}
	} else {
		for _, nid := range req.Es {
			ex, err := s.app.GetExchange(nid)
			if err != nil {
				resp.Messages = append(resp.Messages, err.Error())
				continue
			}
			if ex.CurrentStatus == app.ExchangeStatusDisable {
				resp.Messages = append(resp.Messages, fmt.Sprintf("exchange %s is %s", ex.NID(), ex.CurrentStatus))
				continue
			}
			exs = append(exs, ex)
		}
	}

	wg := &sync.WaitGroup{}
	for _, exc := range exs {
		wg.Add(1)
		go func(ex *app.Exchange) {
			defer wg.Done()
			ps, err := s.app.GetAllPairsByExchange(ex)
			if err != nil {
				resp.Messages = append(resp.Messages, err.Error())
				return
			}

			resp.Exchanges[ex.NID()] = &dto.Exchange{
				Status: ex.CurrentStatus,
			}
			for _, p := range ps {
				dp := dto.PairDTO(p)
				resp.Exchanges[ex.NID()].Pairs = append(resp.Exchanges[ex.NID()].Pairs, dp)
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
	if ex.CurrentStatus == app.ExchangeStatusDisable {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("exchange %s is %s", ex.NID(), ex.CurrentStatus))))
		return
	}

	err = s.app.RemovePair(ex.Exchange, bc, qc, req.Force)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(200, fmt.Sprintf("pair '%s/%s' removed from %s", bc, qc, ex.NID()))
}
