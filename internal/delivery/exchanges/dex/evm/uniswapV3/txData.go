package uniswapV3

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3/contracts"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *dex) TxData(in, out ts.Token, sender, receiver common.Address,
	amount *big.Int) ([]byte, error) {

	pair, err := p.Pair(in, out)
	if err != nil {
		return nil, err
	}

	data := [][]byte{}
	abi, err := contracts.RouteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           in.Address,
		TokenOut:          out.Address,
		Fee:               pair.FeeTier,
		Recipient:         receiver,
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}
	input, err := abi.Pack("exactInputSingle", params)
	if err != nil {
		return nil, err
	}
	data = append(data, input)

	if out.IsNative() {
		input, err := abi.Pack("unwrapWETH9", common.Big0, receiver)
		if err != nil {
			return nil, err
		}
		data = append(data, input)
	}

	deadline := big.NewInt(time.Now().Add(time.Minute * time.Duration(15)).Unix())
	return abi.Pack("multicall0", deadline, data)

}
