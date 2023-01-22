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

	Swap(o *entity.CexOrder, tIn, tOut Token, value string,
		source, dest common.Address) (tx *types.Transaction, nonce *big.Int, err error)
	TrackSwap(o *entity.CexOrder, p *Pair, index int)
}
