package http

import (
	"exchange-provider/internal/adapter/http/dto"
)

func (s *Server) EstimateAmountOut(ctx Context) {
	req := &dto.EstimateAmountOutReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	in, err := dto.ParseToken(req.In)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	out, err := dto.ParseToken(req.Out)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	ex, amountOut, err := s.app.EstimateAmountOut(in, out, req.AmountIn, req.LP)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &dto.EstimateAmountOutRes{
		In:        req.In,
		Out:       req.Out,
		AmountIn:  req.AmountIn,
		AmountOut: amountOut,
		LP:        ex.Id(),
	}

	ctx.JSON(res, nil)
}
