package uniswapV2

import (
	"crypto/ecdsa"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type dex struct {
	network     string
	chainId     *big.Int
	nativeToken string
	router      common.Address
	prvKey      *ecdsa.PrivateKey
	ps          []*ts.EthProvider
}

func NewUniswapV2Dex(network, nativeToken, router string, chainId int64,
	prvKey *ecdsa.PrivateKey, ps []*ts.EthProvider) (*dex, error) {

	return &dex{
		network:     network,
		chainId:     big.NewInt(chainId),
		nativeToken: nativeToken,
		router:      common.HexToAddress(router),
		prvKey:      prvKey,
		ps:          ps,
	}, nil
}

func (d *dex) Router() common.Address { return d.router }
