package allbridge

import (
	"context"
	"crypto/ecdsa"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/networks/evm"
	"exchange-provider/internal/delivery/exchanges/dex/allbridge/types"
	"exchange-provider/internal/entity"
	"fmt"

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
	Enable      bool          `json:"enable"`
	FeeRate     float64       `json:"feeRate"`
	ExchangeFee float64       `json:"exchangeFee"`
	Networks    mapCfgNetwork `json:"networks"`
	Message     string        `json:"message"`
}

func (ex *exchange) UpdateConfigs(cfgi interface{}, store entity.ExchangeStore) error {
	cfg, ok := cfgi.(*Config)
	if !ok {
		return fmt.Errorf("invalid config")
	}
	if err := cfg.validate(); err != nil {
		return nil
	}
	cfg.Enable = ex.cfg.Enable
	tl, err := getTokenInfo(cfg.Networks)
	if err != nil {
		return err
	}

	ns := make(map[string]types.Network)
	for _, n := range cfg.Networks {
		switch n.Type {
		case types.EvmNetwork:
			net, err := evm.NewEvmNetwork(ex.NID(), n.Network, n.AllbridgeContract,
				tl[n.Network].BridgeAddress, n.MainContract, n.client, n.prvKey)
			if err != nil {
				return err
			}
			ns[n.Network] = net

		default:
			return fmt.Errorf("type %s not supported", n.Type)
		}
	}

	c, err := newCache(ex, false)
	if err != nil {
		return err
	}

	ps, err := ex.createPairs(cfg.ExchangeFee, cfg.FeeRate, tl, true)
	if err != nil {
		return err
	}

	if err := ex.store.UpdateConfigs(ex, cfg); err != nil {
		return err
	}

	if len(ps) > 0 {
		for _, p := range ps {
			fmt.Println(p, p.ExchangeFee)
		}
		err := ex.pairs.Add(ex, ps...)
		if err != nil {
			return err
		}
	}

	ex.tl = tl
	ex.ns = ns
	ex.c = c
	ex.cfg = cfg
	return nil
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
