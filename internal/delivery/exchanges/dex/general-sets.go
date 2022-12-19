package dex

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (u *dex) checkProviders() error {
	var chainId *big.Int
	for i, p := range u.cfg.Providers {
		c, err := ethclient.Dial(p.URL)
		if err != nil {
			return err
		}
		cId, err := c.ChainID(context.Background())
		if err != nil {
			return err
		}
		if i == 0 {
			chainId = cId
		} else {
			if cId != chainId {
				return fmt.Errorf("providers mismatch for chain Id")
			}
		}
		p.Client = c
	}

	u.cfg.ChainId = chainId.Uint64()
	u.cfg.chainId = strconv.Itoa(int(u.cfg.ChainId))
	return nil
}

func (u *dex) setupWallet() error {
	w, err := u.ws.AddWallet(u.cfg.Mnemonic, u.cfg.chainId, u.provider().URL, u.cfg.AccountCount)
	if err != nil {
		return err
	}
	u.wallet = w
	u.cfg.Accounts, _ = w.AllAccounts(u.cfg.AccountCount)
	return nil
}
