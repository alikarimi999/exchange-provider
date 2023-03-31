package kucoin

import (
	"exchange-provider/internal/entity"
	"time"
)

type Token struct {
	Currency  string
	ChainName string
	Chain     string

	DepositAddress string
	DepositTag     string

	BlockTime     time.Duration
	ConfirmBlocks int64

	MinOrderSize float64
	MaxOrderSize float64

	MinWithdrawalSize float64
	MinWithdrawalFee  float64

	WithdrawalPrecision int
	OrderPrecision      int
}

func (k *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		Currency:  k.Currency,
		ChainName: k.ChainName,
		Chain:     k.Chain,

		DepositAddress: k.DepositAddress,
		DepositTag:     k.DepositTag,

		BlockTime:           k.BlockTime,
		ConfirmBlocks:       k.ConfirmBlocks,
		MinOrderSize:        k.MinOrderSize,
		MaxOrderSize:        k.MaxOrderSize,
		MinWithdrawalSize:   k.MinWithdrawalSize,
		MinWithdrawalFee:    k.MinWithdrawalFee,
		WithdrawalPrecision: k.WithdrawalPrecision,
		OrderPrecision:      k.OrderPrecision,
	}
}
