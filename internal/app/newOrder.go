package app

import (
	b "exchange-provider/internal/delivery/exchanges/cex/binance"
	k "exchange-provider/internal/delivery/exchanges/cex/kucoin"
	a "exchange-provider/internal/delivery/exchanges/dex/allbridge"
	e "exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

func (u *OrderUseCase) NewOrder(userId string, sender, refund, reciever entity.Address,
	in, out entity.TokenId, amount float64, lp uint, api *entity.APIToken) (entity.Order, error) {
	es, ex, aLPs, err := u.estimateAmountOut(in, out, amount, lp, api.Level, nil)
	if err != nil {
		return nil, err
	}

	o, err := u.newOrder(userId, sender, refund, reciever, in, out, amount, lp, api, es[0], ex[0])
	if err == nil {
		if err := u.repo.Add(o); err != nil {
			return nil, err
		}
		return o, nil
	}
	exculudeLPs := []uint{ex[0].Id()}
	for i := 0; i < len(aLPs); i++ {
		es0, ex0, _, err := u.estimateAmountOut(in, out, amount, lp, api.Level, exculudeLPs)
		if err != nil {
			return nil, err
		}

		o, err := u.newOrder(userId, sender, refund, reciever, in, out, amount, lp, api, es0[0], ex0[0])
		if err == nil {
			if err := u.repo.Add(o); err != nil {
				return nil, err
			}
			return o, nil
		}
		exculudeLPs = append(exculudeLPs, ex0[0].Id())
	}
	return nil, errors.Wrap(errors.ErrInternal)
}

func (u *OrderUseCase) newOrder(userId string, sender, refund, reciever entity.Address,
	in, out entity.TokenId, amount float64, lp uint, api *entity.APIToken,
	es *entity.EstimateAmount, ex entity.Exchange) (entity.Order, error) {

	if refund.Addr == "" {
		refund = sender
	}

	var (
		d interface{}
	)

	switch ex.Type() {
	case entity.EvmDEX:
		d = &e.NewOrderData{
			UserId:   userId,
			In:       in,
			Out:      out,
			Es:       es,
			Sender:   common.HexToAddress(sender.Addr),
			Reciever: common.HexToAddress(reciever.Addr),
			AmountIn: amount,
		}
	case entity.CrossDex:
		d = &a.NewOrderData{
			UserId:   userId,
			In:       in,
			Out:      out,
			Es:       es,
			Sender:   sender,
			Reciever: reciever,
			AmountIn: amount,
		}
	case entity.CEX:
		switch ex.Name() {
		case "kucoin":
			d = &k.NewOrderData{
				UserId:            userId,
				In:                in,
				Out:               out,
				Es:                es,
				SenderAddress:     sender,
				WithdrawalAddress: reciever,
			}
		case "binance":
			d = &b.NewOrderData{
				UserId:            userId,
				In:                in,
				Out:               out,
				Es:                es,
				SenderAddress:     sender,
				WithdrawalAddress: reciever,
			}
		}
	}
	return ex.NewOrder(d, api)
}
