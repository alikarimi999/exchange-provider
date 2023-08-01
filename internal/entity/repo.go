package entity

type OrderRepo interface {
	Add(Order) error
	Update(Order) error
	Get(*ObjectId) (Order, error)
	GetPaginated(ps *Paginated, onlyCount bool) error
	TxIdExists(txId string) (bool, error)
	GetWithFilter(key string, value interface{}) ([]Order, error)
}

type Paginated struct {
	Page, PerPage, Total int64
	Filters              []*Filter
	Desc                 bool
	Orders               []Order
	Pairs                []*Pair
	Exs                  map[string]Exchange
}
