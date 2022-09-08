package uniswapv3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (u *UniSwapV3) newKeyedTransactorWithChainID(address common.Address, value *big.Int) (*bind.TransactOpts, error) {
	key, err := u.wallet.PrivateKey(address)
	if err != nil {
		return nil, err
	}

	nonce, err := u.wallet.Nonce(address)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(key, u.chainId)
	if err != nil {
		return nil, err
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = value
	opts.GasLimit = 0

	return opts, nil
}
