package uniswapv3

import "github.com/ethereum/go-ethereum/common"

func hashToAddress(h common.Hash) common.Address {
	return common.BytesToAddress(h[:])
}
