package entity

import "time"

type OrderRepo interface {
	Add(Order) error
	Update(Order) error
	Get(*ObjectId) (Order, error)
	GetAll(userId string) ([]Order, error)
	GetPaginated(ps *Paginated) error
	TxIdExists(txId string) (bool, error)
	AddPendingWithdrawal(orderId *ObjectId) error
	GetPendingWithdrawals(end time.Time) ([]*ObjectId, error)
	DelPendingWithdrawal(orderId *ObjectId) error
}

type Paginated struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Orders               []Order
	Pairs                []*Pair
}
