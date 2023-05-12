package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"math/big"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

func (k *kucoinExchange) withdrawal(o *types.Order, wc *Token, p *entity.Pair,
	withdrawalfromMain bool) {

	s := o.Swaps[len(o.Swaps)-1]
	// amountOut := applyFee(s.OutAmount, feeRate)
	amountIn := o.Swaps[0].InAmountExecuted
	amountOut := s.OutAmount
	var price float64
	if wc.Currency == p.T2.ET.(*Token).Currency {
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

	opts := make(map[string]string)
	opts["chain"] = wc.Chain
	opts["memo"] = o.Withdrawal.Tag
	opts["remark"] = o.ID().String()
	opts["feeDeductType"] = "INTERNAL"

	vol := trim(big.NewFloat(o.Withdrawal.Amount).Text('f', 18), wc.WithdrawalPrecision)
	var (
		res *kucoin.ApiResponse
		err error
	)

	if !withdrawalfromMain {
		if err := k.innerTransfer(wc.Currency, vol); err != nil {
			o.Status = types.OWithdrawalFailed
			o.FailedDesc = err.Error()
			return
		}
	}
	// then withdraw from main account
	for i := 0; i <= 10; i++ {
		res, err = k.writeApi.ApplyWithdrawal(wc.Currency, o.Withdrawal.Addr, vol, opts)
		if err = handleSDKErr(err, res); err != nil {
			if withdrawalfromMain && res != nil && res.Code == "400100" &&
				res.Message == "account.available.amount" {
				if err := k.innerTransfer(wc.Currency, vol); err != nil {
					o.Status = types.OWithdrawalFailed
					o.FailedDesc = err.Error()
					return
				}
				withdrawalfromMain = false
				continue
			}

			if i == 10 {
				o.Status = types.OWithdrawalFailed
				o.FailedDesc = err.Error()
				return
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	w := &kucoin.ApplyWithdrawalResultModel{}
	res.ReadData(w)
	o.Withdrawal.Id = w.WithdrawalId
	o.Status = types.OWithdrawalTracking
}

func (k *kucoinExchange) innerTransfer(currency, vol string) error {
	for i := 0; i <= 10; i++ {
		res, err := k.writeApi.InnerTransferV2(uuid.New().String(), currency, "trade", "main", vol)
		if err = handleSDKErr(err, res); err != nil {
			if i == 10 {
				return err
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	return nil
}
