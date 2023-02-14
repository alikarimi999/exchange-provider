package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/delivery/exchanges/dex/types"
	"exchange-provider/internal/entity"
	"exchange-provider/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	em "github.com/ethereum/go-ethereum/common/math"
	ts "github.com/ethereum/go-ethereum/core/types"
)

func (d *EvmDex) approveTx(r *entity.Route, owner common.Address) (*ts.Transaction, error) {
	T1, ok := d.get(r.In.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}
	T2, ok := d.get(r.Out.TokenId)
	if !ok {
		return nil, errors.Wrap(errors.ErrNotFound)
	}

	var in types.Token
	if T1.Symbol == r.In.TokenId {
		in = T1
	} else {
		in = T2
	}

	c, err := contracts.NewIERC20(in.Address, d.provider())
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(d.privateKey, big.NewInt(d.ChainId))
	if err != nil {
		return nil, err
	}
	opts.NoSend = true

	return c.Approve(opts, d.contractAddress, em.MaxBig256)
}

func (d *EvmDex) needApproval(r *entity.Route, owner common.Address, minAmount float64) (bool, error) {
	T1, ok := d.get(r.In.TokenId)
	if !ok {
		return false, errors.Wrap(errors.ErrNotFound)
	}
	T2, ok := d.get(r.Out.TokenId)
	if !ok {
		return false, errors.Wrap(errors.ErrNotFound)
	}

	var in types.Token
	if T1.Symbol == r.In.TokenId {
		in = T1
	} else {
		in = T2
	}
	if in.IsNative() {
		return false, nil
	}

	c, err := contracts.NewIERC20(in.Address, d.provider())
	if err != nil {
		return false, err
	}

	amount, err := c.Allowance(nil, owner, d.contractAddress)
	if err != nil {
		return false, err
	}

	mAmount, _ := big.NewFloat(0).Mul(big.NewFloat(minAmount),
		big.NewFloat(0).SetInt(em.BigPow(10, int64(in.Decimals)))).Int(nil)
	return amount.Cmp(mAmount) == -1, nil
}
