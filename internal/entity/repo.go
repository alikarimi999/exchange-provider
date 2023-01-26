package entity

import "time"

type OrderRepo interface {
	Add(Order) error
	Update(Order) error
	Get(id string) (Order, error)
	GetAll(userId uint64) ([]Order, error)
	GetPaginated(ps *Paginated) error
	TxIdExists(txId string) (bool, error)
	AddPendingWithdrawal(orderId string) error
	GetPendingWithdrawals(end time.Time) ([]string, error)
	DelPendingWithdrawal(orderId string) error
}

type Paginated struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Orders               []Order
	Pairs                []*Pair
}
