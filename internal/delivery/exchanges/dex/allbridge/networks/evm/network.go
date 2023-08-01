package evm

import (
	"context"
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type net struct {
	network              string
	mainContract         common.Address
	ourAllBridgeContract common.Address
	allbridgeContract    common.Address

	provider bind.ContractBackend
	prvKey   *ecdsa.PrivateKey
	chainId  int64
	abi      *abi.ABI
}

func NewEvmNetwork(nid, network, ourAllbridge, allbridge, mainContract string, c *ethclient.Client,
	prvKey *ecdsa.PrivateKey) (types.Network, error) {

	abi, err := contracts.ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	i, err := c.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &net{
		network:              network,
		mainContract:         common.HexToAddress(mainContract),
		ourAllBridgeContract: common.HexToAddress(ourAllbridge),
		allbridgeContract:    common.HexToAddress(allbridge),
		provider:             c,
		prvKey:               prvKey,
		abi:                  abi,
		chainId:              i.Int64(),
	}, nil
}

func (n *net) Type() types.NetworkType { return types.EvmNetwork }
