package uniswapv3

import (
	"exchange-provider/internal/delivery/exchanges/uniswap/v3/contracts"
	"exchange-provider/pkg/utils/numbers"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniSwapV3) wrap(from, contract common.Address, amount string) (*types.Transaction, error) {
	agent := u.agent("wrap")

	var err error
	value, err := numbers.FloatStringToBigInt(amount, ethDecimals)
	if err != nil {
		return nil, err
	}

	opts, err := u.newKeyedTransactorWithChainID(from, value)
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

	c, err := contracts.NewMain(contract, u.provider)
	if err != nil {
		return nil, err
	}

	tx, err := c.Deposit(opts)
	if err != nil {
		u.l.Error(agent, err.Error())
		return nil, err
	}

	u.l.Debug(agent, fmt.Sprintf("wrap `%s` ETH was successfull", amount))
	return tx, nil
}

func (u *UniSwapV3) unwrap(from, contract common.Address, value *big.Int) (*types.Transaction, error) {
	agent := u.agent("unwrap")

	var err error

	opts, err := u.newKeyedTransactorWithChainID(from, big.NewInt(0))
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

	c, err := contracts.NewMain(contract, u.provider)
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
