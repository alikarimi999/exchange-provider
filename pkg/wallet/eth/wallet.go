package eth

import (
	"context"
	"crypto/ecdsa"
	"math/rand"
	"order_service/pkg/errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWallet struct {
	w      *hdwallet.Wallet
	client *ethclient.Client
	*tracker
	mu       *sync.Mutex
	mnemonic string
	p        string
}

func NewWallet(mnemonic string, client *ethclient.Client) (*HDWallet, error) {
	mn := mnemonic
	var err error
	if mnemonic == "" {
		mn, err = hdwallet.NewMnemonic(256)
		if err != nil {
			return nil, err
		}
	}
	return walletFromMnemonic(mn, client)
}

func walletFromMnemonic(mn string, c *ethclient.Client) (*HDWallet, error) {
	hd := &HDWallet{
		p:       ethPath,
		client:  c,
		tracker: newTracker(),
		mu:      &sync.Mutex{},
	}

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
	acc, err := hd.w.Derive(hd.getPath(index), true)
	if err != nil {
		return accounts.Account{}, err
	}
	n, err := hd.client.PendingNonceAt(context.Background(), acc.Address)
	if err != nil {
		return accounts.Account{}, err
	}

	hd.tracker.addAccount(acc.Address, n)
	return acc, nil
}

func (hd *HDWallet) Nonce(address common.Address) (uint64, error) {
	return hd.tracker.nonce(address)
}

func (hd *HDWallet) BurnNonce(address common.Address, nonce uint64) {
	hd.tracker.burnNonce(address, nonce)
}

func (hd *HDWallet) ReleaseNonce(address common.Address, nonce uint64) {
	hd.tracker.releaseNonce(address, nonce)
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
	return hd.getAddress(index)
}
func (hd *HDWallet) getAddress(index uint64) (common.Address, error) {
	acc, err := hd.getAccount(index)
	if err != nil {
		return common.Address{}, err
	}
	return acc.Address, nil
}

func (hd *HDWallet) RandAddress() (common.Address, error) {
	count := len(hd.w.Accounts())
	if count == 0 {
		return common.Address{}, errors.Wrap(errors.ErrNotFound)
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return hd.getAddress(uint64(r.Intn(count)))
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
