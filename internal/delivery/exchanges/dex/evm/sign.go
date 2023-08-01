package evm

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func (d *exchange) sign(data interface{}, prvKey *ecdsa.PrivateKey) ([]byte, error) {
	packed, err := args.Pack(data)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(crypto.Keccak256Hash(packed).Bytes(), prvKey)
}
