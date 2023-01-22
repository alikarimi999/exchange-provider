package entity

const (
	DepositTxIdSet   string = "txId setted"
	DepositConfirmed string = "confirmed"
	DepositFailed    string = "failed"
)

type Address struct {
	Addr string
	Tag  string
}

type Deposit struct {
	Id      string
	OrderId string `bson:"orderId"`

	Status string
	*Token

	TxId   string `bson:"txId"`
	Volume string

	*Address

	FailedDesc string `bson:"failedDesc"`
}
