package entity

type PairsRepo interface {
	Add(ex Exchange, ps ...*Pair) error
	Get(exId uint, t1, t2 string) (*Pair, error)
	GetAll(exId uint) []*Pair
	UpdateAll(cmd string) error
	Update(exId uint, p *Pair) error
	Exists(exId uint, t1, t2 string) bool
	Remove(exId uint, t1, t2 string, hard bool) error
	RemoveAll(exId uint, hard bool) error
	RemoveAllExchanges() error
	GetPaginated(pa *Paginated, admin bool) error
}
