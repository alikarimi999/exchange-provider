package uniswapv3

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

// web3 provider
type Provider struct {
	*ethclient.Client
	URL     string
	counter int
}

func (p *Provider) String() string {
	return p.URL
}

func (p *Provider) ping() error {
	_, err := p.BlockNumber(context.Background())
	return err
}
