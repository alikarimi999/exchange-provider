package services

import (
	"order_service/internal/delivery/services/deposite"
	"order_service/internal/delivery/services/fee"
	"order_service/internal/delivery/services/pairconf"
	"order_service/internal/entity"

	"gorm.io/gorm"
)

type Config struct {
	DepositeServiceURL string
	FeeServiceURL      string
	DB                 *gorm.DB
}

type Services struct {
	Fee      entity.FeeService
	PairConf entity.PairConfigs
	Deposite entity.DepositeService
}

func WrapServices(cfg *Config) (*Services, error) {
	ss := &Services{
		Deposite: deposite.NewDepositeService(cfg.DepositeServiceURL),
	}

	f, err := fee.NewFeeService(cfg.DB)
	if err != nil {
		return nil, err
	}

	s, err := pairconf.NewPairConfigs(cfg.DB)
	if err != nil {
		return nil, err
	}

	ss.PairConf = s

	ss.Fee = f
	return ss, nil
}
