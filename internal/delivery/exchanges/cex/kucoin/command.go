package kucoin

import "exchange-provider/internal/entity"

func (k *exchange) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}
