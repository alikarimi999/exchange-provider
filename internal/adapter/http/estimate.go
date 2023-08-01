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
	ess, exs, err := s.app.EstimateAmountOut(req.Input, req.Output, req.AmountIn, req.LP, api.Level)
	if err != nil {
		ctx.JSON(nil, err)
		return
	}

	res := &dto.EstimateAmountOutRes{
		InUsd:  dto.Number(ess[0].InUsd),
		OutUsd: dto.Number(ess[0].OutUsd),
		EstimateAmount: dto.EstimateAmount{
			Input:             req.Input,
			Output:            req.Output,
			AmountIn:          dto.Number(ess[0].AmountIn),
			AmountOut:         dto.Number(ess[0].AmountOut),
			FeeRate:           dto.Number(ess[0].FeeRate),
			FeeAmount:         dto.Number(ess[0].FeeAmount),
			ExchangeFee:       dto.Number(ess[0].ExchangeFee),
			ExchangeFeeAmount: dto.Number(ess[0].ExchangeFeeAmount),
			FeeCurrency:       ess[0].FeeCurrency,
			LP:                exs[0].Id(),
		},
	}
	if len(ess) == 2 {
		res.ReverseEstimate = dto.EstimateAmount{
			Input:             req.Output,
			Output:            req.Input,
			AmountIn:          dto.Number(ess[1].AmountIn),
			AmountOut:         dto.Number(ess[1].AmountOut),
			FeeRate:           dto.Number(ess[1].FeeRate),
			FeeAmount:         dto.Number(ess[1].FeeAmount),
			ExchangeFee:       dto.Number(ess[1].ExchangeFee),
			ExchangeFeeAmount: dto.Number(ess[1].ExchangeFeeAmount),
			FeeCurrency:       ess[1].FeeCurrency,
			LP:                exs[1].Id(),
		}
	}
	ctx.JSON(res, nil)
}
