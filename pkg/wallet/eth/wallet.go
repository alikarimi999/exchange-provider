package eth

import (
	"crypto/ecdsa"
	"order_service/pkg/errors"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWallet struct {
	w        *hdwallet.Wallet
	mu       *sync.Mutex
	mnemonic string
	p        string
}

func CreateWallet() *HDWallet {
	hd := &HDWallet{p: ethPath}

	mn, _ := hdwallet.NewMnemonic(256)
	hd.mnemonic = mn

	w, _ := hdwallet.NewFromMnemonic(mn)
	hd.w = w

	hd.AddAccount(0)

	return hd
}

func WalletFromMnemonic(mn string) (*HDWallet, error) {
	hd := &HDWallet{p: ethPath}

	w, err := hdwallet.NewFromMnemonic(mn)
	if err != nil {
		return nil, err
	}

	hd.mnemonic = mn
	hd.w = w

	hd.AddAccount(0)
	return hd, nil
}

func (hd *HDWallet) Mnemonic() string {
	hd.mu.Lock()
	defer hd.mu.Unlock()
	return hd.mnemonic
}

func (hd *HDWallet) AddAccount(index uint64) (accounts.Account, error) {
	return hd.w.Derive(hd.getPath(index), true)
}

func (hd *HDWallet) Account(index uint64) (accounts.Account, error) {
	return hd.getAccount(index)
}

func (hd *HDWallet) AccountByAddress(address common.Address) (accounts.Account, error) {
	return hd.getAccountByAddress(address)
}

func (hd *HDWallet) getAccount(index uint64) (accounts.Account, error) {
	as := hd.w.Accounts()
	p := hd.path(index)
	for _, a := range as {
		if a.URL.Path == p {
			return a, nil
		}
	}
	return accounts.Account{}, errors.Wrap(errors.ErrNotFound)
}

func (hd *HDWallet) getAccountByAddress(address common.Address) (accounts.Account, error) {
	as := hd.w.Accounts()
	for _, a := range as {
		if a.Address == address {
			return a, nil
		}
	}
	return accounts.Account{}, errors.Wrap(errors.ErrNotFound)
}

func (hd *HDWallet) Address(index uint64) (common.Address, error) {
	acc, err := hd.getAccount(index)
	if err != nil {
		return common.Address{}, err
	}
	return acc.Address, nil
}

func (hd *HDWallet) PrivateKey(address common.Address) (*ecdsa.PrivateKey, error) {
	acc, err := hd.getAccountByAddress(address)
	if err != nil {
		return nil, err
	}
	return hd.w.PrivateKey(acc)
}

func (hd *HDWallet) PublicKey(address common.Address) (*ecdsa.PublicKey, error) {
	acc, err := hd.getAccountByAddress(address)
	if err != nil {
		return nil, err
	}
	return hd.w.PublicKey(acc)
}
