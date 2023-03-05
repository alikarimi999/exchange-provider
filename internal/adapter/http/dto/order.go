package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
)

type order struct {
	Id        string      `json:"id"`
	Type      string      `json:"type"`
	UserId    string      `json:"userId"`
	CreatedAt int64       `json:"createdAt"`
	Order     interface{} `json:"order"`
}

func userOrderFromEntity(o entity.Order) *order {
	switch o.Type() {
	case entity.CEXOrder:
		s := &userSingleOrder{}
		return s.fromEntity(o.(*entity.CexOrder))
	case entity.EVMOrder:
		m := &userMultiOrder{}
		return m.fromEntity(o.(*entity.EvmOrder))
	default:
		return nil
	}
}

func adminOrderFromEntity(o entity.Order) *order {
	switch o.Type() {
	case entity.CEXOrder:
		s := &adminSingleOrder{}
		return s.fromEntity(o.(*entity.CexOrder))
	case entity.EVMOrder:
		m := &adminMultiOrder{}
		return m.fromEntity(o.(*entity.EvmOrder))
	default:
		return nil
	}
}

type CreateOrderRequest struct {
	UserId   string  `json:"userId"`
	In       string  `json:"input"`
	Out      string  `json:"output"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Tag      string  `json:"tag"`
	AmountIn float64 `json:"amountIn"`
	Msg      string  `json:"message"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.UserId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("userId is required"))
	}
	if r.Sender == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("sender is required"))
	}
	if r.Receiver == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("receiver is required"))
	}
	if r.In == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("input is required"))
	}
	if r.Out == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("output is required"))
	}
	if r.AmountIn == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("amountIn is required"))
	}
	return nil
}

const (
	singleStep string = "SingleStep"
	multiSteps string = "MultiStep"
)

type createOrderResponse struct {
	OrderId    string `json:"orderId"`
	Type       string `json:"type"`
	TotalSteps int    `json:"totalSteps"`
}

func CreateOrderResponse(o entity.Order) *createOrderResponse {
	r := &createOrderResponse{OrderId: o.ID().String()}
	if o.Type() == entity.CEXOrder {
		r.Type = singleStep
		r.TotalSteps = 1
	} else {
		r.Type = multiSteps
		r.TotalSteps = len(o.(*entity.EvmOrder).Steps)
	}
	return r
}
