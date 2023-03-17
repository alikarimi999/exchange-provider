package uniswapV2

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2/contracts"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/logger"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type dex struct {
	id          uint
	network     string
	chainId     *big.Int
	nativeToken string

	router   common.Address
	contract common.Address
	factory  common.Address

	prvKey *ecdsa.PrivateKey
	ps     []*ts.EthProvider

	l logger.Logger
}

func NewUniswapV2Dex(id uint, network, nativeToken, router, contract string, chainId int64,
	prvKey *ecdsa.PrivateKey, ps []*ts.EthProvider, l logger.Logger) (*dex, error) {

	d := &dex{
		id:          id,
		network:     network,
		chainId:     big.NewInt(chainId),
		nativeToken: nativeToken,
		router:      common.HexToAddress(router),
		contract:    common.HexToAddress(contract),
		prvKey:      prvKey,
		ps:          ps,
		l:           l,
	}

	r, err := contracts.NewContract(d.router, d.provider())
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

func (d *dex) agent(fn string) string {
	return fmt.Sprintf("%s.%s", d.id, fn)
}
