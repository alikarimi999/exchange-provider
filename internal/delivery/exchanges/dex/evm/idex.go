package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IDex interface {
	SaveAvailablePairs(ps []types.Pair, file string)
	Prices(ps []*entity.Pair) error
	EstimateAmountOut(in, out *types.Token, amountIn float64) (amountOut float64, fee uint64, err error)
	TxData(in, out *types.Token, sender, receiver common.Address, amount *big.Int, fee int64) ([]byte, error)
	Router() common.Address
}
