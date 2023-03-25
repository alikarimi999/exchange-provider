package entity

type PairsRepo interface {
	Add(ex Exchange, ps ...*Pair) error
	Get(exId uint, t1, t2 string) (*Pair, bool)
	GetAll(exId uint) []*Pair
	Update(exId uint, p *Pair)
	Exists(exId uint, t1, t2 string) bool
	Remove(exId uint, t1, t2 string) error
	GetPaginated(*Paginated) error
}
