package dto

import (
	bt "exchange-provider/internal/delivery/exchanges/cex/binance/types"
	kt "exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"strings"
)

type userSingleOrder struct {
	Status     string `json:"status"`
	FailReason string `json:"failReason,omitempty"`

	Input             entity.TokenId `json:"input"`
	Output            entity.TokenId `json:"output"`
	SetInAmount       Number         `json:"setInAmount"`
	EstimateAmountOut Number         `json:"estimateAmountOut"`

	InAmount  Number `json:"inAmount"`
	OutAmount Number `json:"outAmount"`
	Duration  string `json:"duration,omitempty"`

	FeeRate           Number         `json:"feeRate"`
	FeeRateAmount     Number         `json:"feeRateAmount"`
	ExchangeFee       Number         `json:"exchangeFee"`
	ExchangeFeeAmount Number         `json:"exchangeFeeAmount"`
	FeeCurrency       entity.TokenId `json:"feeCurrency"`

	Deposit    entity.Address `json:"deposit"`
	Refund     entity.Address `json:"refund"`
	Withdrawal entity.Address `json:"withdrawal"`

	DepositTxId    string `json:"depositTxId"`
	WithdrawalTxId string `json:"withdrawTxId"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	ExpireAt       int64  `json:"expireAt"`
}

func (s *userSingleOrder) fromEntity(ord entity.Order) *order {
	switch strings.Split(ord.ExchangeNid(), "-")[0] {
	case "kucoin":
		o := ord.(*kt.Order)
		so := userSingleOrder{
			Status:            o.STATUS().String(),
			Input:             o.In,
			Output:            o.Out,
			SetInAmount:       Number(o.SetAmountIn),
			EstimateAmountOut: Number(o.EstimateAmountOut),
			InAmount:          Number(o.Deposit.Amount),
			OutAmount:         Number(o.Withdrawal.Amount),
			FeeRate:           Number(o.FeeRate),
			FeeRateAmount:     Number(o.FeeAmount),
			ExchangeFee:       Number(o.ExchangeFee),
			ExchangeFeeAmount: Number(o.ExchangeFeeAmount),
			FeeCurrency:       o.FeeAndSpreadCurrency,

			Deposit:    o.Deposit.Address,
			Withdrawal: o.Withdrawal.Address,

			DepositTxId:    o.Deposit.TxId,
			WithdrawalTxId: o.Withdrawal.TxId,
			CreatedAt:      o.CreatedAT,
			UpdatedAt:      o.UpdatedAt,
			ExpireAt:       o.ExpireAt,
		}

		return &order{
			Id:        ord.ID().String(),
			Type:      singleStep,
			UserId:    o.UserID,
			CreatedAt: o.CreatedAT,
			Order:     so,
		}

	case "binance":
		o := ord.(*bt.Order)
		so := userSingleOrder{
			Status:            o.STATUS().String(),
			Input:             o.In,
			Output:            o.Out,
			SetInAmount:       Number(o.SetAmountIn),
			EstimateAmountOut: Number(o.EstimateAmountOut),
			InAmount:          Number(o.Deposit.Amount),
			OutAmount:         Number(o.Withdrawal.Amount),
			FeeRate:           Number(o.FeeRate),
			FeeRateAmount:     Number(o.FeeAmount),
			ExchangeFee:       Number(o.ExchangeFee),
			ExchangeFeeAmount: Number(o.ExchangeFeeAmount),
			FeeCurrency:       o.FeeAndSpreadCurrency,

			Deposit:    o.Deposit.Address,
			Withdrawal: o.Withdrawal.Address,

			DepositTxId:    o.Deposit.TxId,
			WithdrawalTxId: o.Withdrawal.TxId,
			CreatedAt:      o.CreatedAT,
			UpdatedAt:      o.UpdatedAt,
			ExpireAt:       o.ExpireAt,
		}

		return &order{
			Id:        ord.ID().String(),
			Type:      singleStep,
			UserId:    o.UserID,
			CreatedAt: o.CreatedAT,
			Order:     so,
		}

	}
	return nil
}
