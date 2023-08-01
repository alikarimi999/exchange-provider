package types

import "fmt"

type Number float64

func (n Number) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.*f", 8, n)
	return []byte(s), nil
}

type StepStatus struct {
	Status     string `json:"status"`
	TxId       string `json:"txId"`
	AmountIn   Number `json:"amountIn"`
	AmountOut  Number `json:"amountOut"`
	FailedDesc string `json:"failedDesc"`
}
type OrderStatus map[int]*StepStatus
