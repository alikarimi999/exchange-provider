package dex

import (
	"context"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (u *dex) generalSets() error {
	for _, p := range u.cfg.Providers {
		c, err := ethclient.Dial(p.URL)
		if err != nil {
			return err
		}
		cId, err := c.ChainID(context.Background())
		if err != nil {
			return err
		} else {
			u.cfg.ChainId = cId.Uint64()
			u.cfg.chainId = fmt.Sprintf("%d", u.cfg.ChainId)
		}

		p.Client = c
	}

	if err := u.setWallet(); err != nil {
		return err
	}

	return nil
}

func (u *dex) setWallet() error {
	if u.cfg.AccountCount == 0 {
		u.cfg.AccountCount = 1
	}
	if u.cfg.Mnemonic == "" {
		u.cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}
	w, err := eth.NewWallet(u.cfg.Mnemonic, u.provider().Client, u.cfg.AccountCount)
	if err != nil {
		return err
	}
	u.wallet = w
	u.cfg.Accounts, _ = w.AllAccounts()
	return nil
}
