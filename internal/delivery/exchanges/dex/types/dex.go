package types

import (
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Dex interface {
	PairWithPrice(bt, qt Token) (*Pair, error)
	Pair(bt, qt Token) (*Pair, error)

	Swap(o *entity.UserOrder, tIn, tOut Token, value string, source, dest common.Address) (tx *types.Transaction, nonce *big.Int, err error)
	ParseSwapLogs(o *entity.UserOrder, tx *types.Transaction, pair *Pair, receipt *types.Receipt) (amount, fee string, err error)
}
