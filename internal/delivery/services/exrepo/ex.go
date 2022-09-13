package exrepo

import (
	"crypto/rsa"
	"order_service/internal/app"
	"order_service/internal/entity"
	"order_service/pkg/logger"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ExchangeRepo struct {
	db *gorm.DB
	v  *viper.Viper
	l  logger.Logger
	rc *redis.Client

	prv *rsa.PrivateKey
}

func NewExchangeRepo(db *gorm.DB, v *viper.Viper, rc *redis.Client, l logger.Logger, prvKey *rsa.PrivateKey) app.ExchangeRepo {
	return &ExchangeRepo{
		db:  db,
		v:   v,
		l:   l,
		rc:  rc,
		prv: prvKey,
	}
}

func (a *ExchangeRepo) Add(ex *app.Exchange) error {

	e, err := a.encryptConfigs(ex)
	if err != nil {
		return err
	}
	return a.db.Save(e).Error

}

func (a *ExchangeRepo) UpdateStatus(ex entity.Exchange, s string) error {

	return a.db.Model(&Exchange{}).Where("id = ?", ex.AccountId()).Update("status", s).Error

}

func (a *ExchangeRepo) GetAll() ([]*app.Exchange, error) {

	agent := "ExchangeRepo.GetAll"

	var exs []*app.Exchange
	var exchanges []Exchange
	if err := a.db.Find(&exchanges).Error; err != nil {
		return nil, err
	}

	for _, ex := range exchanges {

		exc, err := a.decrypt(&ex)
		if err != nil {
			a.l.Error(agent, err.Error())
			continue
		}

		exs = append(exs, &app.Exchange{
			Exchange:       exc,
			CurrentStatus:  ex.Status,
			LastChangeTime: time.Now(),
		})
	}

	return exs, nil
}

func (a *ExchangeRepo) Remove(ex entity.Exchange) error {

	return a.db.Delete(&Exchange{}, "id = ?", ex.AccountId()).Error

}
