package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
)

type UserOrder struct {
	Id            int64          `json:"order_id"`
	UserId        int64          `json:"user_id"`
	CreatedAt     int64          `json:"created_at"`
	Status        string         `json:"status"`
	Deposit       *Deposit       `json:"deposit"`
	Exchange      string         `json:"exchange"`
	Withdrawal    *Withdrawal    `json:"withdrawal"`
	RequestCoin   string         `json:"request_coin"`
	RequestChain  string         `json:"request_chain"`
	ProvideCoin   string         `json:"provide_coin"`
	ProvideChain  string         `json:"provide_chain"`
	ExchangeOrder *ExchangeOrder `json:"exchange_order"`
	Broken        bool           `json:"broken"`
	BrokeReason   string         `json:"broke_reason"`
}

func UoFromEntity(o *entity.UserOrder) *UserOrder {
	return &UserOrder{
		Id:            o.Id,
		UserId:        o.UserId,
		CreatedAt:     o.CreatedAt,
		Status:        string(o.Status),
		Deposit:       DFromEntity(o.Deposite),
		Exchange:      o.Exchange,
		Withdrawal:    WFromEntity(o.Withdrawal),
		RequestCoin:   o.RequestCoin.Id,
		RequestChain:  o.RequestCoin.Chain.Id,
		ProvideCoin:   o.ProvideCoin.Id,
		ProvideChain:  o.ProvideCoin.Chain.Id,
		ExchangeOrder: EoFromEntity(o.ExchangeOrder),
		Broken:        o.Broken,
		BrokeReason:   o.BrokeReason,
	}
}

type CreateOrderRequest struct {
	UserId  int64  `json:"user_id"`
	Address string `json:"address"`
	RCoin   string `json:"r_coin"`
	RChain  string `json:"r_chain"`

	PCoin  string `json:"p_coin"`
	PChain string `json:"p_chain"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.UserId == 0 {
		return errors.Wrap(errors.NewMesssage("user_id is required"))
	}
	if r.Address == "" {
		return errors.Wrap(errors.NewMesssage("address is required"))
	}
	if r.RCoin == "" {
		return errors.Wrap(errors.NewMesssage("r_coin is required"))
	}
	if r.RChain == "" {
		return errors.Wrap(errors.NewMesssage("r_chain is required"))
	}

	if r.PCoin == "" {
		return errors.Wrap(errors.NewMesssage("p_coin is required"))
	}
	if r.PChain == "" {
		return errors.Wrap(errors.NewMesssage("p_chain is required"))
	}

	return nil
}

type CreateOrderResponse struct {
	Id              int64  `json:"id"`
	DepositeId      int64  `json:"deposite_id"`
	DepositeAddress string `json:"deposite_address"`
}

type GetOrderResponse struct {
}
