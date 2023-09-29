package entity

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Allowance struct {
	Token   *Token
	Owner   string
	Spender string
	Amount  *big.Int
}

type EVMDex interface {
	Exchange
	Network() string
	Standard() string
	CreateTx(Order) ([]Tx, error)
	CreateSwapBytes(in, out TokenId, tokenOwner, sender, receiver,
		mainContract common.Address, amount, feeAmount float64,
		prvKey *ecdsa.PrivateKey) ([]byte, error)
	ExchangeFeeAmount(in TokenId, p *Pair, exchangeFee float64) (efa float64, price float64, err error)
	Allowance(t *Token, owner string) (*Allowance, error)
	GetToken(tId TokenId) *Token
}
