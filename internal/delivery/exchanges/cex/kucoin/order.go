package kucoin

import (
	"exchange-provider/internal/delivery/exchanges/cex/kucoin/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"time"
)

type NewOrderData struct {
	UserId            string
	In                entity.TokenId
	Out               entity.TokenId
	AmountIn          float64
	Es                *entity.EstimateAmount
	SenderAddress     entity.Address
	WithdrawalAddress entity.Address
}

func (k *kucoinExchange) NewOrder(data interface{}, api *entity.APIToken) (entity.Order, error) {
	d := data.(*NewOrderData)
	p := d.Es.P
	var (
		dt *Token
	)
	if p.T1.Id.String() == d.In.String() {
		dt = p.T1.ET.(*Token)
	} else {
		dt = p.T2.ET.(*Token)
	}

	if err := k.getAddress(dt); err != nil {
		return nil, errors.Wrap(errors.ErrUnknown,
			errors.NewMesssage("there is a problem about this pair right now"))
	}

	t := time.Now()
	o := &types.Order{
		Status: types.OCreated,
		UserID: d.UserId,
		ExNid:  k.NID(),
		In:     d.In,
		Out:    d.Out,

		ApiKey: api.Id,
		BusId:  api.BusId,
		Level:  api.Level,

		SetAmountIn:               d.AmountIn,
		InitialPrice:              d.Es.Price,
		EstimateAmountOut:         d.Es.AmountOut,
		EstimateExchangeFeeAmount: d.Es.ExchangeFeeAmount,
		EstimateFeeAmount:         d.Es.FeeAmount,
		SpreadRate:                d.Es.SpreadRate,
		FeeRate:                   d.Es.FeeRate,
		ExchangeFee:               p.ExchangeFee,
		FeeAndSpreadCurrency:      d.Out,

		Deposit: types.Deposit{
			Address: entity.Address{
				Addr: dt.DepositAddress,
				Tag:  dt.DepositTag,
			},
		},
		Swaps: make(map[int]*types.Swap),
		Withdrawal: &types.Withdrawal{
			Address: d.WithdrawalAddress,
		},
		Sender:    d.SenderAddress,
		CreatedAT: t.Unix(),
		UpdatedAt: t.Unix(),
		ExpireAt:  t.Add(2 * time.Hour).Unix(),
	}

	if p.EP != nil && p.EP.(*ExchangePair).HasIntermediaryCoin {
		ep := p.EP.(*ExchangePair)
		o.Swaps[0] = &types.Swap{
			In:   d.In,
			Out:  ep.IC1.toId(),
			Side: sellSide,
		}
		o.Swaps[1] = &types.Swap{
			In:   ep.IC1.toId(),
			Out:  d.Out,
			Side: buySide,
		}
	} else {
		var side string
		if p.T1.String() == d.In.String() {
			side = sellSide
		} else {
			side = buySide
		}
		o.Swaps[0] = &types.Swap{
			Side: side,
			In:   d.In,
			Out:  d.Out,
		}
	}

	return o, nil
}
