package http

import (
	"fmt"
	"order_service/internal/adapter/http/dto"
	"order_service/pkg/errors"
)

func (s *Server) GetMinPairDeposit(ctx Context) {
	req := struct {
		Bc string `json:"base_coin"`
		Qc string `json:"quote_coin"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	bc, err := dto.ParseCoin(req.Bc)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	qc, err := dto.ParseCoin(req.Qc)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	minBc, minQc := s.app.GetMinPairDeposit(bc, qc)

	ctx.JSON(200, struct {
		MinBc float64 `json:"min_base_coin"`
		MinQc float64 `json:"min_quote_coin"`
	}{
		MinBc: minBc,
		MinQc: minQc,
	})

}

func (s *Server) GetAllMinDeposit(ctx Context) {
	ctx.JSON(200, s.app.AllMinDeposit())
}

func (s *Server) ChangeDefaultMinDeposit(ctx Context) {
	req := struct {
		D float64 `json:"default_min_deposit"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if err := s.app.ChangeDefaultMinDeposit(req.D); err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(200, struct {
		D float64 `json:"default_min_deposit"`
	}{
		D: req.D,
	})
}

func (s *Server) GetDefaultMinDeposit(ctx Context) {
	ctx.JSON(200, struct {
		D float64 `json:"default_min_deposit"`
	}{
		D: s.app.GetDefaultMinDeposit(),
	})
}

func (s *Server) ChangeMinDeposit(ctx Context) {
	req := struct {
		Bc    string  `json:"base_coin"`
		MinBc float64 `json:"min_base_coin"`
		Qc    string  `json:"quote_coin"`
		MinQc float64 `json:"min_quote_coin"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	bc, err := dto.ParseCoin(req.Bc)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	qc, err := dto.ParseCoin(req.Qc)
	if err != nil {
		handlerErr(ctx, err)
		return
	}

	if err := s.app.ChangeMinDeposit(bc, qc, req.MinBc, req.MinQc); err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(200, fmt.Sprintf("min deposit for %s/%s changed to %f/%f", bc.String(), qc.String(), req.MinBc, req.MinQc))
}
