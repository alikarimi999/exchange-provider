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

func OrderFromEntityForUser(o entity.Order) *order {
	switch o.Type() {
	case entity.CEXOrder:
		s := &userSingleOrder{}
		return s.fromEntity(o)

	case entity.EVMOrder:
		m := &userMultiOrder{}
		return m.evmFromEntity(o)
	default:
		return nil
	}
}

type CreateOrderRequest struct {
	UserId   string         `json:"userId"`
	Input    entity.TokenId `json:"input"`
	Output   entity.TokenId `json:"output"`
	Sender   entity.Address `json:"sender"`
	Refund   entity.Address `json:"refund"`
	Receiver entity.Address `json:"receiver"`
	AmountIn float64        `json:"amountIn"`
	LP       uint           `json:"lp"`
	Msg      string         `json:"message"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.UserId == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("userId is required"))
	}

	if r.Sender.Addr == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("sender is required"))
	}

	if r.Receiver.Addr == "" {
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
