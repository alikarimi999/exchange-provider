package uniswapV2

import (
	"crypto/ecdsa"
	ts "exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2/contracts"
	"exchange-provider/pkg/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type dex struct {
	network     string
	chainId     *big.Int
	nativeToken string

	priceProvider common.Address
	router        common.Address
	contract      common.Address
	factory       common.Address
	abi           *abi.ABI
	prvKey        *ecdsa.PrivateKey
	ps            []*ts.EthProvider

	l logger.Logger
}

func NewUniswapV2Dex(nid string, network, nativeToken, router, priceProvide, contract string, chainId int64,
	prvKey *ecdsa.PrivateKey, ps []*ts.EthProvider, l logger.Logger) (*dex, error) {

	abi, err := contracts.ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	d := &dex{
		network:       network,
		chainId:       big.NewInt(chainId),
		nativeToken:   nativeToken,
		priceProvider: common.HexToAddress(priceProvide),
		router:        common.HexToAddress(router),
		contract:      common.HexToAddress(contract),
		abi:           abi,
		prvKey:        prvKey,
		ps:            ps,
		l:             l,
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
