package kucoin

import (
	"fmt"
	"order_service/internal/entity"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *kucoinExchange) createOrderRequest(o *entity.UserOrder, sr entity.PairConfigs) (*kucoin.CreateOrderModel, error) {

	p, err := k.exchangePairs.get(o.BC, o.QC)
	if err != nil {
		return nil, err
	}

	if o.Side == "buy" {
		vol, rate, err := sr.ApplySpread(o.BC, o.QC, o.Size)
		if err != nil {
			return nil, err
		}
		o.SpreadRate = rate
		return &kucoin.CreateOrderModel{
			Symbol: p.Symbol,
			Side:   o.Side,
			Type:   "market",
			Size:   trim(vol, p.Bc.orderPrecision),
		}, nil
	} else {
		return &kucoin.CreateOrderModel{
			Symbol: p.Symbol,
			Side:   o.Side,
			Type:   "market",
			Funds:  trim(o.Funds, p.Qc.orderPrecision),
		}, nil
	}

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
