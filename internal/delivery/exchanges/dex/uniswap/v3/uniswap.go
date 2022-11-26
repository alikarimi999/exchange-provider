package uniswapv3

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type UniswapV3 struct {
	id string
	ps []*types.Provider

	factory common.Address
	router  common.Address
	nt      string

	tt       *utils.TxTracker
	wallet   *eth.HDWallet
	chaindId int64

	l logger.Logger
}

func NewUniSwapV3(id, nt string, ps []*types.Provider, f, r common.Address, w *eth.HDWallet, tt *utils.TxTracker, l logger.Logger) (*UniswapV3, error) {
	u := &UniswapV3{
		id: id,
		ps: ps,

		factory: f,
		router:  r,
		nt:      nt,

		tt:     tt,
		wallet: w,
		l:      l,
	}

	c, err := u.provider().Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	u.chaindId = c.Int64()
	return u, nil
}

func (u *UniswapV3) agent(fn string) string {
	return fmt.Sprintf("%s-%s", u.id, fn)
}
