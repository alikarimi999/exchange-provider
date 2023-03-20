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
	Id() uint
	Name() string
	Type() ExType
	EstimateAmountOut(t1, t2 *Token, amount float64) (amountOut, min float64, err error)
	AddPairs(data interface{}) (*AddPairsResult, error)
	RemovePair(t1, t2 *Token) error
	Configs() interface{}
	Command(Command) (CommandResult, error)
	Remove()
}
