package uniswapV3

import (
	"crypto/ecdsa"
	ts "exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3/contracts"
	"exchange-provider/pkg/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type dex struct {
	id          string
	network     string
	chainId     *big.Int
	nativeToken string
	router      common.Address
	contract    common.Address
	factory     common.Address
	prvKey      *ecdsa.PrivateKey
	ps          []*ts.EthProvider
	l           logger.Logger
}

func NewUniswapV3Dex(nid string, network, nativeToken, router, contract string, chainId int64,
	prvKey *ecdsa.PrivateKey, ps []*ts.EthProvider, l logger.Logger) (*dex, error) {

	d := &dex{
		id:          nid,
		network:     network,
		chainId:     big.NewInt(chainId),
		nativeToken: nativeToken,
		router:      common.HexToAddress(router),
		contract:    common.HexToAddress(contract),
		prvKey:      prvKey,
		ps:          ps,
		l:           l,
	}
	r, err := contracts.NewRouter(d.router, d.provider())
	if err != nil {
		return nil, err
	}
	f, err := r.Factory(nil)
	if err != nil {
		return nil, err
	}
	d.factory = f
	return d, nil
}

func (d *dex) Router() common.Address { return d.router }
