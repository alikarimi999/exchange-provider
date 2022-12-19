package dex

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/wallet/eth"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Name        string
	ChainId     uint64
	chainId     string
	Network     string
	NativeToken string
	// TokenStandard string

	Providers []*ts.EthProvider

	Factory       common.Address
	Router        common.Address
	Mnemonic      string
	AccountCount  uint64
	Accounts      []accounts.Account
	ConfirmBlocks uint64
	TokensFile    string

	BlockTime time.Duration
}

func (cfg *Config) Validate(readConfig bool) error {

	if !readConfig {
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

		if cfg.BlockTime == time.Duration(0) {
			return errors.New("block_time cannot be empty")
		}

		cfg.BlockTime += time.Duration(5 * time.Second)
	}

	if cfg.AccountCount == 0 {
		cfg.AccountCount = 1
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
