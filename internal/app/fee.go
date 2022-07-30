package app

func (o *OrderUseCase) GetFee() string {
	return o.fs.GetFee()

}

func (o *OrderUseCase) ChangeFee(f float64) {
	o.fs.ChangeFee(f)
}
