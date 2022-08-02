package http

import (
	"fmt"
	"order_service/internal/adapter/http/dto"
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
	for id, pairs := range eps {
		ex, err := s.app.GetExchange(id)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Exchanges[id] = &dto.AddPairsResult{
					Error: fmt.Sprintf("exchange '%s' not found", id),
				}
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}

		res, err := s.app.AddPairs(ex, pairs)
		if err != nil {
			handlerErr(ctx, err)
			return
		}

		resp.Exchanges[id] = dto.FromEntity(res)

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
		Exchanges: make(map[string][]*dto.Pair),
	}

	var exs []string
	if len(req.Exchanges) == 1 && req.Exchanges[0] == "*" {
		exs = s.app.GetAllExchanges()
	} else {
		exs = req.Exchanges
	}

	for _, ex := range exs {
		ps, err := s.app.GetAllPairs(ex)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Messages = append(resp.Messages, fmt.Sprintf("exchange '%s' not found", ex))
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}

		for _, p := range ps {
			resp.Exchanges[ex] = append(resp.Exchanges[ex], dto.PairDTO(p))
		}
	}

	ctx.JSON(200, resp)
}

func (s *Server) GetPair(ctx Context) {
	req := &dto.GetPairRequest{}

	if err := ctx.Bind(req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	bc, qc, err := req.Parse()
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	exs := []string{}
	if len(req.Exchanges) == 1 && req.Exchanges[0] == "*" {
		exs = s.app.GetAllExchanges()
	} else {
		exs = req.Exchanges
	}

	resp := &dto.GetPairResponse{
		Exchanges: make(map[string]*dto.Pair),
	}
	for _, exc := range exs {
		ex, err := s.app.GetExchange(exc)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Messages = append(resp.Messages, fmt.Sprintf("exchange '%s' not found", exc))
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}
		p, err := s.app.GetPair(ex, bc, qc)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Messages = append(resp.Messages, fmt.Sprintf("pair '%s/%s' not found in %s", bc, qc, exc))
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}
		resp.Exchanges[exc] = dto.PairDTO(p)

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

	resp := &dto.RemovePairResponse{
		Exchanges: make(map[string]string),
	}
	for _, exc := range req.Exchanges {
		ex, err := s.app.GetExchange(exc)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Exchanges[exc] = fmt.Sprintf("exchange '%s' not found", exc)
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}
		err = s.app.RemovePair(ex, bc, qc)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Exchanges[exc] = fmt.Sprintf("pair '%s/%s' not found in %s", bc, qc, exc)
				continue
			default:
				handlerErr(ctx, err)
				return
			}
		}
		resp.Exchanges[exc] = fmt.Sprintf("pair '%s/%s' removed from %s", bc, qc, exc)
	}

	ctx.JSON(200, resp)
}
