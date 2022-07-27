package entity

type Deposit struct {
	Id         int64
	UserId     int64
	OrderId    int64
	Exchange   string
	Volume     string
	Fullfilled bool
	Address    string
}

type DepositeService interface {
	New(userId, orderId int64, coin *Coin, exchange string) (*Deposit, error)
	Supported(exchang string, coins ...*Coin) ([]*Coin, error)
}
