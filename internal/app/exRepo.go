package app

import "exchange-provider/internal/entity"

type ExchangeRepo interface {
	Add(ex entity.Exchange) error
	GetAll() ([]entity.Exchange, error)
	Remove(ex entity.Exchange) error
}
