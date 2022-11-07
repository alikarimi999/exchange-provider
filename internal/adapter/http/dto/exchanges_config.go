package dto

import (
	"exchange-provider/internal/delivery/exchanges/dex"
	"exchange-provider/internal/delivery/exchanges/dex/types"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ChianId       uint64 `json:"chian_id,omitempty"`
	Network       string `json:"network,omitempty"`
	NativeToken   string `json:"native_token,omitempty"`
	TokenStandard string `json:"token_standard,omitempty"`

	Factory string `json:"factory,omitempty"`
	Router  string `json:"router,omitempty"`

	Providers []string `json:"providers,omitempty"`

	Mnemonic     string             `json:"mnemonic,omitempty"`
	AccountCount uint64             `json:"account_count,omitempty"`
	Accounts     []accounts.Account `json:"accounts,omitempty"`

	BlockTime     string `json:"block_time,omitempty"`
	ConfirmBlocks uint64 `json:"confirm_blocks,omitempty"`
	TokensFile    string `json:"tokens_file,omitempty"`
	Msg           string `json:"msg,omitempty"`
}

func (cfg *Config) Map() (*dex.Config, error) {
	c := &dex.Config{
		Name:          cfg.Name,
		ChianId:       cfg.ChianId,
		Network:       cfg.Network,
		TokenStandard: cfg.TokenStandard,
		NativeToken:   cfg.NativeToken,

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
		c.Providers = append(c.Providers, &types.Provider{URL: url})
	}

	return c, nil
}
