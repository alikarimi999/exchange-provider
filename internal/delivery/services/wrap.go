package services

import (
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/delivery/services/exrepo"
	"exchange-provider/internal/delivery/services/fee"
	"exchange-provider/internal/delivery/services/pairconf"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	V      *viper.Viper
	L      logger.Logger
	RC     *redis.Client
	PrvKey *rsa.PrivateKey
}

type Services struct {
	Fee      entity.FeeService
	PairConf entity.PairConfigs
	ExRepo   app.ExchangeRepo
}

func WrapServices(cfg *Config) (*Services, error) {
	ss := &Services{
		ExRepo: exrepo.NewExchangeRepo(cfg.DB, cfg.V, cfg.RC, cfg.L, cfg.PrvKey),
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
