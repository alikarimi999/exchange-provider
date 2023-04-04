package entity

type ExType uint

const (
	CEX ExType = iota
	EvmDEX
	CossDex
)

type Exchange interface {
	Id() uint
	Name() string
	NID() string
	Type() ExType
	EstimateAmountOut(t1, t2 *Token, amount float64) (amountOut, min float64, err error)
	AddPairs(data interface{}) (*AddPairsResult, error)
	RemovePair(t1, t2 *Token) error
	Configs() interface{}
	Command(Command) (CommandResult, error)
	Remove()
}
