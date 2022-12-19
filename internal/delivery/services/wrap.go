package services

import (
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/delivery/services/exrepo"
	"exchange-provider/internal/delivery/services/fee"
	"exchange-provider/internal/delivery/services/pairconf"
	walletstore "exchange-provider/internal/delivery/services/wallet-store"
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
	entity.FeeService
	entity.PairConfigs
	app.ExchangeRepo
	app.WalletStore
}

func WrapServices(cfg *Config) (*Services, error) {
	ws := walletstore.NewWalletStore()
	ss := &Services{
		ExchangeRepo: exrepo.NewExchangeRepo(cfg.DB, ws, cfg.V, cfg.RC, cfg.L, cfg.PrvKey),
		WalletStore:  ws,
	}

	f, err := fee.NewFeeService(cfg.DB, cfg.V)
	if err != nil {
		return nil, err
	}

	s, err := pairconf.NewPairConfigs(cfg.DB, cfg.V, cfg.L)
	if err != nil {
		return nil, err
	}

	ss.PairConfigs = s

	ss.FeeService = f

	return ss, nil
}
