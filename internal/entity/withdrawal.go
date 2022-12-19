package entity

const (
	WithdrawalSucceed string = "succeed"
	WithdrawalFailed  string = "failed"
	WithdrawalPending string = "pending"
)

type Withdrawal struct {
	Id      uint64
	OrderId int64

	Status string
	TxId   string
	*Address

	*Token
	Unwrapped bool
	Volume    string

	Fee         string
	FeeCurrency string

	FailedDesc string
}
