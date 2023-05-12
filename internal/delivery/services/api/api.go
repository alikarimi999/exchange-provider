package api

import (
	"context"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiService struct {
	c      *mongo.Collection
	maxIps uint
	l      logger.Logger
}

func NewApiService(db *mongo.Database, maxIps uint, l logger.Logger) (entity.ApiService, error) {
	return &ApiService{
		c:      db.Collection("api"),
		maxIps: maxIps,
		l:      l,
	}, nil
}

func (a *ApiService) AddApiToken(api *entity.APIToken) error {
	_, err := a.c.InsertOne(context.Background(), api)
	return err
}

func (a *ApiService) Get(id string) (*entity.APIToken, error) {
	agent := a.agent("Get")
	res := a.c.FindOne(context.Background(), bson.D{{"_id", id}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, errors.Wrap(errors.ErrNotFound)
		}
		a.l.Debug(agent, res.Err().Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}
	api := &entity.APIToken{}
	if err := res.Decode(api); err != nil {
		a.l.Debug(agent, res.Err().Error())
		return nil, errors.Wrap(errors.ErrInternal)
	}
	return api, nil
}

func (a *ApiService) Update(api *entity.APIToken) error {
	_, err := a.c.ReplaceOne(context.Background(), bson.D{{"_id", api.Id}}, api)
	return err
}

func (a *ApiService) Remove(id string) error {
	res, err := a.c.DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return errors.Wrap(errors.ErrInternal, err)
	}
	if res.DeletedCount == 0 {
		return errors.Wrap(errors.ErrNotFound)
	}
	return nil
}

func (a *ApiService) MaxIps() uint { return a.maxIps }
