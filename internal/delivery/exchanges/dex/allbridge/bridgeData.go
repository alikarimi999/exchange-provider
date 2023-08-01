package allbridge

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/networks/evm"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (ex *exchange) getBridgeInput(o *types.Order, r *types.Route, n types.Network, in,
	out *entity.Token, afterSwap bool, amountIn, feeAmount float64) (*evm.Input, error) {

	if n.Type() == types.EvmNetwork {
		in, err := ex.tl.getTokenInfo(in.Id)
		if err != nil {
			return nil, err
		}
		out, err := ex.tl.getTokenInfo(out.Id)
		if err != nil {
			return nil, err
		}
		messenger := chooseMessenger(in.TransferTime[out.Network])
		bf, err := getBridgeFee(in.ChainId, out.ChainId, messenger)
		if err != nil {
			return nil, err
		}

		var receiver string
		if len(o.Steps) == 2 {
			receiver = ex.cfg.Networks.network(o.Steps[1].Routes[0].In.Network).MainContract
		} else {
			receiver = o.Receiver
		}
		o.Steps[0].Receiver = receiver
		n, _ := new(big.Int).SetString(o.Nonce, 10)
		return &evm.Input{
			In:        in,
			Out:       out,
			Sender:    common.HexToAddress(o.Sender),
			Receiver:  common.HexToAddress(receiver),
			Nonce:     n,
			AfterSwap: afterSwap,
			AmountIn:  amountIn,
			FeeAmount: feeAmount,
			Messenger: int64(messenger),
			BridgeFee: bf,
		}, nil
	}
	return nil, errors.Wrap(errors.ErrUnknown)
}
