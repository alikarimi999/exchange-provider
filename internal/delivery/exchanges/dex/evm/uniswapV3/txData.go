package uniswapV3

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3/contracts"
	"exchange-provider/internal/entity"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *dex) TxData(in, out *entity.Token, sender, receiver common.Address,
	amount *big.Int, fee int64) ([]byte, error) {

	data := [][]byte{}
	abi, err := contracts.RouteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	var rec common.Address
	if out.Native {
		rec = p.router
	} else {
		rec = receiver
	}
	params := contracts.IV3SwapRouterExactInputSingleParams{
		TokenIn:           common.HexToAddress(in.Address),
		TokenOut:          common.HexToAddress(out.Address),
		Fee:               big.NewInt(fee),
		Recipient:         rec,
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	input, err := abi.Pack("exactInputSingle", params)
	if err != nil {
		return nil, err
	}

	data = append(data, input)

	if out.Native {
		input, err := abi.Pack("unwrapWETH9", common.Big0, receiver)
		if err != nil {
			return nil, err
		}
		data = append(data, input)
	}

	deadline := big.NewInt(time.Now().Add(time.Minute * time.Duration(15)).Unix())
	return abi.Pack("multicall0", deadline, data)

}
