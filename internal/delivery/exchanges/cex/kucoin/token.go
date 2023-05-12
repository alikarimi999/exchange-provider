package kucoin

import (
	"exchange-provider/internal/entity"
	"time"
)

type Token struct {
	Currency  string `json:"currency"`
	ChainName string `json:"chanName,omitempty"  bson:"chainName,omitempty"`
	Chain     string `json:"chain,omitempty" bson:"chain,omitempty"`

	DepositAddress string `json:"depositAddress,omitempty" bson:"-"`
	DepositTag     string `json:"depositTag,omitempty" bson:"-"`

	BlockTime     time.Duration `json:"blockTime,omitempty" bson:"blockTime,omitempty"`
	ConfirmBlocks int64         `json:"confirmBlocks,omitempty" bson:"-"`

	MinOrderSize float64 `json:"minOrderSize" bson:"-"`
	MaxOrderSize float64 `json:"maxOrderSize" bson:"-"`

	MinWithdrawalSize float64 `json:"minWithdrawalSize,omitempty" bson:"-"`
	MinWithdrawalFee  float64 `json:"minWithdrawalFee,omitempty" bson:"-"`

	WithdrawalPrecision int `json:"withdrawalPrecision,omitempty" bson:"-"`
	OrderPrecision      int `json:"orderPrecision" bson:"-"`
}

func (k *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		Currency:  k.Currency,
		ChainName: k.ChainName,
		Chain:     k.Chain,

		DepositAddress: k.DepositAddress,
		DepositTag:     k.DepositTag,

		BlockTime:     k.BlockTime,
		ConfirmBlocks: k.ConfirmBlocks,

		MinOrderSize: k.MinOrderSize,
		MaxOrderSize: k.MaxOrderSize,

		MinWithdrawalSize: k.MinWithdrawalSize,
		MinWithdrawalFee:  k.MinWithdrawalFee,

		WithdrawalPrecision: k.WithdrawalPrecision,
		OrderPrecision:      k.OrderPrecision,
	}
}

func (k *Token) snapshot() *Token {
	return &Token{
		Currency:  k.Currency,
		ChainName: k.ChainName,
		Chain:     k.Chain,

		DepositAddress: k.DepositAddress,
		DepositTag:     k.DepositTag,

		BlockTime:     k.BlockTime,
		ConfirmBlocks: k.ConfirmBlocks,

		MinOrderSize: k.MinOrderSize,
		MaxOrderSize: k.MaxOrderSize,

		MinWithdrawalSize: k.MinWithdrawalSize,
		MinWithdrawalFee:  k.MinWithdrawalFee,

		WithdrawalPrecision: k.WithdrawalPrecision,
		OrderPrecision:      k.OrderPrecision,
	}
}
