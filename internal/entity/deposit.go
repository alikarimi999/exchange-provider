package entity

type Deposit struct {
	Id       int64
	UserId   int64
	OrderId  int64
	Exchange string
	Volume   string

	Fullfilled bool
	Address    string
	Tag        string
}

type DepositeService interface {
	New(userId, orderId int64, coin *Coin, exchange string) (*Deposit, error)
	SetTxId(userId, orderId, depositeId int64, txId string) error
}

type Depositcoin struct {
	CoinId   string
	ChainId  string
	SetChain bool
}
