package binance

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/cex/binance/types"
	"exchange-provider/internal/entity"
	"math/big"
	"time"

	"github.com/adshao/go-binance/v2"
)

func (ex *exchange) withdrawal(o *types.Order, wc *Token, p *entity.Pair) {

	s := o.Swaps[len(o.Swaps)-1]
	// amountOut := applyFee(s.OutAmount, feeRate)
	amountIn := o.Swaps[0].InAmountExecuted
	amountOut := s.OutAmount
	var price float64
	if wc.Coin == p.T2.ET.(*Token).Coin {
		o.SpreadAmount = (amountOut * o.SpreadRate)
		amountOut = amountOut - o.SpreadAmount
		price = s.OutAmount / amountIn
	} else {
		price = amountIn / s.OutAmount
		o.SpreadAmount = amountOut - (amountIn / (price + (price * o.SpreadRate)))
		amountOut = amountOut - o.SpreadAmount
	}

	amountOut = amountOut - o.ExchangeFeeAmount
	o.FeeAmount = amountOut * o.FeeRate
	o.Withdrawal.Amount = amountOut - o.FeeAmount
	o.ExecutedPrice = price

	vol := trim(big.NewFloat(o.Withdrawal.Amount).Text('f', 18), wc.WithdrawalPrecision)

	var (
		err error
		res *binance.CreateWithdrawResponse
	)
	for i := 0; i <= 10; i++ {
		res, err = ex.c.NewCreateWithdrawService().WithdrawOrderID(o.ID().String()).
			Address(o.Withdrawal.Addr).AddressTag(o.Withdrawal.Tag).Amount(vol).
			Coin(wc.Coin).Network(wc.Network).Do(context.Background())
		if err != nil {
			if i == 10 {
				o.Status = types.OWithdrawalFailed
				o.FailedDesc = err.Error()
				return
			}
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	o.Withdrawal.Id = res.ID
	o.Withdrawal.Amount = o.Withdrawal.Amount - wc.MinWithdrawalFee
	o.Withdrawal.BinanceFee = wc.MinWithdrawalFee
	o.Status = types.OWithdrawalTracking
}
