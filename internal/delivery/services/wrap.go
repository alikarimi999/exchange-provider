package services

import (
	"order_service/internal/delivery/services/deposite"
	"order_service/internal/delivery/services/fee"
	"order_service/internal/entity"
)

type Config struct {
	DepositeServiceURL string
	FeeServiceURL      string
}

type Services struct {
	Fee      entity.FeeService
	Deposite entity.DepositeService
}

func WrapServices(cfg *Config) *Services {
	return &Services{
		Fee:      fee.NewFeeService(),
		Deposite: deposite.NewDepositeService(cfg.DepositeServiceURL),
	}
}
