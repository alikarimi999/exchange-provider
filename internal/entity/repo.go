package entity

type OrderRepo interface {
	Add(order *Order) error
	Update(order *Order) error
	UpdateDeposit(d *Deposit) error
	Get(id int64) (*Order, error)
	GetAll(userId int64) ([]*Order, error)
	// get paginated orders
	GetPaginated(ps *PaginatedOrders) error
	CheckTxId(txId string) (bool, error)
}

type PaginatedOrders struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Orders               []*Order
}
