package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"order_service/internal/delivery/storage/cache/dto"
	"order_service/internal/entity"

	"order_service/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderCache struct {
	c   *redis.Client
	ctx context.Context
}

func NewOrderCache(c *redis.Client) entity.OrderCache {
	return &OrderCache{
		c:   c,
		ctx: context.Background(),
	}
}

func (c *OrderCache) Add(order *entity.UserOrder) error {
	const op = errors.Op("OrderCache.Add")

	o := dto.ToDTO(order)
	key := fmt.Sprintf("user:%d:order:%d", o.UserId, o.Id)
	if err := c.c.Set(c.ctx, key, o, 0).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) Get(userId, id int64) (*entity.UserOrder, error) {
	const op = errors.Op("OrderCache.Get")

	key := fmt.Sprintf("user:%d:order:%d", userId, id)
	o := &dto.UserOrder{}
	b, err := c.c.Get(c.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	if err = json.Unmarshal(b, o); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	return o.ToEntity(), nil
}

func (c *OrderCache) GetAll(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("OrderCache.GetAll")

	p := fmt.Sprintf("user:%d:order:*", userId)
	var keys []string
	if err := c.c.Keys(c.ctx, p).ScanSlice(&keys); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	if len(keys) == 0 {
		return nil, errors.Wrap(op, errors.ErrNotFound)
	}

	vals, err := c.c.MGet(c.ctx, keys...).Result()
	if err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	var osDTO []*dto.UserOrder
	for _, v := range vals {
		o := &dto.UserOrder{}
		if err := json.Unmarshal([]byte(v.(string)), o); err != nil {
			return nil, errors.Wrap(err, op, errors.ErrInternal)
		}
		osDTO = append(osDTO, o)
	}

	var os []*entity.UserOrder
	for _, o := range osDTO {
		os = append(os, o.ToEntity())
	}
	return os, nil

}

func (c *OrderCache) Update(o *entity.UserOrder) error {
	const op = errors.Op("OrderCache.Update")
	if err := c.Add(o); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (c *OrderCache) UpdateExchangeOrder(eo *entity.ExchangeOrder) error {
	const op = errors.Op("OrderCache.UpdateExchangeOrder")

	o, err := c.Get(eo.UserId, eo.OrderId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	o.ExchangeOrder = eo
	if err := c.Add(o); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (c *OrderCache) Delete(userId, id int64) error {
	const op = errors.Op("OrderCache.Delete")

	key := fmt.Sprintf("user:%d:order:%d", userId, id)
	if err := c.c.Del(c.ctx, key).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) UpdateWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("OrderCache.UpdateWithdrawal")

	o, err := c.Get(w.UserId, w.OrderId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	o.Withdrawal = w
	if err := c.Add(o); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
