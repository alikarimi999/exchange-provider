package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/logger"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

type EvmDex struct {
	*Config

	dex     IDex
	pairs   entity.PairsRepo
	version int8
	ts      *tokens
	v       *viper.Viper
	l       logger.Logger
}

func NewEvmDex(cfg *Config, pairs entity.PairsRepo, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.EVMDex, error) {
	var (
		d   IDex
		err error
	)

	if err := cfg.Validate(readConfig); err != nil {
		return nil, err
	}

	ex := &EvmDex{
		Config: cfg,
		ts:     newTokens(),
		pairs:  pairs,
		v:      v,
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

	if readConfig {
		ps := ex.pairs.GetAll(ex.Id())
		for _, p := range ps {
			if !ex.ts.exists(p.T1.String()) {
				ex.ts.add(p.T1)
			}
			if !ex.ts.exists(p.T2.String()) {
				ex.ts.add(p.T2)
			}
		}
	}
	if err := ex.checkProviders(); err != nil {
		return nil, err
	}

	switch cfg.Name {
	case "uniswapv3":
		ex.version = 3
		d, err = uniswapV3.NewUniswapV3Dex(ex.Id(), ex.Network, ex.NativeToken, ex.Swapper,
			ex.Contract, ex.ChainId, ex.privateKey, ex.providers, l)

	case "uniswapv2", "panckakeswapv2":
		ex.version = 2
		d, err = uniswapV2.NewUniswapV2Dex(ex.Id(), ex.Network, ex.NativeToken, ex.Swapper, ex.Contract,
			ex.ChainId, ex.privateKey, ex.providers, l)
	default:
		err = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("unsupported '%s'", ex.Name())))
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	ex.dex = d
	return ex, nil
}

func (d *EvmDex) Name() string {
	return d.Config.Name + "-" + d.Config.Network
}

func (d *EvmDex) Id() uint {
	return d.Config.Id
}

func (d *EvmDex) Chain() string {
	return d.Config.Network
}

func (d *EvmDex) Type() entity.ExType {
	return entity.EvmDEX
}

func (d *EvmDex) Configs() interface{} {
	return d.Config
}

func (d *EvmDex) Remove() {}
