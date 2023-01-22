package entity

const (
	SwapSucceed string = "succeed"
	SwapFailed  string = "failed"
)

type Swap struct {
	Id   uint64
	TxId string `bson:"txId"`

	OrderId string `bson:"orderId"`
	Status  string // succed, failed

	InAmount    string `bson:"inAmount"`
	OutAmount   string `bson:"outAmount"`
	Fee         string
	FeeCurrency string `bson:"feeCurrency"`
	FailedDesc  string `bson:"failedDesc"`
	MetaData    `bson:"metaData"`
}

type Cex interface {
	Exchange
	Swap(o *CexOrder, index int) (string, error)
	TrackSwap(o *CexOrder, index int, done chan<- struct{}, proccessed <-chan bool)
	TrackDeposit(o *CexOrder, done chan<- struct{}, proccessed <-chan bool)

	Withdrawal(o *CexOrder) (string, error)
	TrackWithdrawal(w *CexOrder, done chan<- struct{}, proccessedCh <-chan bool)

	Stop()
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
