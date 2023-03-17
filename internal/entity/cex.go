package entity

const (
	SwapSucceed string = "succeed"
	SwapFailed  string = "failed"
)

type Swap struct {
	Id     uint64
	TxId   string
	Status string // succed, failed

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
	TxIdSetted(*CexOrder)
	Run()
	SetDepositddress(o *CexOrder) error
}

type AddPairsResult struct {
	Added   []Pair
	Existed []string
	Failed  []*PairsErr
}

type PairsErr struct {
	Pair string
	Err  error
}
