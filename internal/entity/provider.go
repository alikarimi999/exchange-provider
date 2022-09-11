package entity

import "github.com/ethereum/go-ethereum/ethclient"

// web3 provider
type Provider struct {
	*ethclient.Client
	Url string
}

func (p *Provider) String() string {
	return p.Url
}
