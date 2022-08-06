package http

import (
	"fmt"
	"order_service/internal/adapter/http/dto"
	"order_service/internal/app"
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

func (s *Server) AddPairs(ctx Context) {
	req := &dto.AddPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if err := req.Validate(); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	eps := map[string][]*entity.Pair{}
	for _, p := range req.Pairs {
		bc, err := p.BaseCoin()
		if err != nil {
			handlerErr(ctx, err)
			return
		}
		qc, err := p.QuoteCoin()
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		for ex, p := range p.ExchangePairs(bc, qc) {
			eps[ex] = append(eps[ex], p)
		}

	}

	resp := &dto.AddPairsResponse{
		Exchanges: make(map[string]*dto.AddPairsResult),
	}

	// add pairs to the exchange
	for nid, pairs := range eps {
		ex, err := s.app.GetExchange(nid)
		if err != nil {
			resp.Exchanges[nid] = &dto.AddPairsResult{
				Error: errors.ErrorMsg(err),
			}
			continue
		}
		if ex.CurrentStatus == app.ExchangeStatusDisable {
			resp.Exchanges[nid] = &dto.AddPairsResult{
				Error: fmt.Sprintf("exchange '%s' is %s", nid, ex.CurrentStatus),
			}
			continue
		}

		res, err := s.app.AddPairs(ex, pairs)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		resp.Exchanges[nid] = dto.FromEntity(res)

	}

	// add pair's coins to the supported coins

	ctx.JSON(200, resp)
}

func (s *Server) GetExchangesPairs(ctx Context) {
	req := &dto.GetAllPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	resp := &dto.GetAllPairsResponse{
		Exchanges: make(map[string]*dto.Exchange),
	}

	var exs []*app.Exchange
	if len(req.Exchanges) == 1 && req.Exchanges[0] == "*" {
		for _, ex := range s.app.AllExchanges() {
			if ex.CurrentStatus == app.ExchangeStatusDisable {
				continue
			}
			exs = append(exs, ex)
		}
	} else {
		for _, ex := range s.app.AllExchanges(req.Exchanges...) {
			fmt.Println(req.Exchanges)
			fmt.Println(ex.Name())
			if ex.CurrentStatus == app.ExchangeStatusDisable {
				resp.Messages = append(resp.Messages, fmt.Sprintf("exchange '%s' is %s", ex.NID(), ex.CurrentStatus))
				continue
			}
			exs = append(exs, ex)
		}
	}

	for _, ex := range exs {
		ps, err := s.app.GetAllPairs(ex)
		if err != nil {
			resp.Messages = append(resp.Messages, err.Error())
			continue
		}

		resp.Exchanges[ex.NID()] = &dto.Exchange{
			Status: ex.CurrentStatus,
		}
		for _, p := range ps {
			resp.Exchanges[ex.NID()].Pairs = append(resp.Exchanges[ex.NID()].Pairs, dto.PairDTO(p))
		}
	}

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
