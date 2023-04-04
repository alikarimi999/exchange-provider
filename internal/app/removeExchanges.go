package app

func (o *OrderUseCase) RemoveExchange(id uint, force bool) error {
	return o.exs.remove(id)
}
