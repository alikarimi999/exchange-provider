package app

import (
	"exchange-provider/internal/entity"
	"fmt"
)

func (o *OrderUseCase) AddExchange(ex entity.Exchange) error {
	if exists := o.exs.exists(ex.Id()); !exists {
		return o.exs.AddExchange(ex)
	}
	return fmt.Errorf("exchange %d already exists", ex.Id())
}

func (o *OrderUseCase) GetExchange(id uint) (entity.Exchange, error) {
	return o.exs.get(id)
}

func (o *OrderUseCase) AllExchanges(names ...string) []entity.Exchange {
	return o.exs.getByNames(names...)
}
