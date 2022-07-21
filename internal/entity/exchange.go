package entity

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

type Exchange interface {
	ID() string
	Exchange(from, to Coin, volume string) (string, error)
	TrackOrder(o *ExchangeOrder, done chan<- struct{}, err chan<- error)

	Withdrawal(coin Coin, address string, vol string) (string, error)
	TrackWithdrawal(w *Withdrawal, done chan<- struct{}, err chan<- error, proccessedCh <-chan bool) error
}
