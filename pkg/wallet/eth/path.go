package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
)

const (
	ethPath = "m/44'/60'/0'/0"
)

func (hd *HDWallet) getPath(index uint64) accounts.DerivationPath {
	path, _ := accounts.ParseDerivationPath(hd.path(index))
	return path
}

func (hd *HDWallet) path(index uint64) string {
	hd.mu.Lock()
	defer hd.mu.Unlock()
	return fmt.Sprintf("%s/%d", hd.p, index)
}
