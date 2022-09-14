package kucoin

import "order_service/internal/entity"

func (k *kucoinExchange) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}
