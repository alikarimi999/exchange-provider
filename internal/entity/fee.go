package entity

type FeeService interface {
	ApplyFee(userId int64, total string) (remainder, fee string, err error)
	GetFee() string
	ChangeFee(f float64)
}
