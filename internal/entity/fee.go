package entity

type FeeService interface {
	ApplyFee(userId int64, total string) (remainder, fee string, err error)
	GetUserFee(userId int64) string
	ChangeUserFee(userId int64, fee float64) error
	ChangeDefaultFee(f float64) error
	GetDefaultFee() string
	GetAllUsersFees() map[int64]string
}
