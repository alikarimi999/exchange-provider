package types

import "exchange-provider/internal/entity"

type Deposit struct {
	TxId    string
	Amount  float64
	Address entity.Address
}
