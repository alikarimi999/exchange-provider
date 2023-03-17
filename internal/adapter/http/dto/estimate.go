package dto

type EstimateAmountOutReq struct {
	Input    string
	Output   string
	AmountIn float64 `json:"amountIn"`
	LP       uint    `json:"lp"`
}

type EstimateAmountOutRes struct {
	Input     string  `json:"input"`
	Output    string  `json:"output"`
	AmountIn  float64 `json:"amountIn"`
	AmountOut float64 `json:"amountOut"`
	LP        uint    `json:"lp"`
}
