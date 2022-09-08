package uniswapv3

import "github.com/ethereum/go-ethereum/core/types"

const (
	nonceLocked = iota + 1
	nonceReleased
)

type tx struct {
	*types.Transaction
	nonceStatus int
}
