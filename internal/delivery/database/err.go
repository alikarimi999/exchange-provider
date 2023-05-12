package database

import (
	"exchange-provider/pkg/errors"
	"fmt"
)

func invalidErr(param string, s interface{}) error {
	err := fmt.Errorf("%s: '%v' is invalid", param, s)
	return errors.Wrap(errors.ErrBadRequest, err,
		errors.NewMesssage(err.Error()))
}
