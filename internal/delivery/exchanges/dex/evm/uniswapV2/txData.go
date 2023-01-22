package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2/contracts"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *dex) TxData(in, out ts.Token, sender, receiver common.Address,
	amount *big.Int) ([]byte, error) {

	abi, err := contracts.ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()

	if in.IsNative() {
		return abi.Pack("swapExactETHForTokens", common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	} else if out.IsNative() {
		return abi.Pack("swapExactTokensForETH", amount, common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	} else {
		return abi.Pack("swapExactTokensForTokens", amount, common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	}
}
