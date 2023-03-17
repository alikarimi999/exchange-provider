package dto

import (
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"
	"strings"
)

type adminSingleOrder struct {
	Status     string        `json:"status"`
	Deposit    *Deposit      `json:"deposit"`
	Swaps      map[int]*Swap `json:"swaps"`
	Withdrawal *Withdrawal   `json:"withdraw"`

	Fee         string `json:"fee"`
	FeeCurrency string `json:"fee_currency"`

	SpreadRate     string `json:"spread_rate"`
	SpreadVol      string `json:"spread_vol"`
	SpreadCurrency string `json:"spread_currency"`

	FaileCode       int64  `json:"failed_code,omitempty"`
	FailedDesc      string `json:"failed_desc,omitempty"`
	entity.MetaData `json:"meta_data,omitempty"`
}

func (a *adminSingleOrder) fromEntity(o *entity.CexOrder) *order {
	a = &adminSingleOrder{
		Status:     o.Status,
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
		a.Swaps[k] = swapFromEntity(v, o.Routes[k])
	}
	return &order{
		Id:        o.Id,
		Type:      singleStep,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     a,
	}
}

type userSingleOrder struct {
	Status     string `json:"status"`
	FailReason string `json:"failReason,omitempty"`

	Input     string `json:"input"`
	Output    string `json:"output"`
	InAmount  string `json:"inAmount"`
	OutAmount string `json:"outAmount"`

	FilledPrice string `json:"filledPrice"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"feeCurrency"`

	TransferFee         string `json:"transferFee"`
	TransferFeeCurrency string `json:"transferFee_currency"`

	DepositAddress string `json:"depositAddress"`
	DepositTag     string `json:"depositTag,omitempty"`

	WithdrawalAddress string `json:"withdrawAddress"`
	WithdrawalTag     string `json:"withdrawTag,omitempty"`

	WithdrawalTxId string `json:"withdrawTxId"`
	CreatedAt      int64  `json:"createdAt"`
}

func (s *userSingleOrder) fromEntity(o *entity.CexOrder) *order {
	s = &userSingleOrder{
		Input:     o.Deposit.Token.String(),
		Output:    o.Withdrawal.Token.String(),
		InAmount:  fmt.Sprintf("%v", o.Deposit.Volume),
		OutAmount: o.Withdrawal.Volume,

		Fee:                 o.Fee,
		FeeCurrency:         o.FeeCurrency,
		TransferFee:         o.Withdrawal.Fee,
		TransferFeeCurrency: o.Withdrawal.FeeCurrency,

		DepositAddress: o.Deposit.Address.Addr,
		DepositTag:     o.Deposit.Address.Tag,

		WithdrawalAddress: o.Withdrawal.Addr,
		WithdrawalTag:     o.Withdrawal.Address.Tag,

		WithdrawalTxId: strings.Split(o.Withdrawal.TxId, "-")[0],
		CreatedAt:      o.CreatedAt,
	}
	switch o.Status {
	case entity.ODepositeConfimred:
		s.Status = o.Status

	case entity.OSucceeded:
		s.Status = o.Status
		in := big.NewFloat(o.Deposit.Volume)
		out, _ := numbers.StringToBigFloat(o.Withdrawal.Volume)
		fee, _ := numbers.StringToBigFloat(o.Fee)
		s.FilledPrice = new(big.Float).Quo(new(big.Float).Add(out, fee), in).Text('f', 10)

	case entity.OFailed:
		s.Status = o.Status

		switch o.FailedCode {
		case entity.FCDepositFailed:
			if o.Deposit.FailedDesc == app.BR_InsufficientDepositVolume {
				s.FailReason = app.BR_InsufficientDepositVolume
			}
		}
	default:
		s.Status = "pending"
	}
	return &order{
		Id:        o.Id,
		Type:      singleStep,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     s,
	}
}
