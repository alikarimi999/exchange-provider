package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Input struct {
	In, Out             *types.TokenInfo
	Sender, Receiver    common.Address
	Nonce               *big.Int
	AfterSwap           bool
	AmountIn, FeeAmount float64
	Messenger           int64
	BridgeFee           *big.Int
}

func (d *net) BridgeData(input interface{}) ([]byte, error) {

	i := input.(*Input)
	if i.AfterSwap {
		i.AmountIn = 0
		i.FeeAmount = 0
	}

	allbridgeInput := allBridgeInputBytes(common.HexToAddress(i.In.TokenAddress),
		common.HexToAddress(i.Out.TokenAddress), i.Receiver, i.Nonce, int64(i.Out.ChainId), i.Messenger, d.allbridgeContract)

	amIn, _ := new(big.Float).Mul(big.NewFloat(i.AmountIn), big.NewFloat(math.Pow10(i.In.Decimals))).Int(nil)
	feeAm, _ := new(big.Float).Mul(big.NewFloat(i.FeeAmount), big.NewFloat(math.Pow10(i.In.Decimals))).Int(nil)
	data, err := bridgeInputBytes(d.ourAllBridgeContract, common.HexToAddress(i.In.TokenAddress),
		i.Sender, i.BridgeFee, amIn, feeAm, i.AfterSwap, allbridgeInput)
	if err != nil {
		return nil, err
	}

	b, err := bridgeArgs.Pack(data)
	if err != nil {
		return nil, err
	}

	sig, err := sign(b, d.prvKey)
	if err != nil {
		return nil, err
	}

	return d.abi.Pack("Bridge", data, sig)
}
