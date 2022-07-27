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
	Id          string
	UserId      int64
	OrderId     int64
	Symbol      string
	Exchange    string
	Side        string
	Funds       string
	Size        string
	Fee         string
	FeeCurrency string
	Status      ExOrderStatus // succed, failed
}

type CoinConfig struct {
	*Coin
	SetChain  bool
	Precision int
}

type ExchangePair struct {
	BC *CoinConfig // base coin configs
	QC *CoinConfig // queue coin configs
}

type Exchange interface {
	ID() string
	Exchange(from, to *Coin, volume string) (string, error)
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
	AddPairs(pairs []*ExchangePair)
	// get all pairs from the exchange
	GetPairs() []*Pair
	// check if the exchange support a pair with combination of two coins
	Support(c1, c2 *Coin) bool
}
