package entity

type FeeService interface {
	ApplyFee(userId uint64, total string) (remainder, fee string, err error)
	GetUserFee(userId uint64) string
	ChangeUserFee(userId uint64, fee float64) error
	ChangeDefaultFee(f float64) error
	GetDefaultFee() string
	GetAllUsersFees() map[uint64]string
}
