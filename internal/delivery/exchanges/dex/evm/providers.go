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

func checkProviders(providers []string) ([]*types.EthProvider, *big.Int, error) {
	var chainId *big.Int
	ps := []*types.EthProvider{}
	for i, p := range providers {
		c, err := ethclient.Dial(p)
		if err != nil {
			return nil, nil, err
		}
		cId, err := c.ChainID(context.Background())
		if err != nil {
			return nil, nil, err
		}

		if i == 0 {
			chainId = cId
		} else {
			if cId.Uint64() != chainId.Uint64() {
				return nil, nil, fmt.Errorf("providers mismatch for chain Id")
			}
		}
		ps = append(ps, &types.EthProvider{URL: p, Client: c})
	}

	if len(ps) > 0 {
		return ps, chainId, nil
	}
	return nil, nil, fmt.Errorf("no providers")
}

func (d *exchange) provider() *types.EthProvider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(d.cfg.providers) > 0 {
		p := d.cfg.providers[r.Intn(len(d.cfg.providers))]
		return p
	}
	return &types.EthProvider{}
}
