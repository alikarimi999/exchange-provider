package walletstore

import (
	"context"
	"exchange-provider/internal/app"
	"exchange-provider/pkg/wallet/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type wallet struct {
	m   string
	cId string
	w   *eth.HDWallet
}

type WalletStore struct {
	ws []*wallet
}

func NewWalletStore() app.WalletStore {
	return &WalletStore{ws: make([]*wallet, 0)}
}

func (s *WalletStore) AddWallet(m, cId, p string, accountCount uint64) (*eth.HDWallet, error) {

	if w, err := s.getWallet(m, cId); err == nil {
		if w.Len() < int(accountCount) {
			for i := w.Len(); i < int(accountCount); i++ {
				if _, err := w.AddAccount(uint64(i)); err != nil {
					for j := w.Len(); j < i; j++ {
						w.RemoveAccount(uint64(j))
					}
					return nil, err
				}
			}
		}
		return w, nil
	}

	c, err := ethclient.Dial(p)
	if err != nil {
		return nil, err
	}
	id, err := c.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	if id.String() != cId {
		return nil, fmt.Errorf("chainId mismatch")
	}
	w, err := eth.NewWallet(m, c, accountCount)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (s *WalletStore) getWallet(m, cId string) (*eth.HDWallet, error) {
	for _, w := range s.ws {
		if w.m == m && w.cId == cId {
			return w.w, nil
		}
	}
	return nil, fmt.Errorf("wallet not found")
}
