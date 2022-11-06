package types

import (
	"exchange-provider/internal/entity"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Dex interface {
	PairWithPrice(bt, qt Token) (*Pair, error)
	SetBestPrice(bt, qt Token) (*Pair, error)
	// remember: set Withdrawal.Unwrapped for panckakeSwapp true
	Swap(o *entity.UserOrder, tIn, tOut Token, value string, source, dest common.Address) (*types.Transaction, *Pair, error)
}
