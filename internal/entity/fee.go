package entity

type FeeService interface {
	ApplyFee(userId string, total string) (remainder, fee string, err error)
	GetUserFee(userId string) string
	ChangeUserFee(userId string, fee float64) error
	ChangeDefaultFee(f float64) error
	GetDefaultFee() string
	GetAllUsersFees() map[string]string
}
