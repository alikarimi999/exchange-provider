package entity

type WithdrawalStatus string

const (
	WithdrawalSucceed WithdrawalStatus = "succeed"
	WithdrawalFailed  WithdrawalStatus = "failed"
	WithdrawalPending WithdrawalStatus = "pending"
)

type Withdrawal struct {
	Id      uint64
	OrderId int64

	Status WithdrawalStatus
	TxId   string
	*Address

	*Token
	Unwrapped bool
	Volume    string

	Fee         string
	FeeCurrency string

	FailedDesc string
}
