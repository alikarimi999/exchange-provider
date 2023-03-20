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

	in := req.Input.ToEntity()
	out := req.Output.ToEntity()

	ex, amountOut, err := s.app.EstimateAmountOut(in, out, req.AmountIn, req.LP)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &dto.EstimateAmountOutRes{
		Input:     req.Input,
		Output:    req.Output,
		AmountIn:  req.AmountIn,
		AmountOut: amountOut,
		LP:        ex.Id(),
	}

	ctx.JSON(res, nil)
}
