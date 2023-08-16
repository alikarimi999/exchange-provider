package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV2"
	"exchange-provider/internal/delivery/exchanges/dex/evm/uniswapV3"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
	Id                 uint   `json:"id"`
	Enable             bool   `json:"enable"`
	Name               string `json:"name"`
	Version            uint   `json:"version"`
	NativeToken        string `json:"nativeToken"`
	WrappedNativeToken string `json:"wrappedNativeToken"`
	ChainId            int64  `json:"chainId"`
	TokenStandard      string `json:"tokenStandard"`
	Network            string `json:"network"`
	PriceProvider      string `json:"priceProvider"`
	Contract           string `json:"contract"`
	contractAddress    common.Address
	Swapper            string `json:"swapper"`
	swapperAddress     common.Address
	HexKey             string `json:"hexKey"`
	prvKey             *ecdsa.PrivateKey
	Providers          []string `json:"providers"`
	providers          []*types.EthProvider
	Message            string `json:"message"`
}

func (ex *exchange) UpdateConfigs(cfgi interface{}, store entity.ExchangeStore) error {
	cfg, ok := cfgi.(*Config)
	if !ok {
		return fmt.Errorf("invalid config")
	}

	if err := cfg.validate(); err != nil {
		return err
	}

	if cfg.Id != ex.cfg.Id {
		return errors.Wrap(errors.ErrBadRequest, "the id field is not mutable")
	}

	if cfg.Version != ex.cfg.Version {
		return fmt.Errorf("the version field is not mutable")
	}

	if cfg.Name != ex.cfg.Name {
		return fmt.Errorf("the name field is not mutable")
	}

	if cfg.Network != ex.cfg.Network {
		return fmt.Errorf("the network field is not mutable")
	}

	if cfg.NativeToken != ex.cfg.NativeToken {
		return fmt.Errorf("the nativeToken field is not mutable")
	}

	if cfg.TokenStandard != ex.cfg.TokenStandard {
		return fmt.Errorf("the standard field is not mutable")
	}

	k, err := crypto.HexToECDSA(cfg.HexKey)
	if err != nil {
		return err
	}
	cfg.prvKey = k
	cfg.Enable = ex.cfg.Enable
	cfg.contractAddress = common.HexToAddress(cfg.Contract)
	cfg.swapperAddress = common.HexToAddress(cfg.Swapper)
	cfg.WrappedNativeToken = fmt.Sprintf("W%s", cfg.NativeToken)

	ps, chainId, err := checkProviders(cfg.Providers)
	if err != nil {
		return err
	}

	cfg.providers = ps
	cfg.ChainId = chainId.Int64()
	var (
		dex IDex
	)

	switch cfg.Version {
	case 3:
		dex, err = uniswapV3.NewUniswapV3Dex(ex.NID(), cfg.Network, cfg.NativeToken, cfg.Swapper,
			cfg.PriceProvider, cfg.Contract, cfg.ChainId, cfg.prvKey, cfg.providers, ex.l)

	case 2:
		dex, err = uniswapV2.NewUniswapV2Dex(ex.NID(), cfg.Network, cfg.NativeToken, cfg.Swapper,
			cfg.PriceProvider, cfg.Contract, cfg.ChainId, cfg.prvKey, cfg.providers, ex.l)
	default:
		return fmt.Errorf("only support version '2' and '3'")
	}

	if err != nil {
		return err
	}

	if err := store.UpdateConfigs(ex, cfg); err != nil {
		return err
	}
	ex.dex = dex
	ex.cfg = cfg
	return nil
}

func (c *Config) validate() error {
	if c.Id == 0 {
		return fmt.Errorf("id cannot be empty")
	}
	if c.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if c.Network == "" {
		return fmt.Errorf("network cannot be empty")
	}
	if c.HexKey == "" {
		return fmt.Errorf("hexKey cannot be empty")
	}

	if c.NativeToken == "" {
		return fmt.Errorf("nativeToken cannot be empty")
	}
	if c.TokenStandard == "" {
		return fmt.Errorf("tokenStandard cannot be empty")
	}
	if c.PriceProvider == "" {
		return fmt.Errorf("priceProvider cannot be empty")
	}
	if c.Contract == "" {
		return fmt.Errorf("contract cannot be empty")
	}
	if c.Swapper == "" {
		return fmt.Errorf("swapper cannot be empty")
	}
	if len(c.Providers) == 0 {
		return fmt.Errorf("providers cannot be empty")
	}

	return nil
}
