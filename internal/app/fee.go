package app

func (o *OrderUseCase) GetDefaultFee() string {
	return o.fs.GetDefaultFee()

}

func (o *OrderUseCase) ChangeDefaultFee(f float64) error {
	return o.fs.ChangeDefaultFee(f)
}

func (o *OrderUseCase) GetUserFee(userId string) string {
	return o.fs.GetUserFee(userId)
}

func (o *OrderUseCase) GetAllUsersFee() map[string]string {
	return o.fs.GetAllUsersFees()
}

func (o *OrderUseCase) ChangeUserFee(userId string, fee float64) error {
	return o.fs.ChangeUserFee(userId, fee)
}
