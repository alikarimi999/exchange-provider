package eth

import (
	"exchange-provider/pkg/bind"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (u *HDWallet) NewKeyedTransactorWithChainID(address common.Address, value *big.Int, chainId int64) (*bind.TransactOpts, error) {
	key, err := u.PrivateKey(address)
	if err != nil {
		return nil, err
	}

	nonce, err := u.Nonce(address)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = value
	opts.GasLimit = 0

	return opts, nil
}
