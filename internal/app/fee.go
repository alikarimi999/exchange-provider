package app

func (o *OrderUseCase) GetDefaultFee() string {
	return o.fs.GetDefaultFee()

}

func (o *OrderUseCase) ChangeDefaultFee(f float64) error {
	return o.fs.ChangeDefaultFee(f)
}

func (o *OrderUseCase) GetUserFee(userId int64) string {
	return o.fs.GetUserFee(userId)
}

func (o *OrderUseCase) GetAllUsersFee() map[int64]string {
	return o.fs.GetAllUsersFees()
}

func (o *OrderUseCase) ChangeUserFee(userId int64, fee float64) error {
	return o.fs.ChangeUserFee(userId, fee)
}
