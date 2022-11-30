package dto

import (
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strconv"
)

type UserOrder struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	FailReason  string `json:"fail_reason,omitempty"`
	BaseCoin    string `json:"base_coin"`
	QuoteCoin   string `json:"quote_coin"`
	Side        string `json:"side"`
	FinalSize   string `json:"final_size"`
	FinalFunds  string `json:"final_funds"`
	FilledPrice string `json:"filled_price"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	TransferFee string `json:"transfer_fee"`

	DepositAddress string `json:"deposit_address"`
	DepositTag     string `json:"deposit_tag"`

	WithdrawalAddress string `json:"withdrawal_address"`
	WithdrawalTag     string `json:"withdrawal_tag"`

	WithdrawalTxId string `json:"withdrawal_tx_id"`
	CreatedAt      int64  `json:"created_at"`
}

func UOFromEntity(oe *entity.Order) *UserOrder {
	o := &UserOrder{
		Id: oe.Id,

		Fee:            oe.Withdrawal.Fee,
		TransferFee:    oe.Withdrawal.ExchangeFee,
		DepositAddress: oe.Deposit.Addr,
		DepositTag:     oe.Deposit.Tag,

		WithdrawalAddress: oe.Withdrawal.Addr,
		WithdrawalTag:     oe.Withdrawal.Tag,

		WithdrawalTxId: oe.Withdrawal.TxId,
		CreatedAt:      oe.CreatedAt,
	}

	switch oe.Status {
	case entity.OSDepositeConfimred:
		o.Status = string(oe.Status)

	case entity.OSSucceed:
		o.Status = string(oe.Status)
	case entity.OSFailed:
		o.Status = string(oe.Status)

		switch oe.FailedCode {
		case entity.FCDepositFailed:
			if oe.Deposit.FailedDesc == app.BR_InsufficientDepositVolume {
				o.FailReason = app.BR_InsufficientDepositVolume
			}
		}
	default:
		o.Status = "pending"
	}

	if oe.Withdrawal.Total != "" && oe.Deposit.Volume != "" {

		wt, _ := strconv.ParseFloat(oe.Withdrawal.Total, 64)
		dt, _ := strconv.ParseFloat(oe.Deposit.Volume, 64)
		if o.Side == "buy" {
			o.FilledPrice = strconv.FormatFloat(dt/wt, 'f', 8, 64)
		} else {
			o.FilledPrice = strconv.FormatFloat(wt/dt, 'f', 8, 64)
		}

	}

	return o
}

type AdminOrder struct {
	Id     int64 `json:"order_id"`
	UserId int64 `json:"user_id"`

	Status    string `json:"status"`
	BaseCoin  string `json:"base_coin"`
	QuoteCoin string `json:"quote_coin"`
	Side      string `json:"side"`

	SpreadRate string `json:"spread_rate"`
	SpreadVol  string `json:"spread_vol"`

	Exchange string `json:"exchange"`

	Deposit         *Deposit `json:"deposit"`
	Routes          map[int]*entity.Route
	Swaps           map[int]*Swap `json:"swaps"`
	Withdrawal      *Withdrawal   `json:"withdrawal"`
	CreatedAt       int64         `json:"created_at"`
	FaileCode       int64         `json:"failed_code"`
	FailedDesc      string        `json:"failed_desc"`
	entity.MetaData `json:"meta_data"`
}

func AdminOrderFromEntity(o *entity.Order) *AdminOrder {
	ord := &AdminOrder{
		Id:     o.Id,
		UserId: o.UserId,

		CreatedAt:  o.CreatedAt,
		Status:     string(o.Status),
		Deposit:    DFromEntity(o.Deposit),
		Routes:     make(map[int]*entity.Route),
		Withdrawal: WFromEntity(o.Withdrawal),

		SpreadRate: o.SpreadRate,
		SpreadVol:  o.SpreadVol,

		FaileCode:  o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   o.MetaData,
	}
	for k, v := range o.Swaps {
		ord.Swaps[k] = SwapFromEntity(v, o.Routes[k])
	}
	return ord
}

type CreateOrderRequest struct {
	In  string `json:"input"`
	Out string `json:"output"`

	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.Address == "" {
		return errors.Wrap(errors.NewMesssage("address is required"))
	}
	if r.In == "" {
		return errors.Wrap(errors.NewMesssage("input is required"))
	}

	if r.Out == "" {
		return errors.Wrap(errors.NewMesssage("output is required"))
	}

	return nil
}

type CreateOrderResponse struct {
	OrderId         int64   `json:"order_id"`
	DC              string  `json:"deposit_coin"`
	MinDeposit      float64 `json:"min_deposit"`
	DepositeAddress string  `json:"deposit_address"`
	AddressTag      string  `json:"address_tag"`
}

type GetOrderResponse struct {
}
