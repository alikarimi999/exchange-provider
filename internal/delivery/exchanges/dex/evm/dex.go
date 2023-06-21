package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type evmDex struct {
	*Config
	dex   IDex
	pairs entity.PairsRepo
	repo  entity.OrderRepo
	abi   *abi.ABI
	l     logger.Logger
}

func NewEvmDex(cfg *Config, repo entity.OrderRepo, pairs entity.PairsRepo,
	l logger.Logger) (entity.EVMDex, error) {
	var (
		d   IDex
		err error
	)

	abi, err := contracts.ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ex := &evmDex{
		Config: cfg,
		pairs:  pairs,
		repo:   repo,
		abi:    abi,
		l:      l,
	}

	k, err := crypto.HexToECDSA(cfg.HexKey)
	if err != nil {
		return nil, err
	}
	ex.Config.privateKey = k
	ex.contractAddress = common.HexToAddress(ex.Contract)
	ex.swapperAddress = common.HexToAddress(ex.Swapper)
	ex.WrappedNativeToken = fmt.Sprintf("W%s", ex.NativeToken)

	if err := ex.checkProviders(); err != nil {
		return nil, err
	}

	switch cfg.Version {
	case 3:
		d, err = uniswapV3.NewUniswapV3Dex(ex.NID(), ex.Network, ex.NativeToken, ex.Swapper,
			ex.PriceProvider, ex.Contract, ex.ChainId, ex.privateKey, ex.providers, l)

	case 2:
		d, err = uniswapV2.NewUniswapV2Dex(ex.NID(), ex.Network, ex.NativeToken, ex.Swapper,
			ex.PriceProvider, ex.Contract, ex.ChainId, ex.privateKey, ex.providers, l)
	default:
		err = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("uniswap only support version '2' and '3'"))
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	ex.dex = d
	return ex, nil
}

func (d *evmDex) Name() string {
	return d.Config.Name
}

func (d *evmDex) Id() uint {
	return d.Config.Id
}

func (d *evmDex) NID() string {
	return fmt.Sprintf("%s-%d", d.Name(), d.Id())
}

func (d *evmDex) Chain() string {
	return d.Config.Network
}

func (d *evmDex) Type() entity.ExType {
	return entity.EvmDEX
}

func (d *evmDex) Configs() interface{} {
	return d.Config
}

func (d *evmDex) EnableDisable(enable bool) {
	d.Enable = enable
}
func (d *evmDex) IsEnable() bool {
	return d.Enable
}

func (d *evmDex) Remove() {}
