package eth

import (
	"order_service/pkg/errors"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type tracker struct {
	mux      *sync.Mutex
	accounts map[string]*account
}

func newTracker() *tracker {
	return &tracker{
		mux:      &sync.Mutex{},
		accounts: make(map[string]*account),
	}
}

func (t *tracker) addAccount(address common.Address, latest_nonce uint64) {
	t.mux.Lock()
	defer t.mux.Unlock()

	for s := range t.accounts {
		if s == address.String() {
			return
		}
	}
	t.accounts[address.String()] = newAccount(address, latest_nonce)
}

func (t *tracker) nonce(address common.Address) (uint64, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if a, ok := t.accounts[address.String()]; ok {
		return a.nonce(), nil
	}
	return 0, errors.Wrap(errors.ErrNotFound)
}

func (t *tracker) burnNonce(address common.Address, nonce uint64) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if a, ok := t.accounts[address.String()]; ok {
		a.burn(nonce)
	}
}

func (t *tracker) releaseNonce(address common.Address, nonce uint64) {
	t.mux.Lock()
	defer t.mux.Unlock()

	a, ok := t.accounts[address.String()]
	if ok {
		a.release(nonce)
	}

}
