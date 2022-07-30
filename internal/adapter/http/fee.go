package http

import (
	"fmt"
	"order_service/pkg/errors"
)

func (s *Server) ChangeFee(ctx Context) {
	req := struct {
		FeeRate float64 `json:"fee_rate"`
	}{}

	if err := ctx.Bind(&req); err != nil {
		handlerErr(ctx, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(err.Error())))
		return
	}

	s.app.ChangeFee(req.FeeRate)
	ctx.JSON(200, fmt.Sprintf("fee rate changed to %f", req.FeeRate))
}

func (s *Server) GetFee(ctx Context) {
	ctx.JSON(200, fmt.Sprintf("fee is %s", s.app.GetFee()))
}
