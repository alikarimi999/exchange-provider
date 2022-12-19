package entity

import (
	"sync"
)

type ExType string

const (
	DEX ExType = "DEX"
	CEX ExType = "CEX"
)

const (
	SwapSucceed string = "succeed"
	SwapFailed  string = "failed"
)

type Swap struct {
	Id   uint64
	TxId string

	OrderId int64
	Status  string // succed, failed

	InAmount    string
	OutAmount   string
	Fee         string
	FeeCurrency string
	FailedDesc  string
	MetaData
}

type Exchange interface {
	Id() string
	Name() string
	Exchange(o *Order, index int) (string, error)
	TrackExchangeOrder(o *Order, index int, done chan<- struct{}, proccessed <-chan bool)
	TrackDeposit(o *Order, done chan<- struct{}, proccessed <-chan bool)

	Withdrawal(o *Order) (string, error)
	TrackWithdrawal(w *Order, done chan<- struct{}, proccessedCh <-chan bool)

	ExchangeManager
}

type ExchangeManager interface {
	Type() ExType
	Stop()

	Command(Command) (CommandResult, error)

	Run(wg *sync.WaitGroup)
	Configs() interface{}

	// add pairs to the exchange, if pair exist ignore it
	AddPairs(data interface{}) (*AddPairsResult, error)
	// get all pairs from the exchange
	GetAllPairs() []*Pair
	GetPair(c1, c2 *Token) (*Pair, error)

	RemovePair(c1, c2 *Token) error

	// check if the exchange support a pair with combination of two coins
	Support(in, out *Token) bool

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
