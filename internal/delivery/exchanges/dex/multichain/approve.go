package multichain

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type approveList struct {
	mux  *sync.Mutex
	list []struct {
		token    common.Address
		spender  common.Address
		owner    common.Address
		approved bool
	}
}
