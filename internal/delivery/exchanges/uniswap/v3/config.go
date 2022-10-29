package uniswapv3

import (
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Id            string
	Name          string
	ChianId       uint64
	Network       string
	NativeToken   string
	TokenStandard string

	Providers []*Provider

	Factory       common.Address
	Router        common.Address
	Mnemonic      string
	AccountCount  uint64
	Accounts      []accounts.Account
	ConfirmBlocks uint64
	TokensFile    string
}

func (cfg *Config) Validate(readConfig bool) error {

	if !readConfig {
		if cfg.TokenStandard == "" {
			return errors.New("token-standard cannot be empty")
		}
		if cfg.NativeToken == "" {
			return errors.New("native-token cannot be empty")
		}
		if cfg.Factory == common.BytesToAddress([]byte{0}) {
			return errors.New("factory address cannot be empty")
		}

		if cfg.Router == common.BytesToAddress([]byte{0}) {
			return errors.New("router address cannot be empty")
		}
		if cfg.TokensFile == "" {
			return errors.New("tokens_file cannot be empty")
		}
	}

	if cfg.Name == "" {
		return errors.New("name cannot be empty")
	}

	if cfg.Network == "" {
		return errors.New("network cannot be empty")
	}

	if cfg.ConfirmBlocks == 0 {
		cfg.ConfirmBlocks = 1
	}

	if cfg.Mnemonic == "" {
		cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}

	return nil
}

func (c *Config) wrapNative() string {
	return fmt.Sprintf("W%s", c.NativeToken)
}
