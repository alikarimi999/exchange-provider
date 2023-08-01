package evm

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/evm/types"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (u *exchange) checkProviders() error {
	var chainId *big.Int
	ps := []*types.EthProvider{}
	for i, p := range u.cfg.Providers {
		c, err := ethclient.Dial(p)
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
			if cId.Uint64() != chainId.Uint64() {
				return fmt.Errorf("providers mismatch for chain Id")
			}
		}
		ps = append(ps, &types.EthProvider{URL: p, Client: c})
	}

	u.cfg.ChainId = chainId.Int64()
	u.cfg.providers = ps
	return nil
}

func (d *exchange) provider() *types.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(d.cfg.providers) > 0 {
		p := d.cfg.providers[r.Intn(len(d.cfg.providers))]
		return p
	}
	return &types.EthProvider{}
}
