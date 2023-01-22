package app

func (o *OrderUseCase) GetDefaultFee() string {
	return o.fs.GetDefaultFee()

}

func (o *OrderUseCase) ChangeDefaultFee(f float64) error {
	return o.fs.ChangeDefaultFee(f)
}

func (o *OrderUseCase) GetUserFee(userId uint64) string {
	return o.fs.GetUserFee(userId)
}

func (o *OrderUseCase) GetAllUsersFee() map[uint64]string {
	return o.fs.GetAllUsersFees()
}

func (o *OrderUseCase) ChangeUserFee(userId uint64, fee float64) error {
	return o.fs.ChangeUserFee(userId, fee)
}
