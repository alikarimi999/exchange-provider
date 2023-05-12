package entity

const (
	SwapSucceed string = "succeed"
	SwapFailed  string = "failed"
)

type Swap struct {
	Id     uint64
	TxId   string
	Status string // succed, failed

	In          *Token
	Out         *Token
	InAmount    string
	OutAmount   string
	Fee         string
	FeeCurrency string
	Duration    string
	FailedDesc  string
	MetaData
}

type Cex interface {
	Exchange
	TxIdSetted(Order, string) error
}

type AddPairsResult struct {
	Added   []string
	Existed []string
	Failed  []*PairsErr
}

type UpdatePairResult struct {
	Updated []string
	Failed  []*PairsErr
}

type PairsErr struct {
	Pair string
	Err  error
}
