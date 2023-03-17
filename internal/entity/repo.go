package entity

type OrderRepo interface {
	Add(Order) error
	Update(Order) error
	Get(*ObjectId) (Order, error)
	GetAll(userId string) ([]Order, error)
	GetPaginated(ps *Paginated) error
	TxIdExists(txId string) (bool, error)
	GetWithFilter(key string, value string) (Order, error)
}

type Paginated struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Orders               []Order
	Pairs                []*Pair
}
