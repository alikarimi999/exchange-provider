package app

import "exchange-provider/internal/entity"

type ExchangeRepo interface {
	Add(ex entity.Exchange) error
	GetAll() ([]entity.Exchange, error)
	EnableDisable(exId uint, enable bool) error
	EnableDisableAll(enable bool) error
	Remove(ex entity.Exchange) error
	RemoveAll() error
}
