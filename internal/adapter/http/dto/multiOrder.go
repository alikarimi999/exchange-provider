package dto

import (
	"exchange-provider/internal/entity"
)

type evmStep struct {
	*route
	IsApprove bool
	Approved  bool
}
type route struct {
	Input    string
	Output   string
	Exchange string
}

type adminMultiOrder struct {
	Steps    map[uint]*evmStep
	Sender   string
	Receiver string
	AmountIn float64 `bson:"amountIn"`
	FeeRate  float64 `bson:"feeRate"`
}

func (a *adminMultiOrder) fromEntity(o *entity.DexOrder) *order {
	a = &adminMultiOrder{
		Steps:    make(map[uint]*evmStep),
		Sender:   o.Sender.Hex(),
		Receiver: o.Receiver.Hex(),
		AmountIn: a.AmountIn,
		FeeRate:  a.FeeRate,
	}
	for i, s := range o.Steps {
		a.Steps[i] = &evmStep{
			route:     &route{Input: s.In.String(), Output: s.Out.String(), Exchange: s.Exchange},
			IsApprove: s.NeedApprove,
			Approved:  s.Approved,
		}
	}
	return &order{
		Id:        o.ObjectId.String(),
		Type:      multiSteps,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     a,
	}
}

type userMultiOrder struct {
	TotalStpes int `bson:"totalStpes"`
	Sender     string
	Receiver   string
	AmountIn   float64 `bson:"amountIn"`
	FeeRate    float64 `bson:"feeRate"`
}

func (u *userMultiOrder) fromEntity(o *entity.DexOrder) *order {
	u = &userMultiOrder{
		TotalStpes: len(o.Steps),
		Sender:     o.Sender.Hex(),
		Receiver:   o.Sender.Hex(),
		AmountIn:   o.AmountIn,
		FeeRate:    o.FeeRate,
	}
	return &order{
		Id:        o.ObjectId.String(),
		Type:      multiSteps,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     u,
	}
}
