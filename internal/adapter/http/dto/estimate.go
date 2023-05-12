package dto

import "exchange-provider/internal/entity"

type EstimateAmountOutReq struct {
	Input    entity.TokenId `json:"input"`
	Output   entity.TokenId `json:"output"`
	AmountIn float64        `json:"amountIn"`
	LP       uint           `json:"lp"`
}

type EstimateAmountOutRes struct {
	Input             entity.TokenId `json:"input"`
	Output            entity.TokenId `json:"output"`
	AmountIn          Number         `json:"amountIn"`
	AmountOut         Number         `json:"amountOut"`
	FeeRate           Number         `json:"feeRate"`
	FeeAmount         Number         `json:"feeAmount"`
	ExchangeFee       Number         `json:"exchangeFee"`
	ExchangeFeeAmount Number         `json:"exchangeFeeAmount"`
	FeeCurrency       entity.TokenId `json:"feeCurrency"`
	LP                uint           `json:"lp"`
}
