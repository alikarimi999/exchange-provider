package store

import (
	"context"
	"crypto/rsa"
	"exchange-provider/internal/app"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type exchangeRepo struct {
	db     *mongo.Collection
	repo   entity.OrderRepo
	pairs  entity.PairsRepo
	exs    entity.ExchangeStore
	fee    entity.FeeTable
	spread entity.SpreadTable
	l      logger.Logger
	app.WalletStore
	prv *rsa.PrivateKey
}

func newExchangeRepo(db *mongo.Database, ws app.WalletStore, pairs entity.PairsRepo,
	repo entity.OrderRepo, fee entity.FeeTable, spread entity.SpreadTable,
	l logger.Logger, prvKey *rsa.PrivateKey) *exchangeRepo {
	return &exchangeRepo{
		db:          db.Collection("exchange-repository"),
		repo:        repo,
		pairs:       pairs,
		spread:      spread,
		fee:         fee,
		l:           l,
		WalletStore: ws,
		prv:         prvKey,
	}
}

func (a *exchangeRepo) add(ex entity.Exchange) error {
	e, err := a.encryptConfigs(ex)
	if err != nil {
		return err
	}
	_, err = a.db.InsertOne(context.Background(), e)
	return err
}

func (a *exchangeRepo) getAll(lastUpdate time.Time) ([]entity.Exchange, error) {
	agent := "ExchangeRepo.GetAll"

	var exs []entity.Exchange
	var exchanges []*Exchange
	cur, err := a.db.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	cur.All(context.Background(), &exchanges)
	wg := &sync.WaitGroup{}
	var allbridge *Exchange
	for _, exc := range exchanges {
		if exc.Name == "allbridge" {
			allbridge = exc
			continue
		}
		wg.Add(1)
		go func(ex *Exchange) {
			defer wg.Done()
			exc, err := a.decrypt(ex, lastUpdate)
			if err != nil {
				a.pairs.RemoveAll(ex.Id, false)
				a.l.Error(agent, err.Error())
				return
			}

			exs = append(exs, exc)
		}(exc)
	}
	wg.Wait()

	if allbridge != nil {
		exc, err := a.decrypt(allbridge, lastUpdate)
		if err != nil {
			a.pairs.RemoveAll(allbridge.Id, false)
			a.l.Error(agent, err.Error())
		} else {
			exs = append(exs, exc)
		}
	}

	return exs, nil
}

func (a *exchangeRepo) enableDisable(exId uint, enable bool) error {
	update := bson.M{"$set": bson.M{"enable": enable}}
	_, err := a.db.UpdateByID(context.Background(), exId, update)
	return err
}
func (a *exchangeRepo) enableDisableAll(enable bool) error {
	update := bson.M{"$set": bson.M{"enable": enable}}
	_, err := a.db.UpdateMany(context.Background(), bson.M{}, update)
	return err
}

func (a *exchangeRepo) eemove(ex entity.Exchange) error {
	if err := a.pairs.RemoveAll(ex.Id(), true); err != nil {
		return err
	}
	d, err := a.db.DeleteOne(context.Background(), bson.M{"_id": ex.Id()})
	if d.DeletedCount > 0 {
		a.l.Debug("ExchangeRepo.Remove", fmt.Sprintf("exchange %s deleted", ex.NID()))
	}
	return err
}

func (a *exchangeRepo) removeAll() error {
	if err := a.pairs.RemoveAllExchanges(); err != nil {
		return err
	}

	_, err := a.db.DeleteMany(context.Background(), bson.M{})
	return err
}
