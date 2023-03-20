package entity

type PairsRepo interface {
	Add(exId uint, ps ...*Pair)
	Get(exId uint, t1, t2 string) (*Pair, bool)
	Update(exId uint, p *Pair)
	Exists(exId uint, t1, t2 string) bool
	Remove(exId uint, t1, t2 string)
	GetPaginated(*Paginated) error
}
