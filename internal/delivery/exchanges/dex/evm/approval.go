package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	em "github.com/ethereum/go-ethereum/common/math"
	ts "github.com/ethereum/go-ethereum/core/types"
)

func (d *EvmDex) approveTx(r *entity.Route, owner common.Address) (*ts.Transaction, error) {
	t1, err := d.ts.get(r.In.String())
	if err != nil {
		return nil, err
	}
	t2, err := d.ts.get(r.Out.String())
	if err != nil {
		return nil, err
	}

	var in *entity.Token
	if t1.Equal(r.In) {
		in = t1
	} else {
		in = t2
	}

	c, err := contracts.NewIERC20(common.HexToAddress(in.Address), d.provider())
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
	in, err := d.ts.get(r.In.String())
	if err != nil {
		return false, err
	}

	if in.Native {
		return false, nil
	}

	c, err := contracts.NewIERC20(common.HexToAddress(in.Address), d.provider())
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
