package exrepo

import (
	"crypto/rsa"
	"fmt"
	"order_service/internal/app"
	"order_service/internal/delivery/exchanges/kucoin"
	"order_service/internal/entity"
	"order_service/pkg/errors"
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

	switch ex.Name() {
	case "kucoin":

		cfg, err := a.encryptKucoinConfigs(ex)
		if err != nil {
			return err
		}
		return a.db.Save(cfg).Error

	default:
		return nil
	}
}

func (a *ExchangeRepo) UpdateStatus(ex entity.Exchange, s string) error {
	switch ex.Name() {
	case "kucoin":
		return a.db.Model(&KucoinExchange{}).Where("id = ?", ex.AccountId()).Update("status", s).Error
	default:
		return nil
	}
}

func (a *ExchangeRepo) GetAll() ([]*app.Exchange, error) {
	const op = errors.Op("ExchangeRepo.GetAllConfigs")
	var exs []*app.Exchange
	var kucoinExs []KucoinExchange
	if err := a.db.Find(&kucoinExs).Error; err != nil {
		return nil, err
	}
	for _, ex := range kucoinExs {
		cfg, err := a.decryptKucoinConfigs(&ex)
		if err != nil {
			return nil, err
		}

		exc, err := kucoin.NewKucoinExchange(cfg, a.rc, a.v, a.l, true)

		if err != nil {
			a.l.Error(string(op), fmt.Sprintf("error creating exchange: %s", err.Error()))
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
	switch ex.Name() {
	case "kucoin":
		return a.db.Delete(&KucoinExchange{}, "id = ?", ex.AccountId()).Error
	default:
		return nil
	}
}
