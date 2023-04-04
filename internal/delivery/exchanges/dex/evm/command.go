package evm

import "exchange-provider/internal/entity"

func (d *evmDex) Command(entity.Command) (entity.CommandResult, error) {
	return nil, nil
}
