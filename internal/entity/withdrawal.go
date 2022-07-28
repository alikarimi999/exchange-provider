package entity

import (
	"fmt"
)

type WithdrawalStatus string

const (
	WithdrawalSucceed WithdrawalStatus = "succeed"
	WithdrawalFailed  WithdrawalStatus = "failed"
	WithdrawalPending WithdrawalStatus = "pending"
)

type Withdrawal struct {
	Id      string
	OrderId int64
	UserId  int64
	Address string

	Exchange string

	Total       string
	Fee         string
	ExchangeFee string
	Executed    string

	TxId   string
	Status WithdrawalStatus
}

func (w *Withdrawal) String() string {
	return fmt.Sprintf("%+v", *w)
}
