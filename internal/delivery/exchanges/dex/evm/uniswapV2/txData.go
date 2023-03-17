package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *dex) TxData(in, out *types.Token, sender, receiver common.Address,
	amount *big.Int, fee int64) ([]byte, error) {

	abi, err := contracts.ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()

	if in.Native {
		return abi.Pack("swapExactETHForTokens", common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	} else if out.Native {
		return abi.Pack("swapExactTokensForETH", amount, common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	} else {
		return abi.Pack("swapExactTokensForTokens", amount, common.Big0,
			[]common.Address{in.Address, out.Address}, receiver, big.NewInt(d))
	}
}
