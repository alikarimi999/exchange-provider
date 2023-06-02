package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/pkg/errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
)

func (k *exchange) swap(o *types.Order, bc, qc *Token, index int) error {

	var size, funds string
	side := o.Swaps[index].Side
	if side == sellSide {
		size = big.NewFloat(o.Swaps[index].InAmountRequested).Text('f', 18)
	} else {
		funds = big.NewFloat(o.Swaps[index].InAmountRequested).Text('f', 18)
	}

	req := k.createOrderRequest(bc, qc, side, size, funds,
		fmt.Sprintf("%s-%d", o.ID().String(), index))

	if side == sellSide {
		o.Swaps[index].InAmountRequested, _ = strconv.ParseFloat(req.Size, 64)
	} else {
		o.Swaps[index].InAmountRequested, _ = strconv.ParseFloat(req.Funds, 64)
	}

	// res, err := k.writeApi.InnerTransferV2(uuid.New().String(), in.Symbol, "main", "trade", amount)
	// if err = handleSDKErr(err, res); err != nil {
	// 	return "", errors.Wrap(err, op, errors.ErrBadRequest)
	// }

	// k.l.Debug(agent, fmt.Sprintf("%s %s transferred from main account to trade account",
	// 	amount, in.Symbol))

	for i := 0; i <= 10; i++ {
		res, err := k.writeApi.CreateOrder(req)
		if err = handleSDKErr(err, res); err != nil {
			if i == 10 {
				return err
			}
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return nil

}

func (k *exchange) trackSwap(o *types.Order, bc, qc *Token, index int) error {
	s := o.Swaps[index]

	var (
		resp *kucoin.ApiResponse
		err  error
	)

	for i := 0; i <= 10; i++ {
		resp, err = k.readApi.OrderByClient(fmt.Sprintf("%s-%d", o.ID().String(), index))
		if err == nil && resp.Code == "400100" && i == 0 {
			time.Sleep(2 * time.Second)
			continue
		} else if err == nil && resp.Code == "400100" && i == 1 {
			return errors.Wrap(errors.ErrNotFound)
		}
		if err = handleSDKErr(err, resp); err != nil {
			if i == 10 {
				return err
			}
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	order := &kucoin.OrderModel{}
	if err := resp.ReadData(order); err != nil {
		return err
	}

	fee, _ := strconv.ParseFloat(order.Fee, 64)
	if order.Side == sellSide {
		s.InAmountExecuted, _ = strconv.ParseFloat(order.DealSize, 64)
		outAmountS := trim(order.DealFunds, qc.OrderPrecision)
		outAmountF, _ := strconv.ParseFloat(outAmountS, 64)
		s.OutAmount = outAmountF
	} else {
		amIn, _ := strconv.ParseFloat(order.DealFunds, 64)
		s.InAmountExecuted = amIn
		s.OutAmount, _ = strconv.ParseFloat(order.DealSize, 64)
	}

	s.KucoinFee = fee
	s.FeeCurrency = order.FeeCurrency
	return nil
}
