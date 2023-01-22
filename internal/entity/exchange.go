package entity

type ExType uint

func (e ExType) String() string {
	switch e {
	case CEX:
		return "CEX"
	case EvmDEX:
		return "EvmDEX"
	default:
		return ""
	}
}

const (
	CEX ExType = iota
	EvmDEX
)

type Exchange interface {
	Id() string
	Name() string
	Type() ExType
	Price(t1, t2 *Token) (*Pair, error)
	Support(t1, t2 *Token) bool
	AddPairs(data interface{}) (*AddPairsResult, error)
	RemovePair(t1, t2 *Token) error
	Configs() interface{}
	Command(Command) (CommandResult, error)
}
