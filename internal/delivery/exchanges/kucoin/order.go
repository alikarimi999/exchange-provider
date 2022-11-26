package kucoin

import (
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) createOrderRequest(p *pair, side, size, funds string) (*kucoin.CreateOrderModel, error) {
	return &kucoin.CreateOrderModel{
		ClientOid: uuid.New().String(),
		Symbol:    p.Symbol(),
		Side:      side,
		Type:      "market",
		Size:      trim(size, p.BC.orderPrecision),
		Funds:     trim(funds, p.QC.orderPrecision),
	}, nil

}

// trim returns a string with the given precision.
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
