package allbridge

import (
	"crypto/rand"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
)

type NewOrderData struct {
	UserId   string
	In       entity.TokenId
	Out      entity.TokenId
	Es       *entity.EstimateAmount
	Sender   entity.Address
	Reciever entity.Address
	AmountIn float64
}

func (ex *exchange) NewOrder(data interface{}, api *entity.APIToken) (entity.Order, error) {
	d := data.(*NewOrderData)
	p, err := ex.pairs.Get(ex.Id(), d.In.String(), d.Out.String())
	if err != nil {
		return nil, err
	}

	var in, out *entity.Token
	if p.T1.Id.String() == d.In.String() {
		in = p.T1
		out = p.T2
	} else {
		in = p.T2
		out = p.T1
	}
	t := time.Now()
	o := &types.Order{
		UserID: d.UserId,
		Status: entity.OCreated,
		ExNid:  ex.NID(),
		In:     in.Id,
		Out:    out.Id,

		ApiKey: api.Id,
		BusId:  api.BusId,
		Level:  api.Level,

		Sender:            d.Sender.Addr,
		Receiver:          d.Reciever.Addr,
		Steps:             make(map[int]*types.Step),
		AmountIn:          d.AmountIn,
		EstimateAmountOut: d.Es.AmountOut,
		FeeRate:           d.Es.FeeRate,
		FeeAmount:         d.Es.FeeAmount,
		ExchangeFee:       d.Es.ExchangeFee,
		ExchangeFeeAmount: d.Es.ExchangeFeeAmount,

		FeeCurrency: d.In,
		CreatedAT:   t.Unix(),
		UpdatedAt:   t.Unix(),
		ExpireAt:    t.Add(time.Hour).Unix(),
	}

	r := d.Es.Data.(steps)[0][0]
	p0, err := ex.pairs.Get(r.ex.Id(), r.in.String(), r.out.String())
	if err != nil {
		return nil, err
	}

	var In *entity.Token
	if p0.T1.String() == in.String() {
		In = p0.T1
	} else {
		In = p0.T2
	}

	approve, err := ex.ns[r.in.Network].NeedApproval(In, o.Sender, o.AmountIn)
	if err != nil {
		return nil, err
	}
	n, err := rand.Int(rand.Reader, math.MaxBig256)
	if err != nil {
		return nil, err
	}
	o.Nonce = n.String()
	for i, rs := range d.Es.Data.(steps) {
		for j, r := range rs {
			var needApproval bool
			if i == 0 && j == 0 {
				needApproval = approve
			}
			if j == 0 {
				o.Steps[i] = &types.Step{
					Status: entity.OCreated.String(),
					Routes: make(map[int]*types.Route),
				}
			}
			o.Steps[i].Routes[len(o.Steps[i].Routes)] = &types.Route{
				In:                r.in,
				Out:               r.out,
				AmountIn:          r.amountIn,
				EstimateAmountOut: r.amountOut,
				ExchangeNid:       r.ex.NID(),
				NeedApprove:       needApproval,
			}
		}
	}

	return o, nil

}
