package entity

import "time"

type OrderCache interface {
	Add(order *UserOrder) error
	Get(userId, id int64) (*UserOrder, error)
	GetAll(userId int64) ([]*UserOrder, error)
	Update(order *UserOrder) error
	UpdateExchangeOrder(eo *ExchangeOrder) error
	UpdateWithdrawal(w *Withdrawal) error
	Delete(userId, id int64) error
}

type WithdrawalCache interface {
	AddPendingWithdrawal(w *Withdrawal) error
	GetPendingWithdrawals(end time.Time) ([]*Withdrawal, error)
	DelPendingWithdrawal(w *Withdrawal) error
}
