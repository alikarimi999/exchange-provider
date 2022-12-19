package dto

import (
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/utils/numbers"
	"math/big"
	"strings"
)

type UserOrder struct {
	Id         int64  `json:"id"`
	Status     string `json:"status"`
	FailReason string `json:"fail_reason,omitempty"`

	Input     string `json:"input"`
	Output    string `json:"output"`
	InAmount  string `json:"in_amount"`
	OutAmount string `json:"out_amount"`

	FilledPrice string `json:"filled_price"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`

	TransferFee         string `json:"transfer_fee"`
	TransferFeeCurrency string `json:"transfer_fee_currency"`

	DepositAddress string `json:"deposit_address"`
	DepositTag     string `json:"deposit_tag,omitempty"`

	WithdrawalAddress string `json:"withdrawal_address"`
	WithdrawalTag     string `json:"withdrawal_tag,omitempty"`

	WithdrawalTxId string `json:"withdrawal_tx_id"`
	CreatedAt      int64  `json:"created_at"`
}

func UOFromEntity(oe *entity.Order) *UserOrder {
	o := &UserOrder{
		Id: oe.Id,

		Input:     oe.Deposit.Token.String(),
		Output:    oe.Withdrawal.Token.String(),
		InAmount:  oe.Deposit.Volume,
		OutAmount: oe.Withdrawal.Volume,

		Fee:                 oe.Fee,
		FeeCurrency:         oe.FeeCurrency,
		TransferFee:         oe.Withdrawal.Fee,
		TransferFeeCurrency: oe.Withdrawal.FeeCurrency,

		DepositAddress: oe.Deposit.Addr,
		DepositTag:     oe.Deposit.Tag,

		WithdrawalAddress: oe.Withdrawal.Addr,
		WithdrawalTag:     oe.Withdrawal.Tag,

		WithdrawalTxId: strings.Split(oe.Withdrawal.TxId, "-")[0],
		CreatedAt:      oe.CreatedAt,
	}

	switch oe.Status {
	case entity.OSDepositeConfimred:
		o.Status = string(oe.Status)

	case entity.OSSucceed:
		o.Status = string(oe.Status)
		in, _ := numbers.StringToBigFloat(oe.Deposit.Volume)
		out, _ := numbers.StringToBigFloat(oe.Withdrawal.Volume)
		fee, _ := numbers.StringToBigFloat(oe.Fee)
		o.FilledPrice = new(big.Float).Quo(new(big.Float).Add(out, fee), in).Text('f', 10)

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

	return o
}

type AdminOrder struct {
	Id     int64 `json:"order_id"`
	UserId int64 `json:"user_id"`

	Status string `json:"status"`

	Deposit    *Deposit      `json:"deposit"`
	Swaps      map[int]*Swap `json:"swaps"`
	Withdrawal *Withdrawal   `json:"withdrawal"`

	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`

	SpreadRate     string `json:"spread_rate"`
	SpreadVol      string `json:"spread_vol"`
	SpreadCurrency string `json:"spread_currency"`

	CreatedAt int64 `json:"created_at"`

	FaileCode       int64  `json:"failed_code,omitempty"`
	FailedDesc      string `json:"failed_desc,omitempty"`
	entity.MetaData `json:"meta_data,omitempty"`
}

func AdminOrderFromEntity(o *entity.Order) *AdminOrder {
	ord := &AdminOrder{
		Id:     o.Id,
		UserId: o.UserId,

		CreatedAt:  o.CreatedAt,
		Status:     string(o.Status),
		Deposit:    DFromEntity(o.Deposit),
		Swaps:      make(map[int]*Swap),
		Withdrawal: WFromEntity(o.Withdrawal),

		Fee:         o.Fee,
		FeeCurrency: o.FeeCurrency,

		SpreadRate:     o.SpreadRate,
		SpreadVol:      o.SpreadVol,
		SpreadCurrency: o.SpreadCurrency,

		FaileCode:  o.FailedCode,
		FailedDesc: o.FailedDesc,
		MetaData:   o.MetaData,
	}
	for k, v := range o.Swaps {
		ord.Swaps[k] = swapFromEntity(v, o.Routes[k])
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
	DC              string  `json:"deposit_token"`
	MinDeposit      float64 `json:"min_deposit"`
	DepositeAddress string  `json:"deposit_address"`
	AddressTag      string  `json:"address_tag"`
}

type GetOrderResponse struct {
}
