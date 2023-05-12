package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
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
	TotalStpes uint `bson:"totalStpes"`
	Sender     string
	Receiver   string
	AmountIn   Number `bson:"amountIn"`
}

func (u *userMultiOrder) evmFromEntity(ord entity.Order) *order {
	o := ord.(*types.Order)
	u = &userMultiOrder{
		TotalStpes: ord.Steps(),
		Sender:     o.Sender.Hex(),
		Receiver:   o.Sender.Hex(),
		AmountIn:   Number(o.AmountIn),
	}
	return &order{
		Id:        o.ObjectId.String(),
		Type:      multiSteps,
		UserId:    o.UserID,
		CreatedAt: o.CreatedAT,
		Order:     u,
	}
}
