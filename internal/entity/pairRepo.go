package entity

type PairRepo interface {
	Add(ex Exchange, ps ...*Pair)
	Get(ex string, t1, t2 *Token) (*Pair, error)
	Exists(ex string, t1, t2 *Token) bool
	Remove(ex string, t1, t2 *Token) error
	GetPaginated(p *Paginated) error
}
