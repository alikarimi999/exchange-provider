package dto

import (
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"strings"
)

func ParseToken(t string) (*entity.Token, error) {
	parts := strings.Split(t, "-")
	if len(parts) != 2 {
		return nil, errors.Wrap(errors.ErrBadRequest,
			errors.NewMesssage("token must be in format: <tokenId>-<chainId>"))
	}

	return &entity.Token{
		TokenId: parts[0],
		ChainId: parts[1],
	}, nil
}
