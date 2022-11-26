package cache

import (
	"exchange-provider/internal/delivery/storage/cache/dto"
	"exchange-provider/internal/entity"
	"fmt"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

var prefix_pending_key = "pending_withdrawals"

func (c *OrderCache) AddPendingWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("WithdrawalCache.AddPendingWithdrawal")

	dto := dto.OWToDTO(w)

	if err := c.r.ZAdd(c.ctx, prefix_pending_key, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: dto,
	}).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) GetPendingWithdrawals(end time.Time) ([]*entity.Withdrawal, error) {
	const op = errors.Op("WithdrawalCache.GetPendingWithdrawals")

	ws := []*dto.Withdrawal{}
	err := c.r.ZRangeByScore(c.ctx, prefix_pending_key, &redis.ZRangeBy{
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

func (c *OrderCache) DelPendingWithdrawal(w entity.Withdrawal) error {
	const op = errors.Op("WithdrawalCache.DelPendingWithdrawal")

	if err := c.r.ZRem(c.ctx, prefix_pending_key, dto.OWToDTO(&w)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}
