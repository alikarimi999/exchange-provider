package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/delivery/exchanges/dex/utils"
	"exchange-provider/pkg/logger"
	"exchange-provider/pkg/wallet/eth"
)

type Multichain struct {
	nid string

	c      chains
	wallet *eth.HDWallet

	tt *utils.TxTracker
	am *utils.ApproveManager
	ps []*types.Provider

	l logger.Logger

	apiUrl string
}

func (m *Multichain) Pair(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt}, nil
}

func (m *Multichain) PairWithPrice(bt, qt types.Token) (*types.Pair, error) {
	return &types.Pair{T1: bt, T2: qt, Price: "1"}, nil

}
