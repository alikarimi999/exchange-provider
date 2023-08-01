package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts/erc20"
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

type exchange struct {
	cfg   *Config
	dex   IDex
	pairs entity.PairsRepo
	repo  entity.OrderRepo
	abi   *abi.ABI
	erc20 *abi.ABI

	l logger.Logger
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

	erc20, err := erc20.ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	ex := &exchange{
		cfg:   cfg,
		pairs: pairs,
		repo:  repo,
		abi:   abi,
		erc20: erc20,
		l:     l,
	}

	k, err := crypto.HexToECDSA(cfg.HexKey)
	if err != nil {
		return nil, err
	}
	ex.cfg.prvKey = k
	ex.cfg.contractAddress = common.HexToAddress(ex.cfg.Contract)
	ex.cfg.swapperAddress = common.HexToAddress(ex.cfg.Swapper)
	ex.cfg.WrappedNativeToken = fmt.Sprintf("W%s", ex.cfg.NativeToken)

	if err := ex.checkProviders(); err != nil {
		return nil, err
	}

	switch cfg.Version {
	case 3:
		d, err = uniswapV3.NewUniswapV3Dex(ex.NID(), ex.cfg.Network, ex.cfg.NativeToken, ex.cfg.Swapper,
			ex.cfg.PriceProvider, ex.cfg.Contract, ex.cfg.ChainId, ex.cfg.prvKey, ex.cfg.providers, l)

	case 2:
		d, err = uniswapV2.NewUniswapV2Dex(ex.NID(), ex.cfg.Network, ex.cfg.NativeToken, ex.cfg.Swapper,
			ex.cfg.PriceProvider, ex.cfg.Contract, ex.cfg.ChainId, ex.cfg.prvKey, ex.cfg.providers, l)
	default:
		err = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("only support version '2' and '3'"))
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	ex.dex = d
	return ex, nil
}

func (d *exchange) Name() string {
	return d.cfg.Name
}

func (d *exchange) Id() uint {
	return d.cfg.Id
}

func (d *exchange) NID() string {
	return fmt.Sprintf("%s-%d", d.Name(), d.Id())
}

func (d *exchange) Network() string  { return d.cfg.Network }
func (d *exchange) Standard() string { return d.cfg.TokenStandard }
func (d *exchange) Type() entity.ExType {
	return entity.EvmDEX
}

func (d *exchange) Configs() interface{} {
	return d.cfg
}

func (d *exchange) EnableDisable(enable bool) {
	d.cfg.Enable = enable
}
func (d *exchange) IsEnable() bool {
	return d.cfg.Enable
}

func (d *exchange) Remove() {}
