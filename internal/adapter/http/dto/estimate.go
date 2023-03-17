package dto

type EstimateAmountOutReq struct {
	In       string
	Out      string
	AmountIn float64 `json:"amountIn"`
	LP       uint    `json:"lp"`
}

type EstimateAmountOutRes struct {
	In        string
	Out       string
	AmountIn  float64 `json:"amountIn"`
	AmountOut float64 `json:"amountOut"`
	LP        uint    `json:"lp"`
}
