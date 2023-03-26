package kucoin

import (
	"exchange-provider/internal/entity"
	"fmt"
	"time"
)

type Token struct {
	TokenId string
	Network string

	Address string
	Tag     string

	BlockTime     time.Duration
	ConfirmBlocks int64

	MinOrderSize float64
	MaxOrderSize float64

	MinWithdrawalSize float64
	MinWithdrawalFee  float64

	WithdrawalPrecision int
	OrderPrecision      int

	NeedChain bool
}

func (k *Token) String() string {
	return fmt.Sprintf("%s-%s", k.TokenId, k.Network)
}

func (k *Token) Snapshot() entity.ExchangeToken {
	return &Token{
		TokenId:             k.TokenId,
		Network:             k.Network,
		Address:             k.Address,
		Tag:                 k.Tag,
		BlockTime:           k.BlockTime,
		ConfirmBlocks:       k.ConfirmBlocks,
		MinOrderSize:        k.MinOrderSize,
		MaxOrderSize:        k.MaxOrderSize,
		MinWithdrawalSize:   k.MinWithdrawalSize,
		MinWithdrawalFee:    k.MinWithdrawalFee,
		WithdrawalPrecision: k.WithdrawalPrecision,
		OrderPrecision:      k.OrderPrecision,
		NeedChain:           k.NeedChain,
	}
}
