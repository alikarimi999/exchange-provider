package uniswapv3

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type UniswapV3 struct {
	Id       string
	Ps       []*types.Provider
	Factory  common.Address
	Router   common.Address
	Wallet   *eth.HDWallet
	ChaindId int64

	L logger.Logger
}

func NewUniSwapV3(id string, ps []*types.Provider, f, r common.Address, w *eth.HDWallet, l logger.Logger) (*UniswapV3, error) {
	u := &UniswapV3{
		Id:      id,
		Ps:      ps,
		Factory: f,
		Router:  r,
		Wallet:  w,
		L:       l,
	}

	c, err := u.provider().Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	u.ChaindId = c.Int64()
	return u, nil
}

func (u *UniswapV3) agent(fn string) string {
	return fmt.Sprintf("%s-%s", u.Id, fn)
}
