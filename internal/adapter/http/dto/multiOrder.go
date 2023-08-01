package dto

import (
	at "exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	et "exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"strings"

	"exchange-provider/internal/entity"
)

// type evmStep struct {
// 	*route
// 	IsApprove bool
// 	Approved  bool
// }
// type route struct {
// 	Input    string
// 	Output   string
// 	Exchange string
// }

// type adminMultiOrder struct {
// 	Steps    map[uint]*evmStep
// 	Sender   string
// 	Receiver string
// 	AmountIn float64 `bson:"amountIn"`
// 	FeeRate  float64 `bson:"feeRate"`
// }

// func (a *adminMultiOrder) fromEntity(o *entity.DexOrder) *order {
// 	a = &adminMultiOrder{
// 		Steps:    make(map[uint]*evmStep),
// 		Sender:   o.Sender.Hex(),
// 		Receiver: o.Receiver.Hex(),
// 		AmountIn: a.AmountIn,
// 		FeeRate:  a.FeeRate,
// 	}
// 	for i, s := range o.Steps {
// 		a.Steps[i] = &evmStep{
// 			route:     &route{Input: s.In.String(), Output: s.Out.String(), Exchange: s.Exchange},
// 			IsApprove: s.NeedApprove,
// 			Approved:  s.Approved,
// 		}
// 	}
// 	return &order{
// 		Id:        o.ObjectId.String(),
// 		Type:      multiSteps,
// 		UserId:    o.UserId,
// 		CreatedAt: o.CreatedAt,
// 		Order:     a,
// 	}
// }

type userMultiOrder struct {
	TotalStpes        uint           `json:"totalStpes"`
	Input             entity.TokenId `json:"input"`
	Output            entity.TokenId `json:"output"`
	AmountIn          Number         `json:"amountIn"`
	AmountOut         Number         `json:"amountOut"`
	EstimateAmountOut Number         `json:"estimateAmountOut"`
	FeeRate           Number         `json:"feeRate"`
	FeeAmount         Number         `json:"feeAmount"`
	ExchangeFee       Number         `json:"exchangeFee"`
	ExchangeFeeAmount Number         `json:"exchangeFeeAmount"`
	Sender            string         `json:"sender"`
	Receiver          string         `json:"receiver"`
	Steps             interface{}    `json:"steps"`
}

func (u *userMultiOrder) evmFromEntity(ord entity.Order) *order {
	var mo *userMultiOrder

	ss := strings.Split(ord.ExchangeNid(), "-")
	switch ss[0] {
	case "allbridge":
		o := ord.(*at.Order)
		ls := o.Steps[len(o.Steps)-1]
		mo = &userMultiOrder{
			TotalStpes:        ord.StepsCount(),
			Input:             o.In,
			Output:            o.Out,
			FeeRate:           Number(o.FeeRate),
			FeeAmount:         Number(o.FeeAmount),
			ExchangeFee:       Number(o.ExchangeFee),
			ExchangeFeeAmount: Number(o.EstimateAmountOut),

			Sender:            o.Sender,
			Receiver:          o.Receiver,
			AmountIn:          Number(o.AmountIn),
			AmountOut:         Number(ls.AmountOut),
			EstimateAmountOut: Number(o.EstimateAmountOut),
			Steps:             o.StepsStatus(),
		}

	default:
		o := ord.(*et.Order)
		mo = &userMultiOrder{
			TotalStpes:        ord.StepsCount(),
			Input:             o.In,
			Output:            o.Out,
			Sender:            o.Sender.Hex(),
			Receiver:          o.Receiver.Hex(),
			AmountIn:          Number(o.AmountIn),
			AmountOut:         Number(o.AmountOut),
			EstimateAmountOut: Number(o.EstimateAmountOut),
			FeeRate:           Number(o.FeeRate),
			FeeAmount:         Number(o.FeeAmount),
			ExchangeFee:       Number(o.ExchangeFee),
			ExchangeFeeAmount: Number(o.ExchangeFeeAmount),
			Steps:             o.StepsStatus(),
		}
	}

	return &order{
		Id:        ord.ID().String(),
		Type:      multiSteps,
		UserId:    ord.UserId(),
		CreatedAt: ord.CreatedAt(),
		Order:     mo,
	}
}
