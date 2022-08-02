package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) createOrderRequest(o *entity.UserOrder) (*kucoin.CreateOrderModel, error) {

	p, err := k.exchangePairs.get(o.BC, o.QC)
	if err != nil {
		return nil, err
	}

	return &kucoin.CreateOrderModel{
		ClientOid: uuid.New().String(),
		Symbol:    p.symbol,
		Type:      "market",
		Side:      o.Side,
		Size:      trim(o.Size, p.bc.orderPrecision),
		Funds:     trim(o.Funds, p.qc.orderPrecision),
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
	fmt.Printf("trim %s to %s\n", s, result)
	return result
}
