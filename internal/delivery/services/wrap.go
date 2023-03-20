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

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DB     *mongo.Database
	Repo   entity.OrderRepo
	Pairs  entity.PairsRepo
	V      *viper.Viper
	L      logger.Logger
	PrvKey *rsa.PrivateKey
}

type Services struct {
	entity.FeeService
	entity.PairConfigs
	app.ExchangeRepo
	app.WalletStore
}

func WrapServices(cfg *Config) (*Services, error) {

	f, err := fee.NewFeeService(cfg.DB, cfg.V)
	if err != nil {
		return nil, err
	}

	s, err := pairconf.NewPairConfigs(cfg.DB, cfg.V, cfg.L)
	if err != nil {
		return nil, err
	}

	ws := walletstore.NewWalletStore()

	ss := &Services{
		PairConfigs: s,
		FeeService:  f,
		ExchangeRepo: exrepo.NewExchangeRepo(cfg.DB, ws, cfg.Pairs, cfg.Repo,
			f, s, cfg.V, cfg.L, cfg.PrvKey),
		WalletStore: ws,
	}
	return ss, nil
}
