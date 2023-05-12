package http

import (
	"exchange-provider/internal/adapter/http/dto"
)

func (s *Server) EstimateAmountOut(ctx Context) {
	api := ctx.GetApi()
	req := &dto.EstimateAmountOutReq{}
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(nil, err)
		return
	}

	req.Input.ToUpper()
	req.Output.ToUpper()
	es, ex, err := s.app.EstimateAmountOut(req.Input, req.Output, req.AmountIn, req.LP, api.Level)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &dto.EstimateAmountOutRes{
		Input:             req.Input,
		Output:            req.Output,
		AmountIn:          dto.Number(req.AmountIn),
		AmountOut:         dto.Number(es.AmountOut),
		FeeRate:           dto.Number(es.FeeRate),
		FeeAmount:         dto.Number(es.FeeAmount),
		ExchangeFee:       dto.Number(es.ExchangeFee),
		ExchangeFeeAmount: dto.Number(es.ExchangeFeeAmount),
		FeeCurrency:       es.FeeCurrency,
		LP:                ex.Id(),
	}

	ctx.JSON(res, nil)
}
