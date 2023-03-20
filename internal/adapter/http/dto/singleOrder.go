package dto

import (
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/utils/numbers"
	"math/big"
	"strconv"
	"strings"
	"time"
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
		Id:        o.ObjectId.String(),
		Type:      singleStep,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     a,
	}
}

type userSingleOrder struct {
	Status     string `json:"status"`
	FailReason string `json:"failReason,omitempty"`

	Input     Token   `json:"input"`
	Output    Token   `json:"output"`
	InAmount  float64 `json:"inAmount"`
	OutAmount float64 `json:"outAmount"`
	Duration  string  `json:"duration"`

	FilledPrice string `json:"filledPrice"`
	Fee         string `json:"fee"`
	FeeCurrency string `json:"feeCurrency"`

	TransferFee         string `json:"transferFee"`
	TransferFeeCurrency string `json:"transferFee_currency"`

	Deposit    entity.Address `json:"deposit"`
	Refund     entity.Address `json:"refund"`
	Withdrawal entity.Address `json:"withdrawal"`

	WithdrawalTxId string `json:"withdrawTxId"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
	ExpireAt       int64  `json:"expireAt"`
}

func (s *userSingleOrder) fromEntity(o *entity.CexOrder) *order {

	out, _ := strconv.ParseFloat(o.Withdrawal.Volume, 64)
	s = &userSingleOrder{
		Input:     tokenFromEntity(o.Routes[0].In, false),
		Output:    tokenFromEntity(o.Routes[0].Out, false),
		InAmount:  o.Deposit.Volume,
		OutAmount: out,
		Duration:  o.Swaps[0].Duration,

		Fee:                 o.Fee,
		FeeCurrency:         o.FeeCurrency,
		TransferFee:         o.Withdrawal.Fee,
		TransferFeeCurrency: o.Withdrawal.FeeCurrency,

		Deposit:    o.Deposit.Address,
		Refund:     o.Refund,
		Withdrawal: o.Withdrawal.Address,

		WithdrawalTxId: strings.Split(o.Withdrawal.TxId, "-")[0],
		CreatedAt:      o.CreatedAt,
		UpdatedAt:      o.UpdatedAt,
		ExpireAt:       o.Deposit.ExpireAt,
	}

	switch o.Status {
	case entity.OExpired:
		s.Status = o.Status
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
		if o.Deposit.ExpireAt > 0 && time.Now().Unix() >= o.Deposit.ExpireAt {
			s.Status = "expired"
		}
	}
	return &order{
		Id:        o.ObjectId.String(),
		Type:      singleStep,
		UserId:    o.UserId,
		CreatedAt: o.CreatedAt,
		Order:     s,
	}
}
