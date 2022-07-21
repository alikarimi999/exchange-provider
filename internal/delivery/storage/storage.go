package storage

import (
	"order_service/internal/delivery/storage/cache"
	"order_service/internal/delivery/storage/database"
	"order_service/internal/entity"
	"order_service/pkg/logger"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	Repo entity.OrderRepo
	Oc   entity.OrderCache
	Wc   entity.WithdrawalCache
}

func NewStorage(db *gorm.DB, rc *redis.Client, l logger.Logger) *Storage {
	return &Storage{
		Repo: database.NewUserRepo(db),
		Wc:   cache.NewWithdrawalCache(rc),
		Oc:   cache.NewOrderCache(rc),
	}

}
