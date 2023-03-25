package evm

import (
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IDex interface {
	EstimateAmountOut(in, out *entity.Token, amountIn float64) (amountOut float64, fee uint64, err error)
	TxData(in, out *entity.Token, sender, receiver common.Address, amount *big.Int, fee int64) ([]byte, error)
	Router() common.Address
}
