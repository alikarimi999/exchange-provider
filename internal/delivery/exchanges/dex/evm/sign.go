package evm

import "github.com/ethereum/go-ethereum/crypto"

func (d *EvmDex) sign(data interface{}) ([]byte, error) {
	packed, err := args.Pack(data)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(crypto.Keccak256Hash(packed).Bytes(), d.privateKey)
}
