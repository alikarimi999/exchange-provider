package swapspace

import "exchange-provider/internal/entity"

func (ex *exchange) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}
