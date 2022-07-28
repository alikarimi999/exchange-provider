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
		RequestCoin:   o.BC.CoinId,
		RequestChain:  o.BC.ChainId,
		ProvideCoin:   o.QC.CoinId,
		ProvideChain:  o.QC.ChainId,
		ExchangeOrder: EoFromEntity(o.ExchangeOrder),
		Broken:        o.Broken,
		BrokeReason:   o.BrokeReason,
	}
}

type CreateOrderRequest struct {
	UserId int64 `json:"user_id"`

	BC     string `json:"base_coin"`
	BChain string `json:"base_chain"`

	QC     string `json:"quote_coin"`
	QChain string `json:"quote_chain"`

	Side string `json:"side"`

	Address string `json:"address"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.UserId == 0 {
		return errors.Wrap(errors.NewMesssage("user_id is required"))
	}
	if r.Address == "" {
		return errors.Wrap(errors.NewMesssage("address is required"))
	}
	if r.BC == "" {
		return errors.Wrap(errors.NewMesssage("base_coin is required"))
	}
	if r.BChain == "" {
		return errors.Wrap(errors.NewMesssage("base_chain is required"))
	}

	if r.QC == "" {
		return errors.Wrap(errors.NewMesssage("quote_coin is required"))
	}
	if r.QChain == "" {
		return errors.Wrap(errors.NewMesssage("quote_chain is required"))
	}

	if r.Side != "buy" && r.Side != "sell" {
		return errors.Wrap(errors.NewMesssage("only buy or sell is allowed for side"))
	}

	return nil
}

type CreateOrderResponse struct {
	OrderId         int64  `json:"order_id"`
	DepositeId      int64  `json:"deposit_id"`
	DepositeAddress string `json:"deposit_address"`
}

type GetOrderResponse struct {
}
