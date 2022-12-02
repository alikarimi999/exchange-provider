package multichain

import (
	"context"
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	id int64

	am *utils.ApproveManager
	w  *eth.HDWallet

	// nativeToken string
	Providers []*ts.Provider

	l logger.Logger
}

func (m *Multichain) newChain(id ChainId, urls ...string) (*Chain, []URL, error) {
	cId, _ := strconv.Atoi(string(id))
	c := &Chain{
		id: int64(cId),
		l:  m.l,
	}

	us := []URL{}
	for _, url := range urls {
		if err := c.addProvider(url); err != nil {
			us = append(us, URL{Url: url, Msg: err.Error()})
			continue
		}
		us = append(us, URL{Url: url, Msg: addMsg})
	}

	if len(c.Providers) > 0 {
		w, err := eth.NewWallet(m.cfg.Mnemonic, c.provider().Client, m.cfg.AccountCount)
		if err != nil {
			return nil, us, err
		}
		c.w = w
		c.am = utils.NewApproveManager(int64(cId), m.tt, w, m.l, c.Providers)
		return c, us, nil
	}
	return nil, us, fmt.Errorf("all urls was invalid")
}

func (c *Chain) addProvider(url string) error {
	for _, p := range c.Providers {
		if p.URL == url {
			return nil
		}
	}

	p := &ts.Provider{URL: url}
	cl, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	id, err := cl.ChainID(context.Background())
	if err != nil {
		return err
	}
	if id.Int64() != c.id {
		return fmt.Errorf("invalid url")
	}

	p.Client = cl
	c.Providers = append(c.Providers, p)
	return nil
}
