package binance

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
)

func (ex *exchange) swap(o *types.Order, bc, qc *Token, index int) error {

	var (
		amount string
		res    *binance.CreateOrderResponse
		err    error
	)
	side := o.Swaps[index].Side
	if side == binance.SideTypeSell {
		amount = trim(big.NewFloat(o.Swaps[index].InAmountRequested).Text('f', 18), bc.OrderPrecision)
		res, err = sendOrderRequest(ex.c.NewCreateOrderService().Type("MARKET").Side(side).
			Symbol(bc.Coin + qc.Coin).Quantity(amount).NewClientOrderID(fmt.Sprintf("%s-%d", o.ID().String(), index)))
		if err != nil {
			return err
		}

		o.Swaps[index].OutAmount, _ = strconv.ParseFloat(res.CummulativeQuoteQuantity, 64)
		o.Swaps[index].InAmountExecuted, _ = strconv.ParseFloat(res.ExecutedQuantity, 64)
	} else {
		amount = trim(big.NewFloat(o.Swaps[index].InAmountRequested).Text('f', 18), qc.OrderPrecision)
		res, err = sendOrderRequest(ex.c.NewCreateOrderService().Type("MARKET").Side(side).
			Symbol(bc.Coin + qc.Coin).QuoteOrderQty(amount).NewClientOrderID(fmt.Sprintf("%s-%d", o.ID().String(), index)))
		if err != nil {
			return err
		}

		o.Swaps[index].InAmountExecuted, _ = strconv.ParseFloat(res.CummulativeQuoteQuantity, 64)
		o.Swaps[index].OutAmount, _ = strconv.ParseFloat(res.ExecutedQuantity, 64)
	}

	o.Swaps[index].InAmountRequested, _ = strconv.ParseFloat(amount, 64)

	for _, f := range res.Fills {
		fAmount, _ := strconv.ParseFloat(f.Commission, 64)
		o.Swaps[index].BinanceFees = append(o.Swaps[index].BinanceFees, types.BinanceFee{
			Coin:   f.CommissionAsset,
			Amount: fAmount,
		})
	}

	var out string
	if side == binance.SideTypeSell {
		out = qc.Coin
	} else {
		out = bc.Coin
	}

	for _, bf := range o.Swaps[index].BinanceFees {
		if bf.Coin == out {
			o.Swaps[index].OutAmount -= bf.Amount
		}
	}
	return nil
}

func sendOrderRequest(req *binance.CreateOrderService) (*binance.CreateOrderResponse, error) {
	for i := 0; i <= 10; i++ {
		res, err := req.Do(context.Background())
		if err != nil {
			if i == 10 {
				return nil, err
			}
			time.Sleep(5 * time.Second)
			continue
		}

		if res.Status != binance.OrderStatusTypeFilled {
			return nil, fmt.Errorf("binance order status is '%s'", res.Status)
		}
		return res, nil
	}
	return nil, fmt.Errorf("it never happens")
}

func trim(s string, l int) string {
	if s == "" {
		return s
	}

	ss := strings.Split(s, ".")
	if l == 0 {
		return ss[0]
	}
	var result string
	if len(ss) == 2 {
		if len(ss[1]) > l {
			result = ss[0] + "." + ss[1][:l]
		} else {
			result = s
		}

	} else {
		result = ss[0]
	}
	return result
}
