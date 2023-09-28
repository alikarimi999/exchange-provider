package evm

import (
	"exchange-provider/internal/delivery/exchanges/dex/evm/contracts/erc20"
	"exchange-provider/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	em "github.com/ethereum/go-ethereum/common/math"
)

func (d *exchange) approveTx(in *entity.Token) ([]byte, error) {
	return d.erc20.Pack("approve", d.cfg.contractAddress, em.MaxBig256)
}

func (d *exchange) needApproval(in *entity.Token, owner common.Address, minAmount float64) (bool, error) {
	if in.Native {
		return false, nil
	}

	c, err := erc20.NewContracts(common.HexToAddress(in.ContractAddress), d.provider())
	if err != nil {
		return false, err
	}

	amount, err := c.Allowance(nil, owner, d.cfg.contractAddress)
	if err != nil {
		return false, err
	}

	mAmount, _ := big.NewFloat(0).Mul(big.NewFloat(minAmount),
		big.NewFloat(0).SetInt(em.BigPow(10, int64(in.Decimals)))).Int(nil)
	return amount.Cmp(mAmount) == -1, nil
}
