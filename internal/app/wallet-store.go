package app

import (
	"exchange-provider/pkg/wallet/eth"
)

type WalletStore interface {
	AddWallet(mnemonic, chainId, provider string, accountCount uint64) (*eth.HDWallet, error)
}
