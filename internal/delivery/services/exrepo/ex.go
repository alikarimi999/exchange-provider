package exrepo

import (
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"sync"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ExchangeRepo struct {
	db *gorm.DB
	v  *viper.Viper
	l  logger.Logger
	rc *redis.Client
	app.WalletStore
	prv *rsa.PrivateKey
}

func NewExchangeRepo(db *gorm.DB, ws app.WalletStore, v *viper.Viper, rc *redis.Client, l logger.Logger, prvKey *rsa.PrivateKey) app.ExchangeRepo {
	return &ExchangeRepo{
		db:          db,
		v:           v,
		l:           l,
		rc:          rc,
		WalletStore: ws,
		prv:         prvKey,
	}
}

func (a *ExchangeRepo) Add(ex entity.Exchange) error {

	e, err := a.encryptConfigs(ex)
	if err != nil {
		return err
	}
	return a.db.Save(e).Error

}

func (a *ExchangeRepo) GetAll() ([]entity.Exchange, error) {
	agent := "ExchangeRepo.GetAll"

	var exs []entity.Exchange
	var exchanges []Exchange
	if err := a.db.Find(&exchanges).Error; err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	for _, exc := range exchanges {
		wg.Add(1)
		go func(ex Exchange) {
			defer wg.Done()
			exc, err := a.decrypt(&ex)
			if err != nil {
				a.l.Error(agent, err.Error())
				return
			}

			exs = append(exs, exc)
		}(exc)
	}
	wg.Wait()
	return exs, nil
}

func (a *ExchangeRepo) Remove(ex entity.Exchange) error {
	return a.db.Delete(&Exchange{}, "id = ?", ex.Id()).Error
}
