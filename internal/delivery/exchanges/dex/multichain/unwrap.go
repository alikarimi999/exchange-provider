package multichain

import (
	"exchange-provider/internal/delivery/exchanges/dex/uniswap/v3/contracts"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (m *Chain) unwrap(from, contract common.Address, value *big.Int) (*types.Transaction, error) {
	agent := "unwrap"

	var err error

	opts, err := m.w.NewKeyedTransactorWithChainID(from, big.NewInt(0), m.id)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			m.w.ReleaseNonce(from, opts.Nonce.Uint64())
		} else {
			m.w.BurnNonce(from, opts.Nonce.Uint64())

		}
	}()

	c, err := contracts.NewMain(contract, m.provider())
	if err != nil {
		return nil, err
	}

	tx, err := c.Withdraw(opts, value)
	if err != nil {
		m.l.Error(agent, err.Error())
		return nil, err
	}
	return tx, nil
}
