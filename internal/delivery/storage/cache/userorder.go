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

func (c *OrderCache) Add(order *entity.UserOrder) error {

	o := dto.ToDTO(order)
	if err := c.save(o); err != nil {
		return err
	}
	return c.saveOrderSeq(o.UserId, o.Id, o.Seq)
}

func (c *OrderCache) UpdateDeposit(d *entity.Deposit) error {

	o, err := c.get(d.UserId, d.OrderId)
	if err != nil {
		return err
	}

	o.Deposite = dto.DToDto(d)
	return c.save(o)
}

func (c *OrderCache) save(o *dto.UserOrder) error {
	const op = errors.Op("OrderCache.save")

	key := fmt.Sprintf("user:%d:order:%d", o.UserId, o.Id)
	if err := c.r.Set(c.ctx, key, o, time.Duration(48*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}
func (c *OrderCache) saveOrderSeq(uId, oId, seq int64) error {
	const op = errors.Op("OrderCache.saveOrderSeq")

	key := fmt.Sprintf("user:%d:seq:%d", uId, seq)
	if err := c.r.Set(c.ctx, key, oId, time.Duration(48*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *OrderCache) getOrderIdBySeq(uId, seq int64) (int64, error) {
	const op = errors.Op("OrderCache.getOrderIdBySeq")

	key := fmt.Sprintf("user:%d:seq:%d", uId, seq)
	oId, err := c.r.Get(c.ctx, key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.Wrap(err, op, errors.ErrNotFound)
		}
		return 0, errors.Wrap(err, op, errors.ErrInternal)
	}
	return oId, nil
}

func (c *OrderCache) Get(userId, id int64) (*entity.UserOrder, error) {

	o, err := c.get(userId, id)
	if err != nil {
		return nil, err
	}
	return o.ToEntity(), nil
}

func (c *OrderCache) get(userId, id int64) (*dto.UserOrder, error) {
	const op = errors.Op("OrderCache.get")

	key := fmt.Sprintf("user:%d:order:%d", userId, id)
	o := &dto.UserOrder{}
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
	return o, nil
}

func (c *OrderCache) GetBySeq(uId, seq int64) (*entity.UserOrder, error) {
	const op = errors.Op("OrderCache.GetBySeq")

	oId, err := c.getOrderIdBySeq(uId, seq)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return c.Get(uId, oId)

}

func (c *OrderCache) GetAll(userId int64) ([]*entity.UserOrder, error) {
	const op = errors.Op("OrderCache.GetAll")

	p := fmt.Sprintf("user:%d:order:*", userId)
	var keys []string
	if err := c.r.Keys(c.ctx, p).ScanSlice(&keys); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	if len(keys) == 0 {
		return nil, errors.Wrap(op, errors.ErrNotFound)
	}

	vals, err := c.r.MGet(c.ctx, keys...).Result()
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
	if err := c.r.Del(c.ctx, key).Err(); err != nil {
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
