package cache

import (
	"context"
	"encoding/json"
	"exchange-provider/internal/delivery/storage/cache/dto"
	"exchange-provider/internal/entity"
	"fmt"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type OrderCache struct {
	r   *redis.Client
	ctx context.Context
}

func NewOrderCache(c *redis.Client) entity.OrderCache {
	return &OrderCache{
		r:   c,
		ctx: context.Background(),
	}
}

func (c *OrderCache) Add(o entity.Order) error {
	return c.save(o)
}

func (c *OrderCache) save(o entity.Order) error {
	const op = errors.Op("OrderCache.save")

	do, err := dto.ToDto(o)
	if err != nil {
		return err
	}
	if err := c.r.Set(c.ctx, c.key(o.ID()), do,
		time.Duration(48*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) Get(id string) (entity.Order, error) {
	return c.get(id)
}

func (c *OrderCache) get(id string) (entity.Order, error) {
	const op = errors.Op("OrderCache.get")

	o := &dto.Order{}
	b, err := c.r.Get(c.ctx, c.key(id)).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	if err = json.Unmarshal(b, o); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	return o.ToEntity()
}

func (c *OrderCache) Update(o *entity.CexOrder) error {
	return c.save(o)
}

func (c *OrderCache) Delete(id string) error {
	const op = errors.Op("OrderCache.Delete")

	if err := c.r.Del(c.ctx, c.key(id)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) key(id string) string {
	return fmt.Sprintf("order:%s", id)
}
