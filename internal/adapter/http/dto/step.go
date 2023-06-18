package dto

import (
	bt "exchange-provider/internal/delivery/exchanges/cex/binance/types"
	kt "exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strings"
)

type OrderStep struct {
	OrderId     string `json:"orderId"`
	Type        string `json:"type,omitempty"`
	TotalSteps  int    `json:"totalSteps"`
	CurrentStep int    `json:"currentStep"`
}

type SingleStep struct {
	*OrderStep
	Duration string         `json:"duration"`
	Token    entity.TokenId `json:"token"`
	Address  string         `json:"address"`
	Tag      string         `json:"tag"`

	AmountIn                  Number         `json:"amountIn"`
	EstimateAmountOut         Number         `json:"estimateAmountOut"`
	FeeRate                   Number         `json:"feeRate"`
	EstimateFeeAmount         Number         `json:"estimateFeeAmount"`
	ExchangeFee               Number         `json:"exchangeFee"`
	EstimateExchangeFeeAmount Number         `json:"estimateExchangeFeeAmount"`
	FeeCurrency               entity.TokenId `json:"feeCurrency"`
	LP                        uint           `json:"lp"`
	CreatedAt                 int64          `json:"createdAt"`
	UpdatedAt                 int64          `json:"updatedAt"`
	ExpireAt                  int64          `json:"expireAt"`
}

func SingleStepResponse(ord entity.Order) (*SingleStep, error) {
	switch strings.Split(ord.ExchangeNid(), "-")[0] {
	case "kucoin":
		o := ord.(*kt.Order)
		return &SingleStep{
			OrderStep: &OrderStep{OrderId: o.ObjectId.String(), CurrentStep: 1, TotalSteps: 1},
			Token:     o.In,
			Address:   o.Deposit.Address.Addr,
			Tag:       o.Deposit.Address.Tag,

			AmountIn:                  Number(o.SetAmountIn),
			EstimateAmountOut:         Number(o.EstimateAmountOut),
			FeeRate:                   Number(o.FeeRate),
			EstimateFeeAmount:         Number(o.EstimateFeeAmount),
			ExchangeFee:               Number(o.ExchangeFee),
			EstimateExchangeFeeAmount: Number(o.EstimateExchangeFeeAmount),
			FeeCurrency:               o.FeeAndSpreadCurrency,
			LP:                        o.ExLp,

			CreatedAt: o.CreatedAT,
			UpdatedAt: o.UpdatedAt,
			ExpireAt:  o.ExpireAt,
		}, nil

	case "binance":
		o := ord.(*bt.Order)
		return &SingleStep{
			OrderStep: &OrderStep{OrderId: o.ObjectId.String(), CurrentStep: 1, TotalSteps: 1},
			Token:     o.In,
			Address:   o.Deposit.Address.Addr,
			Tag:       o.Deposit.Address.Tag,

			AmountIn:                  Number(o.SetAmountIn),
			EstimateAmountOut:         Number(o.EstimateAmountOut),
			FeeRate:                   Number(o.FeeRate),
			EstimateFeeAmount:         Number(o.EstimateFeeAmount),
			ExchangeFee:               Number(o.ExchangeFee),
			EstimateExchangeFeeAmount: Number(o.EstimateExchangeFeeAmount),
			FeeCurrency:               o.FeeAndSpreadCurrency,
			LP:                        o.ExLp,

			CreatedAt: o.CreatedAT,
			UpdatedAt: o.UpdatedAt,
			ExpireAt:  o.ExpireAt,
		}, nil
	}
	return nil, errors.Wrap(errors.ErrInternal)
}

type multiStep struct {
	*OrderStep

	AmountIn                  Number         `json:"amountIn"`
	EstimateAmountOut         Number         `json:"estimateAmountOut"`
	FeeRate                   Number         `json:"feeRate"`
	EstimateFeeAmount         Number         `json:"estimateFeeAmount"`
	ExchangeFee               Number         `json:"exchangeFee"`
	EstimateExchangeFeeAmount Number         `json:"estimateExchangeFeeAmount"`
	FeeCurrency               entity.TokenId `json:"feeCurrency"`

	Transaction interface{} `json:"transaction"`

	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
	ExpireAt  int64 `json:"expireAt"`
}

func MultiStep(o entity.Order, tx entity.Tx, step, steps int) *multiStep {
	ms := &multiStep{
		OrderStep: &OrderStep{OrderId: o.ID().String(), CurrentStep: step, TotalSteps: steps},
	}
	switch tx.Type() {
	case entity.Evm:
		eo := o.(*types.Order)
		ms.AmountIn = Number(eo.AmountIn)
		ms.EstimateAmountOut = Number(eo.EstimateAmountOut)
		ms.FeeRate = Number(eo.FeeRate)
		ms.EstimateFeeAmount = Number(eo.EstimateFeeAmount)
		ms.ExchangeFee = Number(eo.ExchangeFee)
		ms.EstimateExchangeFeeAmount = Number(eo.EstimateExchangeFeeAmount)
		ms.FeeCurrency = eo.FeeCurrency
		ms.Type = string(tx.Type())
		ms.Transaction = evmTx(tx, eo.Sender.Hex())
		ms.CreatedAt = eo.CreatedAT
		ms.UpdatedAt = eo.UpdatedAt
		ms.ExpireAt = eo.ExpireAt

	}
	return ms
}
