package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IDex interface {
	EstimateAmountOut(in, out *types.Token, amountIn float64) (amountOut float64, fee uint64, err error)
	TxData(in, out *types.Token, receiver common.Address, amount *big.Int, fee int64) ([]byte, error)
	Router() common.Address
}
