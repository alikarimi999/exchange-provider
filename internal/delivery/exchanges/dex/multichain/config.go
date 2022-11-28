package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/wallet/eth"
)

type Config struct {
	Id           string
	Mnemonic     string
	AccountCount uint64
	// Accounts     []accounts.Account
	PL  *ProviderList
	Msg string
}

type ProviderList struct {
	list map[chainId][]*types.Provider
}

func (pl *ProviderList) Add(cId string, ps []*types.Provider) {
	pl.list[chainId(cId)] = append(pl.list[chainId(cId)], ps...)
}

func (cfg *Config) Validate(readConfig bool) error {

	if cfg.Id == "" {
		return errors.New("name cannot be empty")
	}

	if cfg.AccountCount == 0 {
		cfg.AccountCount = 1
	}

	if cfg.Mnemonic == "" {
		cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}

	return nil
}
