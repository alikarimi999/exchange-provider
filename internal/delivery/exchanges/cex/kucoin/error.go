package kucoin

import (
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/Kucoin/kucoin-go-sdk"
)

func handleSDKErr(err error, res *kucoin.ApiResponse) error {

	if err != nil {
		return errors.Wrap(err, "kucoin-sdk", errors.ErrInternal)
	}

	if res != nil && res.Code != "200000" {
		return errors.Wrap(errors.NewMesssage(fmt.Sprintf("%s:%s", res.Message, res.Code)))
	}

	return nil

}
