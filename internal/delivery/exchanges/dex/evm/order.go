package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

type NewOrderData struct {
	UserId   string
	In       entity.TokenId
	Out      entity.TokenId
	Es       *entity.EstimateAmount
	Sender   common.Address
	Reciever common.Address
	AmountIn float64
}

func (ex *evmDex) NewOrder(data interface{}, api *entity.APIToken) (entity.Order, error) {

	d := data.(*NewOrderData)
	p, err := ex.pairs.Get(ex.Id(), d.In.String(), d.Out.String())
	if err != nil {
		return nil, err
	}

	if !p.Enable {
		return nil, errors.Wrap(errors.ErrNotFound, errors.NewMesssage("pair is not enable right now"))
	}

	var in, out *entity.Token
	if p.T1.Id.String() == d.In.String() {
		in = p.T1
		out = p.T2
	} else {
		in = p.T2
		out = p.T1
	}
	o := &types.Order{
		UserID: d.UserId,
		Status: entity.OCreated,
		ExNid:  ex.NID(),
		In:     in.Id,
		Out:    out.Id,

		ApiKey: api.Id,
		BusId:  api.BusId,
		Level:  api.Level,

		Sender:   d.Sender,
		Receiver: d.Reciever,
		AmountIn: d.AmountIn,
	}

	o.NeedApprove, err = ex.needApproval(in, o.Sender, o.AmountIn)
	if err != nil {
		return nil, err
	}

	return o, nil
}
