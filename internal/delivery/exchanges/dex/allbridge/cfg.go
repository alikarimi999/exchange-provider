package allbridge

import (
	"context"
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type cfgNetwork struct {
	Standard          string `json:"standard"`
	Network           string `json:"network"`
	MainContract      string `json:"mainContract"`
	AllbridgeContract string `json:"allbridgeContract"`
	Provider          string `json:"provider"`
	HexKey            string `json:"hexKey"`
	chainId           int64
	prvKey            *ecdsa.PrivateKey
	client            *ethclient.Client
	Type              types.NetworkType `json:"type"`
}

type mapCfgNetwork map[string]*cfgNetwork

func (m mapCfgNetwork) network(network string) *cfgNetwork {
	for _, mn := range m {
		if mn.Network == network {
			return mn
		}
	}
	return nil
}

type Config struct {
	Id          uint          `json:"id"`
	Url         string        `json:"url"`
	Name        string        `json:"name"`
	Enable      bool          `json:"enable"`
	FeeRate     float64       `json:"feeRate"`
	ExchangeFee float64       `json:"exchangeFee"`
	Networks    mapCfgNetwork `json:"networks"`
	Message     string        `json:"message"`
}

func (c *Config) validate() error {
	for _, n := range c.Networks {
		c, err := ethclient.Dial(n.Provider)
		if err != nil {
			return err
		}
		n.client = c
		chainId, err := c.ChainID(context.Background())
		if err != nil {
			return err
		}
		n.chainId = chainId.Int64()
		k, err := crypto.HexToECDSA(n.HexKey)
		if err != nil {
			return err
		}
		n.prvKey = k
	}
	return nil
}
