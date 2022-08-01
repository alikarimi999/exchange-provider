package entity

type OrderRepo interface {
	Add(order *UserOrder) error
	Update(order *UserOrder) error
	Get(userId, id int64) (*UserOrder, error)
	GetAll(userId int64) ([]*UserOrder, error)
	// get paginated orders
	GetPaginated(ps *PaginatedUserOrders) error
}

type PaginatedUserOrders struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Orders               []*UserOrder
}
