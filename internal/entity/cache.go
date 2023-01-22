package entity

import "time"

type OrderCache interface {
	Add(Order) error
	Get(id string) (Order, error)
	Delete(id string) error

	WithdrawalCache
}

type WithdrawalCache interface {
	AddPendingWithdrawal(orderId string) error
	GetPendingWithdrawals(end time.Time) ([]string, error)
	DelPendingWithdrawal(orderId string) error
}
