package entity

const (
	DepositTxIdSet   string = "tx_id_setted"
	DepositConfirmed string = "confirmed"
	DepositFailed    string = "failed"
)

type Address struct {
	Addr string
	Tag  string
}

type Deposit struct {
	Id      int64
	OrderId int64

	Status string
	*Token

	TxId   string
	Volume string

	*Address

	FailedDesc string
}
