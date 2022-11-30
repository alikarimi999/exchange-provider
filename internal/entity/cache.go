package entity

import "time"

type OrderCache interface {
	Add(order *Order) error
	UpdateDeposit(d *Deposit) error
	Get(id int64) (*Order, error)
	Delete(id int64) error

	WithdrawalCache
}

type WithdrawalCache interface {
	AddPendingWithdrawal(orderId int64) error
	GetPendingWithdrawals(end time.Time) ([]int64, error)
	DelPendingWithdrawal(orderId int64) error
}
