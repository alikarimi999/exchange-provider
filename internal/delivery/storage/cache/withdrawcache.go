package cache

import (
	"context"
	"fmt"
	"order_service/internal/delivery/storage/cache/dto"
	"order_service/internal/entity"
	"time"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type WithdrawalCache struct {
	r   *redis.Client
	ctx context.Context
}

func NewWithdrawalCache(r *redis.Client) entity.WithdrawalCache {
	return &WithdrawalCache{
		r:   r,
		ctx: context.Background(),
	}
}

func (c *WithdrawalCache) AddPendingWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("WithdrawalCache.AddPendingWithdrawal")

	dto := dto.WToDTO(w)
	key := fmt.Sprintf("pending_withdrawals")

	if err := c.r.ZAdd(c.ctx, key, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: dto,
	}).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *WithdrawalCache) GetPendingWithdrawals(end time.Time) ([]*entity.Withdrawal, error) {
	const op = errors.Op("WithdrawalCache.GetPendingWithdrawals")

	key := fmt.Sprintf("pending_withdrawals")
	ws := []*dto.PendingWithdrawal{}
	err := c.r.ZRangeByScore(c.ctx, key, &redis.ZRangeBy{
		Min: "-inf",
		Max: fmt.Sprintf("%d", end.Unix()),
	}).ScanSlice(&ws)

	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	ews := []*entity.Withdrawal{}
	for _, w := range ws {
		ews = append(ews, w.ToEntity())
	}

	return ews, nil
}

func (c *WithdrawalCache) DelPendingWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("WithdrawalCache.DelPendingWithdrawal")

	key := fmt.Sprintf("pending_withdrawals")
	if err := c.r.ZRem(c.ctx, key, dto.WToDTO(w)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}
