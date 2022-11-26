package entity

import "time"

type OrderCache interface {
	Add(order *Order) error
	UpdateDeposit(d *Deposit) error
	Get(userId, id int64) (*Order, error)
	GetAll(userId int64) ([]*Order, error)
	Delete(userId, id int64) error

	WithdrawalCache
}

type WithdrawalCache interface {
	AddPendingWithdrawal(w *Withdrawal) error
	GetPendingWithdrawals(end time.Time) ([]*Withdrawal, error)
	DelPendingWithdrawal(w Withdrawal) error
}
