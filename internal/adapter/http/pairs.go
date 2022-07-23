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
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, "invalid request"))
		return
	}

	if err := req.Validate(); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	eps := map[string][]*entity.ExchangePair{}
	coins := make(map[string]*entity.Coin) // map[coinId+chainId]*entity.Coin
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

		coins[bc.Id+bc.Chain.Id] = bc
		coins[qc.Id+qc.Chain.Id] = qc

		for ex, p := range p.ExchangePairs(bc, qc) {
			eps[ex] = append(eps[ex], p)
		}

	}

	// add pairs to the exchange
	for ex, pairs := range eps {
		s.app.AddPairs(ex, pairs)
	}

	// add pair's coins to the supported coins
	s.app.AddCoins(coins)

	ctx.JSON(200, "done")
}

func (s *Server) GetExchangesPairs(ctx Context) {
	req := &dto.GetExchangesPairsRequest{}
	if err := ctx.Bind(req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	resp := &dto.GetExchangesPairsResponse{
		Exchanges: make(map[string][]*dto.GetPair),
	}
	for _, ex := range req.Exchanges {
		ps, err := s.app.SupportedPairs(ex)
		if err != nil {
			switch errors.ErrorCode(err) {
			case errors.ErrNotFound:
				resp.Messages = append(resp.Messages, fmt.Sprintf("exchange '%s' not found", ex))
			default:
				handlerErr(ctx, err)
				return
			}
		}
		for _, p := range ps {
			resp.Exchanges[ex] = append(resp.Exchanges[ex], dto.ToDTO(p))
		}
	}

	ctx.JSON(200, resp)
}
