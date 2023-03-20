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
	UserId   string          `json:"userId"`
	Input    Token           `json:"input"`
	Output   Token           `json:"output"`
	Refund   *entity.Address `json:"refund"`
	Receiver *entity.Address `json:"receiver"`
	AmountIn float64         `json:"amountIn"`
	LP       uint            `json:"lp"`
	Msg      string          `json:"message"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.UserId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("userId is required"))
	}

	if r.Refund == nil || r.Refund.Addr == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("refund is required"))
	}
	if r.Receiver == nil || r.Receiver.Addr == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("receiver is required"))
	}
	if r.Input.Symbol == "" || r.Input.Standard == "" || r.Input.Network == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("input is invalid"))
	}
	if r.Output.Symbol == "" || r.Output.Standard == "" || r.Output.Network == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("output is invalid"))
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
