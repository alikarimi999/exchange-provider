package multichain

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strings"
)

func parseCoin(coin string) (*entity.Coin, error) {
	parts := strings.Split(coin, "-")
	if len(parts) != 2 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage("coin must be in format: <coin_id>-<chain_id>"))
	}

	return &entity.Coin{
		CoinId:  strings.ToUpper(parts[0]),
		ChainId: strings.ToUpper(parts[1]),
	}, nil
}
