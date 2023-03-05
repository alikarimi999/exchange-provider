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
	FailedDesc  string
	MetaData
}

type Cex interface {
	Exchange
	Swap(o *CexOrder, index int) (string, error)
	TrackSwap(o *CexOrder, index int, done chan<- struct{}, proccessed <-chan bool)
	TrackDeposit(o *CexOrder, done chan<- struct{}, proccessed <-chan bool)

	Withdrawal(o *CexOrder) (string, error)
	TrackWithdrawal(w *CexOrder, done chan<- struct{}, proccessedCh <-chan bool)
	Run()

	GetAddress(c *Token) (*Address, error)
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
