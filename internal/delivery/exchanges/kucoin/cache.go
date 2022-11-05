package kucoin

import (
	"context"
	"encoding/json"
	"exchange-provider/internal/delivery/exchanges/kucoin/dto"
	"exchange-provider/pkg/logger"
	"fmt"
	"time"

	"exchange-provider/pkg/errors"

	"github.com/go-redis/redis/v9"
)

type cache struct {
	k   *kucoinExchange
	r   *redis.Client
	ctx context.Context
	l   logger.Logger
}

func newCache(k *kucoinExchange, r *redis.Client, l logger.Logger) *cache {
	return &cache{
		k:   k,
		r:   r,
		ctx: context.Background(),
		l:   l,
	}
}

func (c *cache) recordWithdrawal(w *dto.Withdrawal) error {
	op := errors.Op(fmt.Sprintf("%s.cache.recordWithdrawal", c.k.NID()))
	key := fmt.Sprintf("kucoin:withdrawals:%s", w.Id)
	if err := c.r.Set(c.ctx, key, w, time.Duration(24*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *cache) getWithdrawal(id string) (*dto.Withdrawal, error) {
	op := errors.Op(fmt.Sprintf("%s.cache.getWithdrawal", c.k.NID()))

	key := fmt.Sprintf("kucoin:withdrawals:%s", id)
	v, err := c.r.Get(c.ctx, key).Result()
	if err != nil {
		return nil, err
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

func (c *cache) delWithdrawal(id string) error {
	op := errors.Op(fmt.Sprintf("%s.cache.delWithdrawal", c.k.NID()))

	key := fmt.Sprintf("kucoin:withdrawals:%s", id)
	if err := c.r.Del(c.ctx, key).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *cache) proccessedWithdrawal(id string) error {
	op := errors.Op(fmt.Sprintf("%s.cache.proccessedWithdrawal", c.k.NID()))

	key := fmt.Sprintf("kucoin:proccessed:withdrawals:%s", id)
	if err := c.r.Set(c.ctx, key, "", time.Duration(2*time.Hour)).Err(); err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

// check if withdrawal is processed
func (c *cache) isAddable(id string) (bool, error) {
	op := errors.Op(fmt.Sprintf("%s.cache.isAddable", c.k.NID()))

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

func (c *cache) SaveD(de *depositeRecord) error {
	op := fmt.Sprintf("%s.cache.recordDeposite", c.k.NID())

	key := fmt.Sprintf("kucoin:deposites:%s", de.TxId)
	err := c.r.Set(c.ctx, key, de, time.Duration(12*time.Hour)).Err()
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *cache) GetD(txid string) (*depositeRecord, error) {
	op := errors.Op(fmt.Sprintf("%s.cache.getDepositRecord", c.k.NID()))

	key := fmt.Sprintf("kucoin:deposites:%s", txid)
	d := &depositeRecord{}
	b, err := c.r.Get(c.ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}

	if err := json.Unmarshal(b, d); err != nil {
		return nil, errors.Wrap(err, op, errors.ErrInternal)
	}
	d.TxId = txid
	return d, nil
}

func (c *cache) RemoveD(txid string) error {
	op := errors.Op(fmt.Sprintf("%s.cache.purgeDepositRecord", c.k.NID()))

	key := fmt.Sprintf("kucoin:deposites:%s", txid)
	err := c.r.Del(c.ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, op, errors.ErrInternal)
	}
	return nil
}

func (c *cache) ExistD(txid string) (bool, error) {

	key := fmt.Sprintf("kucoin:deposites:%s", txid)
	i, err := c.r.Exists(c.ctx, key).Result()
	if err != nil {
		return false, err
	}

	if i == 1 {
		return true, nil
	}
	return false, nil
}
