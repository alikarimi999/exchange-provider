package multichain

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	Id           string
	Mnemonic     string
	AccountCount uint64
	// Accounts     []accounts.Account
	PL  *ProviderList
	Msg string
}

func EmptyConfig() *Config {
	return &Config{
		PL: &ProviderList{
			list: make(map[chainId][]*types.Provider),
		},
	}
}

type ProviderList struct {
	list map[chainId][]*types.Provider
}

func (pl *ProviderList) Add(cId string, urls []string) error {
	prs := []*types.Provider{}
	for _, url := range urls {
		c, err := ethclient.Dial(url)
		if err != nil {
			return err
		}
		id, err := c.ChainID(context.Background())
		if err != nil {
			return err
		}
		if id.String() != cId {
			return fmt.Errorf("wrong provider")
		}
		prs = append(prs, &types.Provider{Client: c, URL: url})
	}

	pl.list[chainId(cId)] = append(pl.list[chainId(cId)], prs...)
	return nil
}

func (cfg *Config) Validate() error {

	if cfg.AccountCount == 0 {
		cfg.AccountCount = 1
	}

	if cfg.Mnemonic == "" {
		cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}

	return nil
}
