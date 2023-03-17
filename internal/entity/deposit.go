package entity

const (
	DepositTxIdSet   string = "txId setted"
	DepositConfirmed string = "confirmed"
	DepositFailed    string = "failed"
)

type Address struct {
	Addr string `json:"address"`
	Tag  string `json:"tag,omitempty"`
}

type Deposit struct {
	Id     string
	Status string
	*Token

	TxId   string
	Volume float64

	Address Address

	FailedDesc string
	ExpireAt   int64
}
