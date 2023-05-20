package types

import "exchange-provider/internal/entity"

type ExchangePair struct{}

func (e *ExchangePair) Snapshot() entity.ExchangePair { return &ExchangePair{} }
