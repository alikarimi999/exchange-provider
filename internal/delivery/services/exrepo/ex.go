package exrepo

import (
	"context"
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExchangeRepo struct {
	db   *mongo.Collection
	repo entity.OrderRepo
	fee  entity.FeeService
	pc   entity.PairConfigs
	v    *viper.Viper
	l    logger.Logger
	app.WalletStore
	prv *rsa.PrivateKey
}

func NewExchangeRepo(db *mongo.Database, ws app.WalletStore,
	repo entity.OrderRepo, fee entity.FeeService, pc entity.PairConfigs,
	v *viper.Viper, l logger.Logger, prvKey *rsa.PrivateKey) app.ExchangeRepo {
	return &ExchangeRepo{
		db:          db.Collection("exchange-repository"),
		repo:        repo,
		fee:         fee,
		pc:          pc,
		v:           v,
		l:           l,
		WalletStore: ws,
		prv:         prvKey,
	}
}

func (a *ExchangeRepo) Add(ex entity.Exchange) error {

	e, err := a.encryptConfigs(ex)
	if err != nil {
		return err
	}
	_, err = a.db.InsertOne(context.Background(), e)
	return err
}

func (a *ExchangeRepo) GetAll() ([]entity.Exchange, error) {
	agent := "ExchangeRepo.GetAll"

	var exs []entity.Exchange
	var exchanges []Exchange
	cur, err := a.db.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	cur.All(context.Background(), &exchanges)
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
	d, err := a.db.DeleteOne(context.Background(), bson.D{{"id", ex.Id()}})
	if d.DeletedCount > 0 {
		a.l.Debug("ExchangeRepo.Remove", fmt.Sprintf("exchange %d deleted", ex.Id()))
	}
	return err
}
