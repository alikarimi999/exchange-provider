package entity

const (
	WithdrawalSucceed string = "succeed"
	WithdrawalFailed  string = "failed"
	WithdrawalPending string = "pending"
)

type Withdrawal struct {
	Id      string
	OrderId string `bson:"orderId"`

	Status string
	TxId   string `bson:"txId"`
	*Address

	*Token
	Unwrapped bool
	Volume    string

	Fee         string
	FeeCurrency string `bson:"feeCurrency"`

	FailedDesc string `bson:"failedDesc"`
}
