package multichain

import (
	"context"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	Name         string
	Mnemonic     string
	AccountCount uint64
	// Accounts     []accounts.Account
	Msg string
}

type ProviderList struct {
	list map[ChainId][]*ts.EthProvider
}

func (pl *ProviderList) Add(cId string, urls []string) error {
	prs := []*ts.EthProvider{}
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
		prs = append(prs, &ts.EthProvider{Client: c, URL: url})
	}

	pl.list[ChainId(cId)] = append(pl.list[ChainId(cId)], prs...)
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
