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

	req.Msg = fmt.Sprintf("change was successful")
	ctx.JSON(200, req)
}

func (s *Server) SetDepositVol(ctx Context) {
	req := struct {
		UId int64  `json:"user_id"`
		OId int64  `json:"order_id"`
		DId int64  `json:"deposit_id"`
		Vol string `json:"volume"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	if req.UId == 0 || req.OId == 0 || req.DId == 0 || req.Vol == "" {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid request")))
		return
	}

	if err := s.app.SetDepositeVolume(req.UId, req.OId, req.DId, req.Vol); err != nil {
		handlerErr(ctx, err)
		return
	}

	ctx.JSON(200, fmt.Sprintf("deposit volume for %d/%d/%d changed to %s", req.UId, req.OId, req.DId, req.Vol))
}
