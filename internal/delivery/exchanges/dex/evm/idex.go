package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IDex interface {
	Pair(in, out types.Token) (*types.Pair, error)
	TxData(in, out types.Token, sender, receiver common.Address, amount *big.Int) ([]byte, error)
	Router() common.Address
}
