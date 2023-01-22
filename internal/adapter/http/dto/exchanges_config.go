package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ChianId       uint64 `json:"chianId,omitempty"`
	Network       string `json:"network,omitempty"`
	NativeToken   string `json:"nativeToken,omitempty"`
	TokenStandard string `json:"tokenStandard,omitempty"`

	Factory string `json:"factory,omitempty"`
	Router  string `json:"router,omitempty"`

	Providers []string `json:"providers,omitempty"`

	Mnemonic     string             `json:"mnemonic,omitempty"`
	AccountCount uint64             `json:"account_count,omitempty"`
	Accounts     []accounts.Account `json:"accounts,omitempty"`

	BlockTime     string `json:"blockTime,omitempty"`
	ConfirmBlocks uint64 `json:"confirmBlocks,omitempty"`
	TokensFile    string `json:"tokensFile,omitempty"`
	Msg           string `json:"message,omitempty"`
}

func (cfg *Config) Map() (*dex.Config, error) {
	c := &dex.Config{
		Name:          cfg.Name,
		ChainId:       cfg.ChianId,
		Network:       cfg.Network,
		NativeToken:   cfg.NativeToken,
		TokenStandard: cfg.TokenStandard,

		Factory: common.HexToAddress(cfg.Factory),
		Router:  common.HexToAddress(cfg.Router),

		Mnemonic:     cfg.Mnemonic,
		AccountCount: cfg.AccountCount,

		ConfirmBlocks: cfg.ConfirmBlocks,
		TokensFile:    cfg.TokensFile,
	}

	bt, err := toTime(cfg.BlockTime)
	if err != nil {
		return nil, err
	}

	c.BlockTime = bt

	for _, url := range cfg.Providers {
		c.Providers = append(c.Providers, &ts.EthProvider{URL: url})
	}

	return c, nil
}
