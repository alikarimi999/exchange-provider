package entity

import (
	"order_service/pkg/logger"
	"sync"

	"github.com/go-redis/redis/v9"
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
	Exchange    string
	Symbol      string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string
	Status      ExOrderStatus // succed, failed
}

type Exchange interface {
	ID() string
	Exchange(o *UserOrder) (string, error)
	TrackOrder(o *ExchangeOrder, done chan<- struct{}, err chan<- error)

	Withdrawal(coin *Coin, address string, vol string) (string, error)
	TrackWithdrawal(w *Withdrawal, done chan<- struct{}, err chan<- error, proccessedCh <-chan bool) error

	ExchangeManager
}

type ExchangeManager interface {
	ChangeAccount(cfgi interface{}) error
	Setup(cfg interface{}, rc *redis.Client, l logger.Logger) (Exchange, error)

	Run(wg *sync.WaitGroup)
	Configs() interface{}

	// add pairs to the exchange, if pair exist ignore it
	AddPairs(pairs []*Pair) (*AddPairsResult, error)
	// get all pairs from the exchange
	GetAllPairs() []*Pair
	GetPair(bc, qc *Coin) (*Pair, error)

	RemovePair(bc, qc *Coin) error

	// check if the exchange support a pair with combination of two coins
	Support(bc, qc *Coin) bool
}

type AddPairsResult struct {
	Added   []*Pair
	Existed []*Pair
	Failed  []*AddPairsErr
}

type AddPairsErr struct {
	*Pair
	Err error
}
