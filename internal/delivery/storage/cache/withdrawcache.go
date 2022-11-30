package cache

import (
	"fmt"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

var prefix_pending_key = "pending_withdrawals"

func (c *OrderCache) AddPendingWithdrawal(orderId int64) error {
	const op = errors.Op("WithdrawalCache.AddPendingWithdrawal")

	if err := c.r.ZAdd(c.ctx, prefix_pending_key, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: orderId,
	}).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) GetPendingWithdrawals(end time.Time) ([]int64, error) {
	const op = errors.Op("WithdrawalCache.GetPendingWithdrawals")

	ids := []int64{}
	err := c.r.ZRangeByScore(c.ctx, prefix_pending_key, &redis.ZRangeBy{
		Min: "-inf",
		Max: fmt.Sprintf("%d", end.Unix()),
	}).ScanSlice(&ids)

	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	return ids, nil
}

func (c *OrderCache) DelPendingWithdrawal(orderId int64) error {
	const op = errors.Op("WithdrawalCache.DelPendingWithdrawal")

	if err := c.r.ZRem(c.ctx, prefix_pending_key, orderId).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}
