package kucoin

import "exchange-provider/internal/entity"

func (k *kucoinExchange) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}
