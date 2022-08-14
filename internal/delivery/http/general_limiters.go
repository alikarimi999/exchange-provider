package http

type limiters struct {
	conf     *LimiterConfig
	limiters []*rateLimiter
}

func newGeneralLimiters(conf *LimiterConfig) *limiters {
	return &limiters{
		conf: &LimiterConfig{
			Max:    conf.Max,
			Period: conf.Period,
		},
	}
}

func (l *limiters) addLimiter() *rateLimiter {
	limiter := newLimiter(l.conf)
	l.limiters = append(l.limiters, limiter)
	return limiter
}

func (l *limiters) changeConfigs(conf *LimiterConfig) error {
	l.conf = conf
	for _, limiter := range l.limiters {
		limiter.ChangeConfigs(conf)
	}
	return nil
}
