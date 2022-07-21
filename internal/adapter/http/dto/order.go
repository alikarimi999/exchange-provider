package dto

import "order_service/internal/entity"

type UserOrder struct {
	Id            int64
	UserId        int64
	CreatedAt     int64
	Status        string
	Deposite      *Deposite
	Exchange      string
	Withdrawal    *Withdrawal
	RequestCoin   string
	RequestChain  string
	ProvideCoin   string
	ProvideChain  string
	ExchangeOrder *ExchangeOrder
	Broken        bool
	BrokeReason   string
}

func UoFromEntity(o *entity.UserOrder) *UserOrder {
	return &UserOrder{
		Id:            o.Id,
		UserId:        o.UserId,
		CreatedAt:     o.CreatedAt,
		Status:        string(o.Status),
		Deposite:      DFromEntity(o.Deposite),
		Exchange:      o.Exchange,
		Withdrawal:    WFromEntity(o.Withdrawal),
		RequestCoin:   o.RequestCoin.Symbol,
		RequestChain:  string(o.RequestCoin.Chain),
		ProvideCoin:   o.ProvideCoin.Symbol,
		ProvideChain:  string(o.ProvideCoin.Chain),
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

func (r *CreateOrderRequest) Validate() bool {
	if r.UserId == 0 {
		return false
	}
	if r.Address == "" {
		return false
	}
	if r.RCoin == "" {
		return false
	}
	if r.PCoin == "" {
		return false
	}

	return true
}

type CreateOrderResponse struct {
	Id              int64  `json:"id"`
	DepositeId      int64  `json:"deposite_id"`
	DepositeAddress string `json:"deposite_address"`
}

type GetOrderResponse struct {
}
