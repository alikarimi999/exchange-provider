package uniswapV2

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *dex) TxData(in, out *types.Token, receiver common.Address,
	amount *big.Int, fee int64) ([]byte, error) {

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	inAddress := common.HexToAddress(in.ContractAddress)
	outAddress := common.HexToAddress(out.ContractAddress)
	if in.Native {
		return p.abi.Pack("swapExactETHForTokens", common.Big0,
			[]common.Address{inAddress, outAddress}, receiver, big.NewInt(d))
	} else if out.Native {
		return p.abi.Pack("swapExactTokensForETH", amount, common.Big0,
			[]common.Address{inAddress, outAddress}, receiver, big.NewInt(d))
	} else {
		return p.abi.Pack("swapExactTokensForTokens", amount, common.Big0,
			[]common.Address{inAddress, outAddress}, receiver, big.NewInt(d))
	}
}
