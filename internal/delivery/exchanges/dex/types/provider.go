package types

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthProvider struct {
	URL string
	*ethclient.Client
}
