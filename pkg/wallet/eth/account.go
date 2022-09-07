package eth

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type account struct {
	address common.Address

	next     *atomic.Uint64
	locked   *queue // locked nonces
	released *queue // released nonces
}

func newAccount(address common.Address, n uint64) *account {
	a := &account{
		address:  address,
		next:     new(atomic.Uint64),
		locked:   newQ(),
		released: newQ(),
	}
	a.next.Store(n)
	return a
}

func (a *account) nonce() uint64 {
	if a.released.size() > 0 {
		n := a.released.pop()
		a.locked.push(n)
		return uint64(n)
	}

	n := a.next.Load()
	a.locked.push(int(n))

	a.next.Add(1)

	return uint64(n)
}

func (a *account) burn(n uint64) {
	a.locked.remove(int(n))
}

func (a *account) release(n uint64) {
	a.locked.remove(int(n))
	a.released.push(int(n))
}
