package uniswapv3

import (
	"context"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// web3 provider
type Provider struct {
	*ethclient.Client
	URL string
}

func (p *Provider) String() string {
	return p.URL
}

func (p *Provider) ping() error {
	_, err := p.BlockNumber(context.Background())
	return err
}

func (ex *dex) provider() *Provider {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(ex.cfg.Providers) > 0 {
		return ex.cfg.Providers[r.Intn(len(ex.cfg.Providers))]
	}
	return &Provider{}
}
