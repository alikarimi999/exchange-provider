package types

import "exchange-provider/internal/entity"

type Withdrawal struct {
	Id   string
	TxId string
	entity.Address
	Amount     float64
	BinanceFee float64
}
