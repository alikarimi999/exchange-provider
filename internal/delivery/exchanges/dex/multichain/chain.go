package multichain

import (
	ts "exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"strconv"
)

type Chain struct {
	id int64

	am *utils.ApproveManager
	w  *eth.HDWallet

	// nativeToken string
	ps []*ts.Provider

	l logger.Logger
}

func (m *Multichain) newChain(cId string) (*Chain, error) {
	id, _ := strconv.Atoi(cId)
	c := &Chain{
		id: int64(id),
		ps: m.cfg.PL.list[chainId(cId)],
		l:  m.l,
	}

	w, err := eth.NewWallet(m.cfg.Mnemonic, c.provider().Client, m.cfg.AccountCount)
	if err != nil {
		return nil, err
	}
	c.w = w
	c.am = utils.NewApproveManager(int64(id), m.tt, w, m.l, c.ps)
	return c, nil
}
