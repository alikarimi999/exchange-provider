package dto

import (
	uniswapv3 "exchange-provider/internal/delivery/exchanges/uniswap/v3"

	"github.com/ethereum/go-ethereum/accounts"
)

type Config struct {
	Id            string `json:"id,omitempty"`
	ChianId       uint64 `json:"chian_id,omitempty"`
	Network       string `json:"network,omitempty"`
	NativeToken   string `json:"native_token,omitempty"`
	TokenStandard string `json:"token_standard,omitempty"`

	Providers []string `json:"providers,omitempty"`

	Mnemonic      string             `json:"mnemonic,omitempty"`
	AccountCount  uint64             `json:"account_count,omitempty"`
	Accounts      []accounts.Account `json:"accounts,omitempty"`
	ConfirmBlocks uint64             `json:"confirm_blocks,omitempty"`
	TokensFile    string             `json:"tokens_file,omitempty"`
	Msg           string             `json:"msg,omitempty"`
}

func (cfg *Config) Map() *uniswapv3.Config {
	c := &uniswapv3.Config{
		ChianId:       cfg.ChianId,
		Network:       cfg.Network,
		TokenStandard: cfg.TokenStandard,
		NativeToken:   cfg.NativeToken,

		Mnemonic:      cfg.Mnemonic,
		AccountCount:  cfg.AccountCount,
		ConfirmBlocks: cfg.ConfirmBlocks,
		TokensFile:    cfg.TokensFile,
	}

	for _, url := range cfg.Providers {
		c.Providers = append(c.Providers, &uniswapv3.Provider{URL: url})
	}

	return c
}
