package dto

import "exchange-provider/pkg/errors"

var errInvalidID error = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("invalid id"))
