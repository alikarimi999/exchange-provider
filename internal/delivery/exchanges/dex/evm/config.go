package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
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

func (c *Config) Validate() error {
	if c.Id == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id cannot be empty"))
	}
	if c.Name == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("name cannot be empty"))
	}
	if c.Network == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("network cannot be empty"))
	}
	if c.HexKey == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("hexKey cannot be empty"))
	}

	if c.NativeToken == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("nativeToken cannot be empty"))
	}
	if c.TokenStandard == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tokenStandard cannot be empty"))
	}
	if c.PriceProvider == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("priceProvider cannot be empty"))
	}
	if c.Contract == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("contract cannot be empty"))
	}
	if c.Swapper == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("swapper cannot be empty"))
	}
	if len(c.Providers) == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("providers cannot be empty"))
	}

	return nil
}
