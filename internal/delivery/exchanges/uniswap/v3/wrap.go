package uniswapv3

import (
	"fmt"
	"math/big"
	"order_service/internal/delivery/exchanges/uniswap/v3/contracts"
	"order_service/pkg/utils/numbers"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (u *UniSwapV3) wrap(from, contract common.Address, amount string) (*types.Transaction, error) {
	agent := u.agent("wrap")

	value, err := numbers.FloatStringToBigInt(amount, ethDecimals)
	if err != nil {
		return nil, err
	}

	opts, err := u.newKeyedTransactorWithChainID(from, value)
	if err != nil {
		return nil, err
	}

	c, err := contracts.NewMain(contract, u.provider)
	if err != nil {
		return nil, err
	}

	tx, err := c.Deposit(opts)
	if err != nil {
		u.wallet.ReleaseNonce(from, opts.Nonce.Uint64())
		u.l.Debug(agent, fmt.Sprintf("wrap `%s` ETH failed", amount))
		return nil, err
	}
	u.wallet.BurnNonce(from, tx.Nonce())

	u.l.Debug(agent, fmt.Sprintf("wrap `%s` ETH was successfull", amount))
	return tx, nil
}

func (u *UniSwapV3) unwrap(from, contract common.Address, amount string) (*types.Transaction, error) {
	agent := u.agent("unwrap")

	value, err := numbers.FloatStringToBigInt(amount, ethDecimals)
	if err != nil {
		return nil, err
	}

	opts, err := u.newKeyedTransactorWithChainID(from, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	c, err := contracts.NewMain(contract, u.provider)
	if err != nil {
		return nil, err
	}

	tx, err := c.Withdraw(opts, value)
	if err != nil {
		u.wallet.ReleaseNonce(from, opts.Nonce.Uint64())
		u.l.Debug(agent, fmt.Sprintf("wrap `%s` ETH failed", amount))
		return nil, err
	}
	u.wallet.BurnNonce(from, tx.Nonce())

	u.l.Debug(agent, fmt.Sprintf("unwrap `%s` WETH was successfull", amount))
	return tx, nil
}
