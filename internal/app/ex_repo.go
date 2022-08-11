package app

import "order_service/internal/entity"

type ExchangeRepo interface {
	Add(ex *Exchange) error
	UpdateStatus(ex entity.Exchange, s string) error
	GetAll() ([]*Exchange, error)
	Remove(ex entity.Exchange) error
}
