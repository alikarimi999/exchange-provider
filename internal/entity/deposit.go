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
	Id     string
	Status string
	*Token

	TxId   string
	Volume string

	*Address

	FailedDesc string
}
