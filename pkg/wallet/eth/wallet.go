package eth

import (
	"context"
	"crypto/ecdsa"
	"exchange-provider/pkg/errors"
	"math/rand"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWallet struct {
	w        *hdwallet.Wallet
	client   *ethclient.Client
	tracker  *tracker
	mu       *sync.Mutex
	mnemonic string
	p        string
}

func NewWallet(mnemonic string, client *ethclient.Client, count uint64) (*HDWallet, error) {
	mn := mnemonic
	var err error
	if mnemonic == "" {
		mn, err = hdwallet.NewMnemonic(128)
		if err != nil {
			return nil, err
		}
	}
	return walletFromMnemonic(mn, client, count)
}

func NewMnemonic(bits int) (string, error) {
	return hdwallet.NewMnemonic(bits)
}

func walletFromMnemonic(mn string, c *ethclient.Client, count uint64) (*HDWallet, error) {
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

	for i := 0; i < int(count); i++ {
		if _, err := hd.AddAccount(uint64(i)); err != nil {
			for j := 0; j < i; j++ {
				hd.RemoveAccount(uint64(j))
			}
			return nil, err
		}
	}
	return hd, nil
}

func (hd *HDWallet) Mnemonic() string {
	hd.mu.Lock()
	defer hd.mu.Unlock()
	return hd.mnemonic
}

func (hd *HDWallet) String() string {
	return hd.Mnemonic()
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

func (hd *HDWallet) RemoveAccount(index uint64) {
	a, err := hd.getAccount(index)
	if err != nil {
		return
	}
	hd.w.Unpin(a)
	hd.tracker.removeAccount(a.Address)
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

func (hd *HDWallet) AllAddresses(count uint64) ([]common.Address, error) {
	acc := hd.w.Accounts()[:count]
	addresses := []common.Address{}
	for _, a := range acc {
		addresses = append(addresses, a.Address)
	}

	return addresses, nil
}

func (hd *HDWallet) AllAccounts(count uint64) ([]accounts.Account, error) {
	return hd.w.Accounts()[:count], nil
}

func (hd *HDWallet) Len() int {
	return len(hd.w.Accounts())
}

func (hd *HDWallet) getAddress(index uint64) (common.Address, error) {
	acc, err := hd.getAccount(index)
	if err != nil {
		return common.Address{}, err
	}
	return acc.Address, nil
}

func (hd *HDWallet) RandAddress(count uint64) (common.Address, error) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return hd.getAddress(uint64(r.Intn(int(count))))
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
