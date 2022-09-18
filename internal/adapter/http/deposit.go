package http

import (
	"exchange-provider/internal/adapter/http/dto"
	"exchange-provider/pkg/errors"
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

func (s *Server) ChangeMinDeposit(ctx Context) {
	req := struct {
		Bc    string  `json:"base_coin"`
		MinBc float64 `json:"min_base_coin"`
		Qc    string  `json:"quote_coin"`
		MinQc float64 `json:"min_quote_coin"`
		Msg   string  `json:"message,omitempty"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if req.MinBc <= 0 || req.MinQc <= 0 {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("min deposit must be greater than 0")))
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

	req.Msg = "change was successful"
	ctx.JSON(200, req)
}
