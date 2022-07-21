package kucoin

import (
	"context"
	"encoding/json"
	"fmt"
	"order_service/internal/delivery/exchanges/kucoin/dto"
	"order_service/pkg/logger"
	"time"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type withdrawalCache struct {
	r   *redis.Client
	ctx context.Context
	l   logger.Logger
}

func newWithdrawalCache(r *redis.Client, l logger.Logger) *withdrawalCache {
	return &withdrawalCache{
		r:   r,
		ctx: context.Background(),
		l:   l,
	}
}

func (c *withdrawalCache) recordWithdrawal(w *dto.Withdrawal) error {
	const op = errors.Op("Kucoin.WithdrawalCache.recordWithdrawal")

	key := fmt.Sprintf("kucoin:withdrawals:%s", w.Id)
	if err := c.r.Set(c.ctx, key, w, time.Duration(0)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *withdrawalCache) getWithdrawal(id string) (*dto.Withdrawal, error) {
	const op = errors.Op("Kucoin.WithdrawalCache.getWithdrawal")

	key := fmt.Sprintf("kucoin:withdrawals:%s", id)
	v, err := c.r.Get(c.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, errors.ErrNotFound, op)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	if v == "" {
		return nil, nil
	}
	w := &dto.Withdrawal{}

	if err = json.Unmarshal([]byte(v), w); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	return w, nil
}

func (c *withdrawalCache) delWithdrawal(id string) error {
	const op = errors.Op("Kucoin.WithdrawalCache.delWithdrawal")

	key := fmt.Sprintf("kucoin:withdrawals:%s", id)
	if err := c.r.Del(c.ctx, key).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *withdrawalCache) proccessedWithdrawal(id string) error {
	const op = errors.Op("Kucoin.WithdrawalCache.proccessedWithdrawal")

	key := fmt.Sprintf("kucoin:proccessed:withdrawals:%s", id)
	if err := c.r.Set(c.ctx, key, "", time.Duration(2*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

// check if withdrawal is processed
func (c *withdrawalCache) isAddable(id string) (bool, error) {
	const op = errors.Op("Kucoin.WithdrawalCache.isWithdrawalProcessed")

	key1 := fmt.Sprintf("kucoin:proccessed:withdrawals:%s", id)
	key2 := fmt.Sprintf("kucoin:withdrawals:%s", id)

	i, err := c.r.Exists(c.ctx, key1, key2).Result()
	if err != nil {
		return false, errors.Wrap(err, op, errors.ErrInternal)
	}
	if i == 0 {
		return false, nil
	}

	return true, nil
}
