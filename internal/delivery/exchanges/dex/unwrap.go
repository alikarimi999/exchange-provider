package dex

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *dex) unwrap(from, contract common.Address, value *big.Int) (*types.Transaction, error) {
	agent := u.agent("unwrap")

	var err error

	opts, err := u.wallet.NewKeyedTransactorWithChainID(from, big.NewInt(0), int64(u.cfg.ChianId))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			u.wallet.ReleaseNonce(from, opts.Nonce.Uint64())
		} else {
			u.wallet.BurnNonce(from, opts.Nonce.Uint64())

		}
	}()

	c, err := contracts.NewMain(contract, u.provider())
	if err != nil {
		return nil, err
	}

	tx, err := c.Withdraw(opts, value)
	if err != nil {
		u.l.Error(agent, err.Error())
		return nil, err
	}
	return tx, nil
}
