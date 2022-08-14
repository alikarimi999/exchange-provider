package services

import (
	"crypto/rsa"
	"order_service/internal/app"
	"order_service/internal/delivery/services/deposite"
	"order_service/internal/delivery/services/exrepo"
	"order_service/internal/delivery/services/fee"
	"order_service/internal/delivery/services/pairconf"
	"order_service/internal/entity"
	"order_service/pkg/logger"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	FeeServiceURL string
	DB            *gorm.DB
	V             *viper.Viper
	L             logger.Logger
	RC            *redis.Client
	PrvKey        *rsa.PrivateKey
}

type Services struct {
	Fee      entity.FeeService
	PairConf entity.PairConfigs
	Deposite entity.DepositeService
	ExRepo   app.ExchangeRepo
}

func WrapServices(cfg *Config) (*Services, error) {
	ss := &Services{
		Deposite: deposite.NewDepositeService(cfg.V, cfg.L),
		ExRepo:   exrepo.NewExchangeRepo(cfg.DB, cfg.V, cfg.RC, cfg.L, cfg.PrvKey),
	}

	f, err := fee.NewFeeService(cfg.DB, cfg.V)
	if err != nil {
		return nil, err
	}

	s, err := pairconf.NewPairConfigs(cfg.DB, cfg.V, cfg.L)
	if err != nil {
		return nil, err
	}

	ss.PairConf = s

	ss.Fee = f

	return ss, nil
}
