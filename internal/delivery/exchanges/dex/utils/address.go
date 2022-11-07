package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

func HashToAddress(h common.Hash) common.Address {
	return common.BytesToAddress(h[:])
}
