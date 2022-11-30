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

func (c *OrderCache) Add(o *entity.Order) error {
	return c.save(o)
}

func (c *OrderCache) UpdateDeposit(d *entity.Deposit) error {

	o, err := c.get(d.OrderId)
	if err != nil {
		return err
	}

	o.Deposit = d
	return c.save(o)
}

func (c *OrderCache) save(o *entity.Order) error {
	const op = errors.Op("OrderCache.save")

	key := fmt.Sprintf("order:%d", o.Id)
	if err := c.r.Set(c.ctx, key, &dto.Order{Order: o}, time.Duration(48*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) Get(id int64) (*entity.Order, error) {
	return c.get(id)
}

func (c *OrderCache) get(id int64) (*entity.Order, error) {
	const op = errors.Op("OrderCache.get")

	key := fmt.Sprintf("order:%d", id)
	o := &entity.Order{}
	b, err := c.r.Get(c.ctx, key).Bytes()

	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	if err = json.Unmarshal(b, o); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	if o.MetaData == nil {
		o.MetaData = make(map[string]interface{})
	}
	for _, s := range o.Swaps {
		if s.MetaData == nil {
			s.MetaData = make(map[string]interface{})
		}
	}
	return o, nil
}

func (c *OrderCache) Update(o *entity.Order) error {
	const op = errors.Op("OrderCache.Update")
	if err := c.Add(o); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

func (c *OrderCache) Delete(id int64) error {
	const op = errors.Op("OrderCache.Delete")

	key := fmt.Sprintf("order:%d", id)
	if err := c.r.Del(c.ctx, key).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) UpdateWithdrawal(w *entity.Withdrawal) error {
	const op = errors.Op("OrderCache.UpdateWithdrawal")

	o, err := c.Get(w.OrderId)
	if err != nil {
		return errors.Wrap(err, op)
	}
	o.Withdrawal = w
	if err := c.Add(o); err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}
