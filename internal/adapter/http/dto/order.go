package dto

import (
	"order_service/internal/entity"
	"order_service/pkg/errors"
	"strconv"
)

type UserOrder struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	BaseCoin    string `json:"base_coin"`
	QuoteCoin   string `json:"quote_coin"`
	Side        string `json:"side"`
	FinalSize   string `json:"final_size"`
	FinalFunds  string `json:"final_funds"`
	FilledPrice string `json:"filled_price"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`
	TransferFee string `json:"transfer_fee"`

	DepositAddress    string `json:"deposit_address"`
	WithdrawalAddress string `json:"withdrawal_address"`
	WithdrawalTxId    string `json:"withdrawal_tx_id"`
	CreatedAt         int64  `json:"created_at"`
}

func UOFromEntity(de *entity.UserOrder) *UserOrder {
	d := &UserOrder{
		Id:        de.Id,
		BaseCoin:  de.BC.String(),
		QuoteCoin: de.QC.String(),
		Side:      de.Side,

		Status:            string(de.Status),
		Fee:               de.Withdrawal.Fee,
		TransferFee:       de.Withdrawal.ExchangeFee,
		DepositAddress:    de.Deposite.Address,
		WithdrawalAddress: de.Withdrawal.Address,
		WithdrawalTxId:    de.Withdrawal.TxId,
		CreatedAt:         de.CreatedAt,
	}

	if de.Side == "buy" {
		d.FinalSize = de.Withdrawal.Executed
		d.FinalFunds = de.Deposite.Volume
		d.FeeCurrency = de.BC.String()
	} else {
		d.FinalSize = de.Deposite.Volume
		d.FinalFunds = de.Withdrawal.Executed
		d.FeeCurrency = de.QC.String()
	}

	if de.Withdrawal.Total != "" && de.Deposite.Volume != "" {

		wt, _ := strconv.ParseFloat(de.Withdrawal.Total, 64)
		dt, _ := strconv.ParseFloat(de.Deposite.Volume, 64)
		if d.Side == "buy" {
			d.FilledPrice = strconv.FormatFloat(dt/wt, 'f', 8, 64)
		} else {
			d.FilledPrice = strconv.FormatFloat(wt/dt, 'f', 8, 64)
		}

	}

	return d
}

type AdminUserOrder struct {
	Id        int64  `json:"order_id"`
	UserId    int64  `json:"user_id"`
	Status    string `json:"status"`
	BaseCoin  string `json:"base_coin"`
	QuoteCoin string `json:"quote_coin"`
	Side      string `json:"side"`

	SpreadRate string `json:"spread_rate"`
	SpreadVol  string `json:"spread_vol"`

	Exchange string `json:"exchange"`

	Deposit       *Deposit       `json:"deposit"`
	ExchangeOrder *ExchangeOrder `json:"exchange_order"`
	Withdrawal    *Withdrawal    `json:"withdrawal"`
	CreatedAt     int64          `json:"created_at"`
	Broken        bool           `json:"broken"`
	BreakReason   string         `json:"break_reason"`
}

func AdminUOFromEntity(o *entity.UserOrder) *AdminUserOrder {
	return &AdminUserOrder{
		Id:         o.Id,
		UserId:     o.UserId,
		CreatedAt:  o.CreatedAt,
		Status:     string(o.Status),
		Deposit:    DFromEntity(o.Deposite),
		Exchange:   o.Exchange,
		Withdrawal: WFromEntity(o.Withdrawal),
		BaseCoin:   o.BC.String(),
		QuoteCoin:  o.QC.String(),
		Side:       o.Side,

		SpreadRate: o.SpreadRate,
		SpreadVol:  o.SpreadVol,

		ExchangeOrder: EoFromEntity(o.ExchangeOrder),
		Broken:        o.Broken,
		BreakReason:   o.BreakReason,
	}
}

type CreateOrderRequest struct {
	BC     string `json:"base_coin"`
	BChain string `json:"base_chain"`

	QC     string `json:"quote_coin"`
	QChain string `json:"quote_chain"`

	Side string `json:"side"`

	Address string `json:"address"`
}

func (r *CreateOrderRequest) Validate() error {
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
	OrderId         int64   `json:"order_id"`
	DepositeId      int64   `json:"deposit_id"`
	DC              string  `json:"deposit_coin"`
	MinDeposit      float64 `json:"min_deposit"`
	DepositeAddress string  `json:"deposit_address"`
	AddressTag      string  `json:"address_tag"`
}

type GetOrderResponse struct {
}
