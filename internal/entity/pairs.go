package entity

type ExchangePair interface {
	Snapshot() ExchangePair
}

type Pair struct {
	T1 *Token
	T2 *Token

	LP       uint
	Exchange string
}

func (p *Pair) String() string {
	return p.T1.String() + "/" + p.T2.String()
}

func (p *Pair) Equal(p1 *Pair) bool {
	return (p.T1.Equal(p1.T1) && p.T2.Equal(p1.T2))
}

func (p *Pair) Snapshot() *Pair {
	return &Pair{
		T1:       p.T1.Snapshot(),
		T2:       p.T2.Snapshot(),
		LP:       p.LP,
		Exchange: p.Exchange,
	}
}
