package kucoin

import (
	"order_service/internal/entity"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) createOrderRequest(from, to *entity.Coin, vol string) (*kucoin.CreateOrderModel, error) {

	p, err := k.exchangePairs.get(from, to)
	if err != nil {
		return nil, err
	}

	if from.Id == p.bc.Id {
		return &kucoin.CreateOrderModel{
			ClientOid: uuid.New().String(),
			Symbol:    p.symbol,
			Type:      "market",
			Side:      "sell",
			Size:      trim(vol, p.bc.precision),
		}, nil
	} else {
		return &kucoin.CreateOrderModel{
			ClientOid: uuid.New().String(),
			Symbol:    p.symbol,
			Type:      "market",
			Side:      "buy",
			Funds:     trim(vol, p.qc.precision),
		}, nil
	}

}

func trim(s string, l int) string {
	if s == "" || l == 0 {
		return s
	}

	ss := strings.Split(s, ".")
	var result string

	if len(ss) == 2 {
		if len(ss[1]) > l {
			result = ss[0] + "." + ss[1][:l]
		} else {
			result = s
		}

	} else {
		result = ss[0] + ".0"
	}
	return result
}
