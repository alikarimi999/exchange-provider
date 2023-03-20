package dto

type EstimateAmountOutReq struct {
	Input    Token
	Output   Token
	AmountIn float64 `json:"amountIn"`
	LP       uint    `json:"lp"`
}

type EstimateAmountOutRes struct {
	Input     Token   `json:"input"`
	Output    Token   `json:"output"`
	AmountIn  float64 `json:"amountIn"`
	AmountOut float64 `json:"amountOut"`
	LP        uint    `json:"lp"`
}
