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

func UOFromEntity(oe *entity.UserOrder) *UserOrder {
	o := &UserOrder{
		Id:        oe.Seq,
		BaseCoin:  oe.BC.String(),
		QuoteCoin: oe.QC.String(),
		Side:      oe.Side,

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

	if oe.Side == "buy" {
		o.FinalSize = oe.Withdrawal.Executed
		o.FinalFunds = oe.Deposit.Volume
		o.FeeCurrency = oe.BC.String()
	} else {
		o.FinalSize = oe.Deposit.Volume
		o.FinalFunds = oe.Withdrawal.Executed
		o.FeeCurrency = oe.QC.String()
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

type AdminUserOrder struct {
	Id        int64  `json:"order_id"`
	UserId    int64  `json:"user_id"`
	Seq       int64  `json:"seq"`
	Status    string `json:"status"`
	BaseCoin  string `json:"base_coin"`
	QuoteCoin string `json:"quote_coin"`
	Side      string `json:"side"`

	SpreadRate string `json:"spread_rate"`
	SpreadVol  string `json:"spread_vol"`

	Exchange string `json:"exchange"`

	Deposit         *Deposit       `json:"deposit"`
	ExchangeOrder   *ExchangeOrder `json:"exchange_order"`
	Withdrawal      *Withdrawal    `json:"withdrawal"`
	CreatedAt       int64          `json:"created_at"`
	FaileCode       int64          `json:"failed_code"`
	FailedDesc      string         `json:"failed_desc"`
	entity.MetaData `json:"meta_data"`
}

func AdminUOFromEntity(o *entity.UserOrder) *AdminUserOrder {
	return &AdminUserOrder{
		Id:         o.Id,
		UserId:     o.UserId,
		Seq:        o.Seq,
		CreatedAt:  o.CreatedAt,
		Status:     string(o.Status),
		Deposit:    DFromEntity(o.Deposit),
		Exchange:   o.Exchange,
		Withdrawal: WFromEntity(o.Withdrawal),
		BaseCoin:   o.BC.String(),
		QuoteCoin:  o.QC.String(),
		Side:       o.Side,

		SpreadRate: o.SpreadRate,
		SpreadVol:  o.SpreadVol,

		ExchangeOrder: EoFromEntity(o.ExchangeOrder),
		FaileCode:     o.FailedCode,
		FailedDesc:    o.FailedDesc,
		MetaData:      o.MetaData,
	}
}

type CreateOrderRequest struct {
	BC string `json:"base_coin"`
	QC string `json:"quote_coin"`

	Side    string `json:"side"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.Address == "" {
		return errors.Wrap(errors.NewMesssage("address is required"))
	}
	if r.BC == "" {
		return errors.Wrap(errors.NewMesssage("base_coin is required"))
	}

	if r.QC == "" {
		return errors.Wrap(errors.NewMesssage("quote_coin is required"))
	}

	if r.Side != "buy" && r.Side != "sell" {
		return errors.Wrap(errors.NewMesssage("only buy or sell is allowed for side"))
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
