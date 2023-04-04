package allbridge

import "sync"

type token struct {
	Symbol   string
	Network  string
	Address  string
	Decimals uint
}

type pair struct {
	t1 *token
	t2 *token
}

type pairs struct {
	mux  *sync.RWMutex
	list map[string][]*pair
}
