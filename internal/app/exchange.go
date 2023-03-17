package app

import (
	"exchange-provider/internal/entity"
)

func (o *OrderUseCase) AddExchange(ex entity.Exchange) error {
	return o.exs.AddExchange(ex)

}

func (o *OrderUseCase) ExchangeExists(id uint) bool {
	return o.exs.exists(id)
}

func (o *OrderUseCase) GetExchange(id uint) (entity.Exchange, error) {
	return o.exs.get(id)
}

func (o *OrderUseCase) AllExchanges(names ...string) []entity.Exchange {
	return o.exs.getByNames(names...)
}
