package app

import (
	k "exchange-provider/internal/delivery/exchanges/cex/kucoin"
	e "exchange-provider/internal/delivery/exchanges/dex/evm"
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common"
)

func (u *OrderUseCase) NewOrder(userId string, sender, refund, reciever entity.Address,
	in, out entity.TokenId, amount float64, lp uint, api *entity.APIToken) (entity.Order, error) {

	if refund.Addr == "" {
		refund = sender
	}

	var (
		d interface{}
	)

	es, ex, err := u.estimateAmountOut(in, out, amount, lp, api.Level)
	if err != nil {
		return nil, err
	}

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

	case entity.CEX:
		switch ex.Name() {
		case "kucoin":
			d = &k.NewOrderData{
				UserId:            userId,
				In:                in,
				Out:               out,
				AmountIn:          amount,
				Es:                es,
				SenderAddress:     sender,
				WithdrawalAddress: reciever,
			}
		}
	}
	o, err := ex.NewOrder(d, api)
	if err != nil {
		return nil, err
	}
	return o, u.repo.Add(o)
}
