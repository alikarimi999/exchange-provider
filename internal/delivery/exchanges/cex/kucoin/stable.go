package kucoin

import (
	"fmt"
	"strconv"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) stableTicker(bc, sc string) (float64, error) {
	res, err := k.readApi.TickerLevel1(bc + "-" + sc)
	if err = handleSDKErr(err, res); err != nil {
		return 0, err
	}
	tl := &kucoin.TickerLevel1Model{}
	if err := res.ReadData(tl); err != nil {
		return 0, err
	}

	if tl.Price == "" {
		return 0, fmt.Errorf("pair '%s/%s' not found in kucoin, use another stableToken", bc, sc)
	}
	return strconv.ParseFloat(tl.Price, 64)
}
