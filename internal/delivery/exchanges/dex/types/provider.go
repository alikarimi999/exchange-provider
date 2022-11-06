package types

import (
	"context"

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

func (p *Provider) Ping() error {
	_, err := p.BlockNumber(context.Background())
	return err
}
