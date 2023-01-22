package storage

import (
	"exchange-provider/internal/delivery/storage/cache"
	"exchange-provider/internal/delivery/storage/database"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Repo entity.OrderRepo
	Oc   entity.OrderCache
}

func NewStorage(db *mongo.Database, rc *redis.Client, l logger.Logger) *Storage {
	return &Storage{
		Repo: database.NewUserRepo(db, l),
		Oc:   cache.NewOrderCache(rc),
	}

}
