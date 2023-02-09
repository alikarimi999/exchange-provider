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

	IDex
	pairsRepo entity.PairRepo
	ts        *tokens
	v         *viper.Viper
	l         logger.Logger
}

func NewEvmDex(cfg *Config, pairs entity.PairRepo, v *viper.Viper,
	l logger.Logger, readConfig bool) (entity.EVMDex, error) {
	agent := "NewEvmDex"

	var (
		d   IDex
		err error
	)

	if err := cfg.Validate(readConfig); err != nil {
		return nil, err
	}

	ex := &EvmDex{
		Config:    cfg,
		pairsRepo: pairs,
		v:         v,
		l:         l,
	}
	ex.Config.Id = ex.Config.Name + "-" + ex.Config.Network

	k, err := crypto.HexToECDSA(cfg.HexKey)
	if err != nil {
		return nil, err
	}
	ex.Config.privateKey = k

	if readConfig {
		ex.l.Debug(agent, fmt.Sprintf("Retrieving `%s` data ...", ex.Id()))

		ex.NativeToken = ex.v.GetString(fmt.Sprintf("%s.native_token", ex.Id()))
		ex.TokenStandard = ex.v.GetString(fmt.Sprintf("%s.token_standard", ex.Id()))
		ex.PairsFile = ex.v.GetString(fmt.Sprintf("%s.pairs_file", ex.Id()))
		ex.TokensFile = ex.v.GetString(fmt.Sprintf("%s.tokens_file", ex.Id()))
		ex.Contract = ex.v.GetString(fmt.Sprintf("%s.contract", ex.Id()))
		ex.contractAddress = common.HexToAddress(ex.Contract)
		ex.Swapper = ex.v.GetString(fmt.Sprintf("%s.swapper", ex.Id()))
		ex.swapperAddress = common.HexToAddress(ex.Swapper)

		i := ex.v.Get(fmt.Sprintf("%s.providers", ex.Id()))
		if i == nil {
			return nil, errors.New("no provider available in config file")
		}

		psi := i.(map[string]interface{})
		for _, v := range psi {
			ex.Providers = append(ex.Providers, v.(string))
		}

		if err := ex.checkProviders(); err != nil {
			return nil, err
		}

	} else {
		ex.contractAddress = common.HexToAddress(ex.Contract)
		ex.swapperAddress = common.HexToAddress(ex.Swapper)

		ex.v.Set(fmt.Sprintf("%s.native_token", ex.Id()), ex.NativeToken)
		ex.v.Set(fmt.Sprintf("%s.token_standard", ex.Id()), ex.TokenStandard)
		ex.v.Set(fmt.Sprintf("%s.pairs_file", ex.Id()), ex.PairsFile)
		ex.v.Set(fmt.Sprintf("%s.tokens_file", ex.Id()), ex.TokensFile)
		ex.v.Set(fmt.Sprintf("%s.contract", ex.Id()), ex.Contract)
		ex.v.Set(fmt.Sprintf("%s.swapper", ex.Id()), ex.Swapper)

		if err := ex.checkProviders(); err != nil {
			return nil, err
		}

		for i, p := range ex.providers {
			ex.v.Set(fmt.Sprintf("%s.providers.%d", ex.Id(), i), p.URL)
		}
		if err := ex.v.WriteConfig(); err != nil {
			return nil, err
		}
	}
	ex.WrappedNativeToken = fmt.Sprintf("W%s", ex.NativeToken)
	if err := ex.retreiveTokens(); err != nil {
		return nil, err
	}

	switch ex.Name() {
	case "uniswapv3":
		d, err = uniswapV3.NewUniswapV3Dex(ex.Id(), ex.Network, ex.NativeToken, ex.Swapper,
			ex.Contract, ex.ChainId, ex.privateKey, ex.providers, l)

	case "uniswapv2", "panckakeswapv2":
		d, err = uniswapV2.NewUniswapV2Dex(ex.Id(), ex.Network, ex.NativeToken, ex.Swapper, ex.Contract,
			ex.ChainId, ex.privateKey, ex.providers, l)
	default:
		err = errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("unsupported '%s'", ex.Name())))
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	ex.IDex = d
	ex.findAllPairs()
	return ex, nil
}

func (d *EvmDex) Name() string {
	return d.Config.Name
}

func (d *EvmDex) Id() string {
	return d.Config.Id
}

func (d *EvmDex) Type() entity.ExType {
	return entity.EvmDEX
}

func (d *EvmDex) Configs() interface{} {
	return d.Config
}

func (d *EvmDex) Remove() {
	d.pairsRepo.RemoveExchange(d.Id())
}
