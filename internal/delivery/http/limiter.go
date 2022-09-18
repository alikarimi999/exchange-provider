package http

import (
	"context"
	"order_service/pkg/errors"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

var (
	defaultLimiterMax    = 10000
	defaultLimiterPeriod = 1 * time.Second
)

type LimiterConfig struct {
	Max    uint64        `json:"max"`
	Period time.Duration `json:"period"`
}

type rateLimiter struct {
	conf  *LimiterConfig
	store limiter.Store
}

func newLimiter(conf *LimiterConfig) *rateLimiter {

	var max uint64
	var period time.Duration
	if conf != nil {
		if conf.Max > 0 {
			max = conf.Max
		} else {
			max = uint64(defaultLimiterMax)
		}
		if conf.Period > 0 {
			period = conf.Period
		} else {
			period = defaultLimiterPeriod
		}
	} else {
		max = uint64(defaultLimiterMax)
		period = defaultLimiterPeriod
	}

	s, _ := memorystore.New(&memorystore.Config{
		Tokens:        max,
		Interval:      period,
		SweepInterval: time.Hour,
		SweepMinTTL:   2 * time.Hour,
	})

	return &rateLimiter{
		conf:  &LimiterConfig{Max: max, Period: period},
		store: s,
	}
}

func (l *rateLimiter) ChangeConfigs(conf *LimiterConfig) error {
	var max uint64
	var period time.Duration
	if conf != nil {
		if conf.Max > 0 {
			max = conf.Max
		} else {
			max = uint64(defaultLimiterMax)
		}
		if conf.Period > 0 {
			period = conf.Period
		} else {
			period = defaultLimiterPeriod
		}
	} else {
		return errors.Wrap(errors.NewMesssage("config is nil"))
	}

	s, _ := memorystore.New(&memorystore.Config{
		Tokens:        max,
		Interval:      period,
		SweepInterval: time.Hour,
		SweepMinTTL:   2 * time.Hour,
	})

	l.store.Close(context.Background())
	l.store = s
	l.conf = conf
	return nil
}

func (l *rateLimiter) allow(ctx context.Context, ip, key string) bool {

	if ip != "" {
		_, _, _, ok, _ := l.store.Take(ctx, ip)

		if ok {
			if key != "" {
				_, _, _, ok, _ := l.store.Take(ctx, key)
				return ok
			}
		}
	}

	return false
}
