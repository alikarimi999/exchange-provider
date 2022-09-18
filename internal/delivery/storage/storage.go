package storage

import (
	"exchange-provider/internal/delivery/storage/cache"
	"exchange-provider/internal/delivery/storage/database"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/logger"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	Repo entity.OrderRepo
	Oc   entity.OrderCache
}

func NewStorage(db *gorm.DB, rc *redis.Client, l logger.Logger) *Storage {
	return &Storage{
		Repo: database.NewUserRepo(db),
		Oc:   cache.NewOrderCache(rc),
	}

}
