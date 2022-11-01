package entity

import (
	"sync"
)

type ExType string

const (
	DEX ExType = "DEX"
	CEX ExType = "CEX"
)

type ExOrderStatus string

const (
	ExOrderSucceed ExOrderStatus = "succeed"
	ExOrderFailed  ExOrderStatus = "failed"
	ExOrderPending ExOrderStatus = "pending"
)

type ExchangeOrder struct {
	Id          uint64
	ExId        string
	UserId      int64
	OrderId     int64
	Status      ExOrderStatus // succed, failed
	Exchange    string
	Symbol      string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string
	FailedDesc  string
}

type Exchange interface {
	Name() string
	AccountId() string
	NID() string

	Exchange(o *UserOrder, size, funds string) (string, error)
	TrackExchangeOrder(o *UserOrder, done chan<- struct{}, proccessed <-chan bool)
	TrackDeposit(d *Deposit, done chan<- struct{}, proccessed <-chan bool)

	Withdrawal(o *UserOrder, coin *Coin, address *Address, vol string) (string, error)
	TrackWithdrawal(w *Withdrawal, done chan<- struct{}, proccessedCh <-chan bool)

	ExchangeManager
}

type ExchangeManager interface {
	Type() ExType
	Stop()
	StartAgain() (*StartAgainResult, error)

	Command(Command) (CommandResult, error)

	Run(wg *sync.WaitGroup)
	Configs() interface{}

	// add pairs to the exchange, if pair exist ignore it
	AddPairs(data interface{}) (*AddPairsResult, error)
	// get all pairs from the exchange
	GetAllPairs() []*Pair
	GetPair(bc, qc *Coin) (*Pair, error)

	RemovePair(bc, qc *Coin) error

	// check if the exchange support a pair with combination of two coins
	Support(bc, qc *Coin) bool

	GetAddress(c *Coin) (*Address, error)
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

type StartAgainResult struct {
	Removed []*PairsErr
}
