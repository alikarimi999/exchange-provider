package entity

import "time"

type OrderCache interface {
	Add(order *UserOrder) error
	UpdateDeposit(d *Deposit) error
	Get(userId, id int64) (*UserOrder, error)
	GetAll(userId int64) ([]*UserOrder, error)
	GetBySeq(uId, seq int64) (*UserOrder, error)
	Delete(userId, id int64) error

	WithdrawalCache
}

type WithdrawalCache interface {
	AddPendingWithdrawal(w *Withdrawal) error
	GetPendingWithdrawals(end time.Time) ([]*Withdrawal, error)
	DelPendingWithdrawal(w Withdrawal) error
}
