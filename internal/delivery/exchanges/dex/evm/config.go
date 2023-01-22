package evm

import (
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
)

type Config struct {
	Id                 string
	Name               string
	Network            string
	NativeToken        string
	WrappedNativeToken string
	TokensFile         string
	ChainId            int64
	TokenStandard      string
	Contract           string
	Swapper            string
	HexKey             string
	privateKey         *ecdsa.PrivateKey
	Providers          []string
	providers          []*types.EthProvider
	Message            string
}

func (c *Config) Validate(readConfig bool) error {
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
		if c.TokensFile == "" {
			return errors.Wrap(errors.ErrBadRequest, errors.NewMesssage("tokenFile must not be empty"))
		}
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
