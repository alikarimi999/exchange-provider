package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Id                 uint   `json:"id"`
	Name               string `json:"name"`
	NativeToken        string `json:"nativeToken"`
	WrappedNativeToken string `json:"wrappedNativeToken"`
	ChainId            int64  `json:"chainId"`
	TokenStandard      string `json:"tokenStandard"`
	Network            string `json:"network"`
	Contract           string `json:"contract"`
	contractAddress    common.Address
	Swapper            string `json:"swapper"`
	swapperAddress     common.Address
	HexKey             string `json:"hexKey"`
	privateKey         *ecdsa.PrivateKey
	Providers          []string `json:"providers"`
	providers          []*types.EthProvider
	Message            string `json:"message"`
}

func (c *Config) Validate(readConfig bool) error {
	if c.Id == 0 {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("id must not be empty"))
	}
	if c.Name == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("name must not be empty"))
	}
	if c.Network == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("network must not be empty"))
	}
	if c.HexKey == "" {
		return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("hexKey must not be empty"))
	}

	if !readConfig {

		if c.NativeToken == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("nativeToken must not be empty"))
		}
		if c.TokenStandard == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tokenStandard must not be empty"))
		}
		if c.Contract == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("contract must not be empty"))
		}
		if c.Swapper == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("swapper must not be empty"))
		}
		if len(c.Providers) == 0 {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("providers must not be empty"))
		}
	}
	return nil
}
